// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {Vm} from "forge-std/Vm.sol";
import {KeyperSet} from "../src/common/KeyperSet.sol";
import {KeyperSetManager} from "../src/common/KeyperSetManager.sol";
import {DKGContract} from "../src/common/DKGContract.sol";
import {KeyBroadcastContract} from "../src/common/KeyBroadcastContract.sol";
import {IDKGContract} from "../src/common/intf/IDKGContract.sol";

contract DKGState is Script {
    bytes32 constant SIG_DEALING =
        keccak256("DealingSubmitted(uint64,uint64,uint64,bytes,bytes[])");
    bytes32 constant SIG_ACCUSATION =
        keccak256("AccusationSubmitted(uint64,uint64,uint64,uint64[])");
    bytes32 constant SIG_APOLOGY =
        keccak256("ApologySubmitted(uint64,uint64,uint64,uint64[],bytes[])");
    bytes32 constant SIG_SUCCESS_VOTE =
        keccak256("SuccessVoteSubmitted(uint64,uint64,uint64,bytes)");
    bytes32 constant SIG_DKG_SUCCEEDED =
        keccak256("DKGSucceeded(uint64,uint64,bytes)");

    struct Ctx {
        address managerAddr;
        uint64 setIndex;
        uint64 numKeyperSets;
        uint64 activationBlock;
        address keyperSetAddr;
        address[] members;
        uint64 threshold;
        address publisher;
        bool finalized;
        address dkgAddr;
        uint64 phaseLength;
        uint64 dkgLeadLength;
        uint64 cycleLen;
        int256 dkgStart0;
        bool succeededFlag;
        address keyBroadcastAddr;
        bytes broadcastedKey;
        uint256 currentBlock;
    }

    struct Logs {
        Vm.EthGetLogs[] dealing;
        Vm.EthGetLogs[] accusation;
        Vm.EthGetLogs[] apology;
        Vm.EthGetLogs[] vote;
        Vm.EthGetLogs[] success;
    }

    // First-seen order of unique eon-key hashes across fetched vote+success
    // events. A key's label is its index in this list.
    struct KeyReg {
        bytes32[] hashes;
    }

    function run() external {
        Ctx memory ctx = loadContext();
        printHeader(ctx);
        Logs memory logs = fetchLogs(ctx);
        KeyReg memory reg = buildKeyReg(logs);
        printKeyBroadcastLine(ctx, reg);
        printCurrentContext(ctx);
        printBody(ctx, logs, reg);
    }

    // ---------------- context load ----------------

    function loadContext() internal returns (Ctx memory ctx) {
        ctx.managerAddr = vm.envAddress("KEYPERSETMANAGER_ADDRESS");
        ctx.setIndex = uint64(vm.envUint("KEYPER_SET_INDEX"));

        KeyperSetManager manager = KeyperSetManager(ctx.managerAddr);
        ctx.numKeyperSets = manager.getNumKeyperSets();
        if (ctx.setIndex >= ctx.numKeyperSets) {
            revert(
                string.concat(
                    "KEYPER_SET_INDEX=",
                    vm.toString(uint256(ctx.setIndex)),
                    " out of range (numKeyperSets=",
                    vm.toString(uint256(ctx.numKeyperSets)),
                    " on this RPC; valid indices 0..",
                    ctx.numKeyperSets == 0
                        ? "(none)"
                        : vm.toString(uint256(ctx.numKeyperSets - 1)),
                    ")"
                )
            );
        }
        ctx.activationBlock = manager.getKeyperSetActivationBlock(ctx.setIndex);
        ctx.keyperSetAddr = manager.getKeyperSetAddress(ctx.setIndex);

        KeyperSet ks = KeyperSet(ctx.keyperSetAddr);
        ctx.members = ks.getMembers();
        ctx.threshold = ks.getThreshold();
        ctx.publisher = ks.getPublisher();
        ctx.finalized = ks.isFinalized();
        ctx.dkgAddr = ks.getDKGContract();

        DKGContract dkg = DKGContract(ctx.dkgAddr);
        ctx.phaseLength = dkg.PHASE_LENGTH();
        ctx.dkgLeadLength = dkg.DKG_LEAD_LENGTH();
        ctx.cycleLen = dkg.cycleLength();
        ctx.dkgStart0 = dkg.dkgStart(ctx.setIndex, 0);
        ctx.succeededFlag = dkg.succeeded(ctx.setIndex);
        ctx.keyBroadcastAddr = dkg.keyBroadcastContract();

        if (ctx.keyBroadcastAddr != address(0)) {
            KeyBroadcastContract kb = KeyBroadcastContract(
                ctx.keyBroadcastAddr
            );
            ctx.broadcastedKey = kb.getEonKey(ctx.setIndex);
        }
        ctx.currentBlock = block.number;
    }

    // ---------------- header ----------------

    function printHeader(Ctx memory c) internal view {
        console.log("=== DKG State ===");
        console.log("");
        console.log("KeyperSetManager: ", c.managerAddr);
        console.log("  numKeyperSets:  ", uint256(c.numKeyperSets));
        console.log("");
        console.log(
            string.concat(
                "Keyper set: index ",
                vm.toString(uint256(c.setIndex))
            )
        );
        console.log("  contract:         ", c.keyperSetAddr);
        console.log("  finalized:        ", c.finalized);
        console.log("  threshold:        ", uint256(c.threshold));
        console.log("  publisher:        ", c.publisher);
        console.log("  DKG contract:     ", c.dkgAddr);
        console.log("  activation block: ", uint256(c.activationBlock));
        console.log(
            string.concat(
                "  members (",
                vm.toString(c.members.length),
                "):"
            )
        );
        for (uint256 i = 0; i < c.members.length; i++) {
            console.log(
                string.concat(
                    "    [",
                    vm.toString(i),
                    "] ",
                    vm.toString(c.members[i])
                )
            );
        }
        console.log("");
        console.log("DKG contract config:");
        console.log("  PHASE_LENGTH:      ", uint256(c.phaseLength));
        console.log("  DKG_LEAD_LENGTH:   ", uint256(c.dkgLeadLength));
        console.log("  cycleLength:       ", uint256(c.cycleLen));
        console.log(
            string.concat("  dkgStart(index,0): ", _intToString(c.dkgStart0))
        );
        console.log("  succeeded[index]:  ", c.succeededFlag);
        console.log("");
        console.log("KeyBroadcastContract: ", c.keyBroadcastAddr);
    }

    function printKeyBroadcastLine(
        Ctx memory c,
        KeyReg memory reg
    ) internal view {
        if (c.keyBroadcastAddr == address(0)) {
            console.log("  eon key broadcast: (KeyBroadcastContract unset)");
            return;
        }
        if (c.broadcastedKey.length == 0) {
            console.log("  eon key broadcast: (none)");
        } else {
            bytes32 h = keccak256(c.broadcastedKey);
            string memory desc = _keyDesc(reg, h);
            console.log(
                string.concat("  eon key broadcast: ", desc)
            );
        }
    }

    function printCurrentContext(Ctx memory c) internal view {
        console.log("");
        console.log("Chain context:");
        console.log("  current block: ", c.currentBlock);
        int256 curR = _currentRetry(c);
        if (curR < 0) {
            console.log("  current retry: (DKG has not started yet)");
            return;
        }
        console.log("  current retry: ", uint256(curR));
        if (c.succeededFlag) {
            console.log(
                "  current phase: (DKG already succeeded; no active phase)"
            );
            return;
        }
        uint256 retryStart = _retryStartBlock(c, uint64(uint256(curR)));
        uint256 offset = c.currentBlock - retryStart;
        uint256 phaseIdx = offset / uint256(c.phaseLength);
        uint256 blocksInto = offset % uint256(c.phaseLength);
        string memory phaseName;
        if (phaseIdx == 0) phaseName = "Dealing";
        else if (phaseIdx == 1) phaseName = "Accusing";
        else if (phaseIdx == 2) phaseName = "Apologizing";
        else phaseName = "Finalizing";
        console.log(
            string.concat(
                "  current phase: ",
                phaseName,
                " (block ",
                vm.toString(blocksInto + 1),
                " of ",
                vm.toString(uint256(c.phaseLength)),
                ")"
            )
        );
    }

    // ---------------- fetch logs ----------------

    function fetchLogs(Ctx memory ctx) internal returns (Logs memory logs) {
        uint256 fromBlock = ctx.dkgStart0 < 0
            ? 0
            : uint256(ctx.dkgStart0);
        uint256 toBlock = ctx.currentBlock;
        if (fromBlock > toBlock) {
            logs.dealing = new Vm.EthGetLogs[](0);
            logs.accusation = new Vm.EthGetLogs[](0);
            logs.apology = new Vm.EthGetLogs[](0);
            logs.vote = new Vm.EthGetLogs[](0);
            logs.success = new Vm.EthGetLogs[](0);
            return logs;
        }
        bytes32 setIdxTopic = bytes32(uint256(ctx.setIndex));
        logs.dealing = _fetch(
            fromBlock,
            toBlock,
            ctx.dkgAddr,
            SIG_DEALING,
            setIdxTopic
        );
        logs.accusation = _fetch(
            fromBlock,
            toBlock,
            ctx.dkgAddr,
            SIG_ACCUSATION,
            setIdxTopic
        );
        logs.apology = _fetch(
            fromBlock,
            toBlock,
            ctx.dkgAddr,
            SIG_APOLOGY,
            setIdxTopic
        );
        logs.vote = _fetch(
            fromBlock,
            toBlock,
            ctx.dkgAddr,
            SIG_SUCCESS_VOTE,
            setIdxTopic
        );
        logs.success = _fetch(
            fromBlock,
            toBlock,
            ctx.dkgAddr,
            SIG_DKG_SUCCEEDED,
            setIdxTopic
        );
    }

    function _fetch(
        uint256 fromB,
        uint256 toB,
        address emitter,
        bytes32 sig,
        bytes32 setIdxTopic
    ) internal returns (Vm.EthGetLogs[] memory) {
        bytes32[] memory topics = new bytes32[](2);
        topics[0] = sig;
        topics[1] = setIdxTopic;
        return vm.eth_getLogs(fromB, toB, emitter, topics);
    }

    // ---------------- key registry ----------------

    function buildKeyReg(
        Logs memory logs
    ) internal pure returns (KeyReg memory reg) {
        uint256 maxCount = logs.vote.length + logs.success.length;
        bytes32[] memory buf = new bytes32[](maxCount);
        uint256 n = 0;
        for (uint256 i = 0; i < logs.vote.length; i++) {
            bytes memory k = abi.decode(logs.vote[i].data, (bytes));
            bytes32 h = keccak256(k);
            if (!_contains(buf, n, h)) {
                buf[n++] = h;
            }
        }
        for (uint256 i = 0; i < logs.success.length; i++) {
            bytes memory k = abi.decode(logs.success[i].data, (bytes));
            bytes32 h = keccak256(k);
            if (!_contains(buf, n, h)) {
                buf[n++] = h;
            }
        }
        bytes32[] memory out = new bytes32[](n);
        for (uint256 i = 0; i < n; i++) out[i] = buf[i];
        reg.hashes = out;
    }

    function _contains(
        bytes32[] memory arr,
        uint256 len,
        bytes32 v
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < len; i++) if (arr[i] == v) return true;
        return false;
    }

    function _labelOf(
        KeyReg memory reg,
        bytes32 keyHash
    ) internal pure returns (int256) {
        for (uint256 i = 0; i < reg.hashes.length; i++) {
            if (reg.hashes[i] == keyHash) return int256(i);
        }
        return -1;
    }

    function _keyDesc(
        KeyReg memory reg,
        bytes32 keyHash
    ) internal pure returns (string memory) {
        int256 idx = _labelOf(reg, keyHash);
        if (idx < 0) return "unknown key (not seen in votes)";
        if (reg.hashes.length <= 1) return "the eon key";
        return string.concat("KEY#", vm.toString(uint256(idx)));
    }

    // ---------------- body ----------------

    function printBody(
        Ctx memory c,
        Logs memory logs,
        KeyReg memory reg
    ) internal view {
        console.log("");
        console.log("--- Retries ---");

        int256 currentRetry = _currentRetry(c);
        if (currentRetry < 0) {
            console.log("(no retries have started yet)");
            return;
        }

        int256 succeededAtRetry = -1;
        for (uint256 i = 0; i < logs.success.length; i++) {
            int256 r = int256(
                uint256(uint64(uint256(logs.success[i].topics[2])))
            );
            if (succeededAtRetry < 0 || r < succeededAtRetry) {
                succeededAtRetry = r;
            }
        }
        if (c.succeededFlag && succeededAtRetry < 0) {
            console.log(
                "WARNING: succeeded[index] is true but no DKGSucceeded event"
            );
            console.log(
                "         was fetched in the query range. Iterating to current retry."
            );
        }
        int256 lastRetry = succeededAtRetry >= 0
            ? succeededAtRetry
            : currentRetry;

        int256 condStart = -1;
        int256 condEnd = -1;
        for (int256 r = 0; r <= lastRetry; r++) {
            bool inProgress = (r == currentRetry) && !c.succeededFlag;
            uint64 ru = uint64(uint256(r));
            bool hasEvents = _retryHasEvents(logs, ru);
            if (!hasEvents && !inProgress) {
                if (condStart < 0) condStart = r;
                condEnd = r;
                continue;
            }
            if (condStart >= 0) {
                _printCondensed(condStart, condEnd);
                condStart = -1;
            }
            _printRetry(c, logs, reg, ru, inProgress, ru == uint64(uint256(succeededAtRetry >= 0 ? succeededAtRetry : int256(-1))));
        }
        if (condStart >= 0) _printCondensed(condStart, condEnd);
    }

    function _printCondensed(int256 startR, int256 endR) internal view {
        if (startR == endR) {
            console.log(
                string.concat(
                    "Retry ",
                    vm.toString(uint256(startR)),
                    ": no messages"
                )
            );
        } else {
            console.log(
                string.concat(
                    "Retries ",
                    vm.toString(uint256(startR)),
                    "-",
                    vm.toString(uint256(endR)),
                    ": no messages"
                )
            );
        }
    }

    function _printRetry(
        Ctx memory c,
        Logs memory logs,
        KeyReg memory reg,
        uint64 r,
        bool inProgress,
        bool succeededHere
    ) internal view {
        console.log("");
        uint256 startB = _retryStartBlock(c, r);
        uint256 endB = startB + uint256(c.cycleLen) - 1;
        string memory tag;
        if (succeededHere) tag = " [SUCCEEDED]";
        else if (inProgress) tag = " [current]";
        else tag = "";
        console.log(
            string.concat(
                "Retry ",
                vm.toString(uint256(r)),
                " (blocks ",
                vm.toString(startB),
                "-",
                vm.toString(endB),
                ")",
                tag
            )
        );
        _printPhaseWindows(c, startB, inProgress);
        _printDealings(c, logs, r);
        _printAccusations(c, logs, r);
        _printApologies(c, logs, r);
        _printVotes(c, logs, reg, r, succeededHere);
    }

    function _printPhaseWindows(
        Ctx memory c,
        uint256 startB,
        bool inProgress
    ) internal view {
        uint256 pl = uint256(c.phaseLength);
        if (!inProgress) {
            console.log(
                string.concat(
                    "  Phases: Dealing ",
                    vm.toString(startB),
                    "-",
                    vm.toString(startB + pl - 1),
                    " | Accusing ",
                    vm.toString(startB + pl),
                    "-",
                    vm.toString(startB + 2 * pl - 1),
                    " | Apologizing ",
                    vm.toString(startB + 2 * pl),
                    "-",
                    vm.toString(startB + 3 * pl - 1),
                    " | Finalizing ",
                    vm.toString(startB + 3 * pl),
                    "-",
                    vm.toString(startB + 4 * pl - 1)
                )
            );
        } else {
            uint256 offset = c.currentBlock - startB;
            uint256 phaseIdx = offset / pl;
            uint256 blocksInto = offset % pl;
            string memory phaseName;
            if (phaseIdx == 0) phaseName = "Dealing";
            else if (phaseIdx == 1) phaseName = "Accusing";
            else if (phaseIdx == 2) phaseName = "Apologizing";
            else phaseName = "Finalizing";
            console.log(
                string.concat(
                    "  Currently in ",
                    phaseName,
                    " phase (block ",
                    vm.toString(blocksInto + 1),
                    " of ",
                    vm.toString(pl),
                    ")"
                )
            );
        }
    }

    function _printDealings(
        Ctx memory c,
        Logs memory logs,
        uint64 r
    ) internal view {
        uint256 total = c.members.length;
        bool[] memory submitted = new bool[](total);
        uint256 count = 0;
        for (uint256 i = 0; i < logs.dealing.length; i++) {
            if (uint64(uint256(logs.dealing[i].topics[2])) != r) continue;
            uint64 keyperIdx = uint64(uint256(logs.dealing[i].topics[3]));
            if (keyperIdx < total && !submitted[keyperIdx]) {
                submitted[keyperIdx] = true;
                count++;
            }
        }
        if (count == 0) {
            console.log(
                string.concat(
                    "  Dealings: none (0/",
                    vm.toString(total),
                    ")"
                )
            );
            return;
        }
        console.log(
            string.concat(
                "  Dealings (",
                vm.toString(count),
                "/",
                vm.toString(total),
                "):"
            )
        );
        console.log(
            string.concat("    submitted: ", _indexList(submitted, true))
        );
        if (count < total) {
            console.log(
                string.concat("    SILENT:    ", _indexList(submitted, false))
            );
        }
    }

    function _printAccusations(
        Ctx memory,
        Logs memory logs,
        uint64 r
    ) internal view {
        uint256 n = 0;
        for (uint256 i = 0; i < logs.accusation.length; i++) {
            if (uint64(uint256(logs.accusation[i].topics[2])) == r) n++;
        }
        if (n == 0) {
            console.log("  Accusations: none");
            return;
        }
        console.log(string.concat("  Accusations (", vm.toString(n), "):"));
        for (uint256 i = 0; i < logs.accusation.length; i++) {
            if (uint64(uint256(logs.accusation[i].topics[2])) != r) continue;
            uint64 accuser = uint64(
                uint256(logs.accusation[i].topics[3])
            );
            uint64[] memory accused = abi.decode(
                logs.accusation[i].data,
                (uint64[])
            );
            console.log(
                string.concat(
                    "    keyper[",
                    vm.toString(uint256(accuser)),
                    "] accused ",
                    _uint64ListToString(accused)
                )
            );
        }
    }

    function _printApologies(
        Ctx memory,
        Logs memory logs,
        uint64 r
    ) internal view {
        uint256 n = 0;
        for (uint256 i = 0; i < logs.apology.length; i++) {
            if (uint64(uint256(logs.apology[i].topics[2])) == r) n++;
        }
        if (n == 0) {
            console.log("  Apologies: none");
            return;
        }
        console.log(string.concat("  Apologies (", vm.toString(n), "):"));
        for (uint256 i = 0; i < logs.apology.length; i++) {
            if (uint64(uint256(logs.apology[i].topics[2])) != r) continue;
            uint64 apologizer = uint64(
                uint256(logs.apology[i].topics[3])
            );
            (uint64[] memory accusers, ) = abi.decode(
                logs.apology[i].data,
                (uint64[], bytes[])
            );
            console.log(
                string.concat(
                    "    keyper[",
                    vm.toString(uint256(apologizer)),
                    "] apologized to ",
                    _uint64ListToString(accusers)
                )
            );
        }
    }

    struct VoteTally {
        uint256[] tallies;
        bool[][] voterByKey;
        bool[] voted;
        uint256 votedCount;
    }

    function _tallyVotes(
        Ctx memory c,
        Logs memory logs,
        KeyReg memory reg,
        uint64 r
    ) internal pure returns (VoteTally memory t) {
        uint256 total = c.members.length;
        uint256 numKeys = reg.hashes.length;
        t.tallies = new uint256[](numKeys);
        t.voterByKey = new bool[][](numKeys);
        for (uint256 k = 0; k < numKeys; k++) {
            t.voterByKey[k] = new bool[](total);
        }
        t.voted = new bool[](total);
        for (uint256 i = 0; i < logs.vote.length; i++) {
            if (uint64(uint256(logs.vote[i].topics[2])) != r) continue;
            uint64 keyperIdx = uint64(uint256(logs.vote[i].topics[3]));
            if (keyperIdx >= total) continue;
            bytes memory k = abi.decode(logs.vote[i].data, (bytes));
            int256 label = _labelOf(reg, keccak256(k));
            if (label < 0) continue;
            uint256 li = uint256(label);
            if (!t.voted[keyperIdx]) {
                t.voted[keyperIdx] = true;
                t.votedCount++;
            }
            if (!t.voterByKey[li][keyperIdx]) {
                t.voterByKey[li][keyperIdx] = true;
                t.tallies[li]++;
            }
        }
    }

    function _successKeyHash(
        Logs memory logs,
        uint64 r
    ) internal pure returns (bytes32) {
        for (uint256 i = 0; i < logs.success.length; i++) {
            if (uint64(uint256(logs.success[i].topics[2])) == r) {
                bytes memory k = abi.decode(logs.success[i].data, (bytes));
                return keccak256(k);
            }
        }
        return bytes32(0);
    }

    function _printVotes(
        Ctx memory c,
        Logs memory logs,
        KeyReg memory reg,
        uint64 r,
        bool succeededHere
    ) internal view {
        VoteTally memory t = _tallyVotes(c, logs, reg, r);
        bytes32 successHash = _successKeyHash(logs, r);
        uint256 total = c.members.length;

        if (t.votedCount == 0 && !succeededHere) {
            console.log(
                string.concat(
                    "  Success votes: none (threshold ",
                    vm.toString(uint256(c.threshold)),
                    ")"
                )
            );
            return;
        }
        console.log(
            string.concat(
                "  Success votes (",
                vm.toString(t.votedCount),
                "/",
                vm.toString(total),
                ", threshold ",
                vm.toString(uint256(c.threshold)),
                "):"
            )
        );
        for (uint256 li = 0; li < reg.hashes.length; li++) {
            if (t.tallies[li] == 0) continue;
            _printOneKeyTally(reg, t, li, successHash);
        }
        if (t.votedCount < total) {
            console.log(
                string.concat("    NON-VOTERS: ", _indexList(t.voted, false))
            );
        }
        if (succeededHere) {
            console.log(
                string.concat(
                    "  Outcome: DKGSucceeded emitted (",
                    _keyDesc(reg, successHash),
                    ")"
                )
            );
        }
    }

    function _printOneKeyTally(
        KeyReg memory reg,
        VoteTally memory t,
        uint256 li,
        bytes32 successHash
    ) internal view {
        string memory winnerMark = (successHash != bytes32(0) &&
            reg.hashes[li] == successHash)
            ? "  <- SUCCEEDED"
            : "";
        console.log(
            string.concat(
                "    ",
                _keyDesc(reg, reg.hashes[li]),
                ": ",
                vm.toString(t.tallies[li]),
                " votes",
                winnerMark
            )
        );
        console.log(
            string.concat(
                "      voters: ",
                _indexList(t.voterByKey[li], true)
            )
        );
    }

    // ---------------- helpers ----------------

    function _currentRetry(Ctx memory c) internal pure returns (int256) {
        int256 rel = int256(c.currentBlock) - c.dkgStart0;
        if (rel < 0) return -1;
        return rel / int256(uint256(c.cycleLen));
    }

    function _retryStartBlock(
        Ctx memory c,
        uint64 r
    ) internal pure returns (uint256) {
        int256 s = c.dkgStart0 + int256(uint256(r)) * int256(uint256(c.cycleLen));
        // Caller must not invoke for retries with negative start; guarded by
        // _currentRetry returning -1 upstream.
        return uint256(s);
    }

    function _retryHasEvents(
        Logs memory logs,
        uint64 r
    ) internal pure returns (bool) {
        return
            _anyAt(logs.dealing, r) ||
            _anyAt(logs.accusation, r) ||
            _anyAt(logs.apology, r) ||
            _anyAt(logs.vote, r) ||
            _anyAt(logs.success, r);
    }

    function _anyAt(
        Vm.EthGetLogs[] memory arr,
        uint64 r
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < arr.length; i++) {
            if (uint64(uint256(arr[i].topics[2])) == r) return true;
        }
        return false;
    }

    function _indexList(
        bool[] memory flags,
        bool matchValue
    ) internal pure returns (string memory) {
        string memory out = "";
        bool first = true;
        for (uint256 i = 0; i < flags.length; i++) {
            if (flags[i] != matchValue) continue;
            if (first) {
                out = string.concat("[", vm.toString(i));
                first = false;
            } else {
                out = string.concat(out, ",", vm.toString(i));
            }
        }
        if (first) return "[]";
        return string.concat(out, "]");
    }

    function _uint64ListToString(
        uint64[] memory xs
    ) internal pure returns (string memory) {
        if (xs.length == 0) return "[]";
        string memory out = string.concat("[", vm.toString(uint256(xs[0])));
        for (uint256 i = 1; i < xs.length; i++) {
            out = string.concat(out, ",", vm.toString(uint256(xs[i])));
        }
        return string.concat(out, "]");
    }

    function _intToString(int256 x) internal pure returns (string memory) {
        if (x < 0) return string.concat("-", vm.toString(uint256(-x)));
        return vm.toString(uint256(x));
    }
}
