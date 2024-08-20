// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/common/intf/IValidatorRegistry.sol";
import "../src/gnosh/ValidatorRegistry.sol";

contract ValidatorRegistryTest is Test {
    event Updated(bytes message, bytes signature);

    IValidatorRegistry registry;

    function setUp() public {
        registry = new ValidatorRegistry();
    }

    function testUpdate() public {
        bytes memory message = bytes("msg");
        bytes memory signature = bytes("sig");

        vm.expectEmit(address(registry));
        emit Updated(message, signature);

        registry.update(message, signature);
    }

    function testGetters() public {
        assertEq(registry.getNumUpdates(), 0);
        vm.expectRevert();
        registry.getUpdate(0);

        IValidatorRegistry.Update memory up1 = IValidatorRegistry.Update({
            message: bytes("msg1"),
            signature: bytes("sig1")
        });
        IValidatorRegistry.Update memory up2 = IValidatorRegistry.Update({
            message: bytes("msg2"),
            signature: bytes("sig1")
        });
        IValidatorRegistry.Update memory up3 = IValidatorRegistry.Update({
            message: bytes("msg3"),
            signature: bytes("sig3")
        });

        registry.update(up1.message, up1.signature);
        registry.update(up2.message, up2.signature);
        registry.update(up3.message, up3.signature);

        assertEq(registry.getNumUpdates(), 3);
        assertTrue(cmpUpdates(registry.getUpdate(0), up1));
        assertTrue(cmpUpdates(registry.getUpdate(1), up2));
        assertTrue(cmpUpdates(registry.getUpdate(2), up3));
        vm.expectRevert();
        registry.getUpdate(3);
    }

    function cmpUpdates(
        IValidatorRegistry.Update memory up1,
        IValidatorRegistry.Update memory up2
    ) internal pure returns (bool) {
        return keccak256(abi.encode(up1)) == keccak256(abi.encode(up2));
    }
}
