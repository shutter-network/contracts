// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../src/common/DKGContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyBroadcastContract.sol";

/// @title DKG Contract gas benchmark
/// @notice Measures total on-chain gas for a complete DKG run at multiple
/// keyper counts. Setup gas (contract deployment, keyper set creation) is
/// excluded — only protocol operation cost is reported.
contract DKGBenchmark is Test {
    // Phase parameters mirror the existing DKGContract.t.sol conventions so
    // the benchmark uses block windows the test runner already understands.
    uint64 constant PHASE_LENGTH = 10;
    uint64 constant DKG_LEAD_LENGTH = 40;
    uint64 constant ACTIVATION_BLOCK = 1000;

    // Non-zero fill byte. Real BLS/ECIES payloads are non-zero, and calldata
    // gas is 16 per non-zero byte vs. 4 per zero byte — so using 0xAB makes
    // the measurement reflect production cost.
    bytes1 constant FILL = 0xAB;

    DKGContract dkg;
    KeyperSetManager ksm;
    KeyBroadcastContract kbc;
    KeyperSet ks;
    address[] keypers;

    uint256 totalGas;

    // -------------------------------------------------------------------
    // Setup helpers — deploy fresh contracts and a new keyper set per run
    // to keep warm-storage effects from leaking between scenarios.
    // -------------------------------------------------------------------

    function _setup(uint64 n) internal {
        address initializer = address(69);
        address dao = address(42);
        address sequencer = address(420);

        ksm = new KeyperSetManager(initializer);
        vm.prank(initializer);
        ksm.initialize(dao, sequencer);

        kbc = new KeyBroadcastContract(address(ksm));
        dkg = new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            address(ksm),
            address(kbc)
        );

        ks = new KeyperSet();
        address[] memory members = new address[](n);
        delete keypers;
        for (uint64 i = 0; i < n; i++) {
            address k = address(uint160(0x1000 + i));
            members[i] = k;
            keypers.push(k);
        }
        ks.addMembers(members);
        ks.setThreshold(_threshold(n));
        ks.setPublisher(address(dkg));
        ks.setDKGContract(address(dkg));
        ks.setFinalized();

        vm.prank(dao);
        ksm.addKeyperSet(ACTIVATION_BLOCK, address(ks));
    }

    // ⌈2n/3⌉ via integer arithmetic.
    function _threshold(uint64 n) internal pure returns (uint64) {
        return (2 * n + 2) / 3;
    }

    // -------------------------------------------------------------------
    // Data builders — produce realistic-sized non-zero payloads.
    // -------------------------------------------------------------------

    function _fill(uint256 size) internal pure returns (bytes memory result) {
        result = new bytes(size);
        for (uint256 i = 0; i < size; i++) {
            result[i] = FILL;
        }
    }

    // PolyCommitment: one compressed BLS12-381 G2 point per polynomial
    // coefficient, plus a 4-byte length prefix → 4 + threshold × 96 bytes.
    function _commitment(uint64 threshold) internal pure returns (bytes memory) {
        return _fill(4 + uint256(threshold) * 96);
    }

    // PolyEvals: one ECIES ciphertext per other keyper, each 129 bytes
    // (65 ephemeral pubkey + 32 ciphertext + 32 HMAC-SHA256).
    function _polyEvals(uint64 n) internal pure returns (bytes[] memory) {
        bytes[] memory evals = new bytes[](n - 1);
        for (uint64 i = 0; i < n - 1; i++) {
            evals[i] = _fill(129);
        }
        return evals;
    }

    // Eon public key: compressed BLS12-381 G2 point, 96 bytes.
    function _eonKey() internal pure returns (bytes memory) {
        return _fill(96);
    }

    // Accusation: every other keyper's index, n-1 uint64 entries.
    function _accusedIndices(
        uint64 n,
        uint64 selfIdx
    ) internal pure returns (uint64[] memory) {
        uint64[] memory accused = new uint64[](n - 1);
        uint64 j = 0;
        for (uint64 i = 0; i < n; i++) {
            if (i == selfIdx) continue;
            accused[j++] = i;
        }
        return accused;
    }

    // Apology polyEvalData: one plaintext BLS12-381 scalar per accuser, 32
    // bytes each (NOT 129-byte ECIES — the accused keyper reveals the
    // evaluation in the clear during apology).
    function _apologyEvals(uint64 n) internal pure returns (bytes[] memory) {
        bytes[] memory evals = new bytes[](n - 1);
        for (uint64 i = 0; i < n - 1; i++) {
            evals[i] = _fill(32);
        }
        return evals;
    }

    // -------------------------------------------------------------------
    // Scenario runner — happy path or worst case.
    //   happy path: n dealings + ⌈2n/3⌉ success votes.
    //   worst case: same + n accusations (all-vs-all) + n apologies (all
    //   accusers acknowledged with plaintext scalars).
    // -------------------------------------------------------------------

    function _run(uint64 n, bool worstCase) internal {
        _setup(n);
        uint64 threshold = _threshold(n);
        totalGas = 0;

        uint64 dealingBlock = ACTIVATION_BLOCK - DKG_LEAD_LENGTH;
        bytes memory commitment = _commitment(threshold);
        bytes[] memory polyEvals = _polyEvals(n);

        vm.roll(dealingBlock);
        for (uint64 i = 0; i < n; i++) {
            vm.prank(keypers[i]);
            uint256 before = gasleft();
            dkg.submitDealing(0, 0, i, commitment, polyEvals);
            totalGas += before - gasleft();
        }

        if (worstCase) {
            uint64 accusingBlock = dealingBlock + PHASE_LENGTH;
            vm.roll(accusingBlock);
            for (uint64 i = 0; i < n; i++) {
                uint64[] memory accused = _accusedIndices(n, i);
                vm.prank(keypers[i]);
                uint256 before = gasleft();
                dkg.submitAccusation(0, 0, i, accused);
                totalGas += before - gasleft();
            }

            uint64 apologizingBlock = dealingBlock + 2 * PHASE_LENGTH;
            bytes[] memory apologyEvals = _apologyEvals(n);
            vm.roll(apologizingBlock);
            for (uint64 i = 0; i < n; i++) {
                uint64[] memory accusers = _accusedIndices(n, i);
                vm.prank(keypers[i]);
                uint256 before = gasleft();
                dkg.submitApology(0, 0, i, accusers, apologyEvals);
                totalGas += before - gasleft();
            }
        }

        uint64 finalizingBlock = dealingBlock + 3 * PHASE_LENGTH;
        bytes memory eonKey = _eonKey();

        vm.roll(finalizingBlock);
        for (uint64 i = 0; i < threshold; i++) {
            vm.prank(keypers[i]);
            uint256 before = gasleft();
            dkg.submitSuccessVote(0, 0, i, eonKey);
            totalGas += before - gasleft();
        }

        assertTrue(dkg.succeeded(0), "DKG must reach success");
        console.log(
            worstCase
                ? "DKGBenchmark scenario=worst-case n=%s threshold=%s gas=%s"
                : "DKGBenchmark scenario=happy-path n=%s threshold=%s gas=%s",
            uint256(n),
            uint256(threshold),
            totalGas
        );
    }

    function test_happyPath_n3() public {
        _run(3, false);
    }

    function test_happyPath_n5() public {
        _run(5, false);
    }

    function test_happyPath_n10() public {
        _run(10, false);
    }

    function test_happyPath_n20() public {
        _run(20, false);
    }

    function test_happyPath_n50() public {
        _run(50, false);
    }

    function test_worstCase_n3() public {
        _run(3, true);
    }

    function test_worstCase_n5() public {
        _run(5, true);
    }

    function test_worstCase_n10() public {
        _run(10, true);
    }

    function test_worstCase_n20() public {
        _run(20, true);
    }

    function test_worstCase_n50() public {
        _run(50, true);
    }
}
