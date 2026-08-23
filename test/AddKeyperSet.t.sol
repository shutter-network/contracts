// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import {
    ActivationDeltaTooLow,
    AddKeyperSet,
    DuplicateKeyper,
    EmptyKeyperSet,
    ThresholdExceedsKeyperSetSize,
    ThresholdTooLow
} from "../script/AddKeyperSet.s.sol";

contract AddKeyperSetHarness is AddKeyperSet {
    function validateKeyperConfig(
        address[] memory keypers,
        uint256 threshold
    ) external pure {
        _validateKeyperConfig(keypers, threshold);
    }
}

contract AddKeyperSetTest is Test {
    AddKeyperSetHarness internal addKeyperSet;

    function setUp() public {
        addKeyperSet = new AddKeyperSetHarness();
    }

    function testRejectsEmptyKeyperSet() public {
        address[] memory keypers = new address[](0);

        vm.expectRevert(EmptyKeyperSet.selector);
        addKeyperSet.validateKeyperConfig(keypers, 1);
    }

    function testRejectsZeroThreshold() public {
        address[] memory keypers = _keypers(1);

        vm.expectRevert(ThresholdTooLow.selector);
        addKeyperSet.validateKeyperConfig(keypers, 0);
    }

    function testRejectsThresholdAboveKeyperSetSize() public {
        address[] memory keypers = _keypers(2);

        vm.expectRevert(
            abi.encodeWithSelector(
                ThresholdExceedsKeyperSetSize.selector,
                3,
                keypers.length
            )
        );
        addKeyperSet.validateKeyperConfig(keypers, 3);
    }

    function testRejectsDuplicateKeyper() public {
        address[] memory keypers = _keypers(3);
        keypers[2] = keypers[0];

        vm.expectRevert(
            abi.encodeWithSelector(DuplicateKeyper.selector, keypers[0])
        );
        addKeyperSet.validateKeyperConfig(keypers, 2);
    }

    function testAcceptsSingleKeyperWithThresholdOne() public view {
        address[] memory keypers = _keypers(1);

        addKeyperSet.validateKeyperConfig(keypers, 1);
    }

    function testAcceptsThresholdEqualToKeyperSetSize() public view {
        address[] memory keypers = _keypers(3);

        addKeyperSet.validateKeyperConfig(keypers, keypers.length);
    }

    function testAcceptsThresholdBelowSimpleMajority() public view {
        address[] memory keypers = _keypers(3);

        addKeyperSet.validateKeyperConfig(keypers, 1);
    }

    function testRunRejectsInvalidActivationDeltaBeforeBroadcastSetup() public {
        vm.setEnv("PRIVATE_KEY", "1");
        vm.setEnv("ACTIVATION_DELTA", "0");
        vm.setEnv("KEYPER_ADDRESSES", "not-an-address");

        vm.expectRevert(ActivationDeltaTooLow.selector);
        addKeyperSet.run();
    }

    function testRunRejectsInvalidKeyperConfigBeforeBroadcastSetup() public {
        vm.setEnv("PRIVATE_KEY", "1");
        vm.setEnv("ACTIVATION_DELTA", "1");
        vm.setEnv("KEYPERSETMANAGER_ADDRESS", "not-an-address");
        vm.setEnv(
            "KEYPER_ADDRESSES",
            "0x0000000000000000000000000000000000000001"
        );
        vm.setEnv("THRESHOLD", "2");

        vm.expectRevert(
            abi.encodeWithSelector(
                ThresholdExceedsKeyperSetSize.selector,
                2,
                1
            )
        );
        addKeyperSet.run();
    }

    function _keypers(
        uint256 count
    ) internal pure returns (address[] memory keypers) {
        keypers = new address[](count);
        for (uint256 i = 0; i < count; i++) {
            keypers[i] = address(uint160(i + 1));
        }
    }
}
