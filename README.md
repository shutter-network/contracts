# Shutter Network Contracts

This repository contains the EVM contracts for [Shutter](https://www.shutter.network/).

## Installation

Go bindings for the shutter contracts can be generated by running

    ./gen_bindings.sh

Released versions are available as go package under `github.com/shutter-network/contracts/bindings/${contract name in lowercase}`.

## Scripts

There are some helper scripts (deployment, executing special contract methods, ...) in the `script/` directory.

**TODO: Add information for the involved scripts.**

## Tests

To manually run the tests, execute

    forge test -vvv

## Contribution

Contributions are welcome! Feel free to [create issues](https://github.com/shutter-network/contracts/issues/new/choose) and/or open [pull requests](https://github.com/shutter-network/contracts/compare).

## References / other repositories

- [shutter.network](https://shutter.network)
- [shutter-network/rolling-shutter](https://github.com/shutter-network/rolling-shutter)
- [Gnosis chain spec for Shutter](https://github.com/gnosischain/specs/tree/master/shutter)
