# gnosh-contracts

## Running the tests

Install [foundry](https://book.getfoundry.sh/getting-started/installation).

Run `forge test -vvv`

## Using the deploy script

Create a `.env` file with the following content, fill in the missing values:

```shell
GOERLI_RPC_URL=https://ethereum-goerli.publicnode.com
PRIVATE_KEY=0xe1...
ETHERSCAN_API_KEY=BIR44...
```

Run the following command:

```shell
source .env
forge script script/deploy.s.sol --rpc-url $GOERLI_RPC_URL --broadcast --verify -vvvv
```
