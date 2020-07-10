# Stakebird DAO

Stakebird is a [content curation DAO](https://ethresear.ch/t/prediction-markets-for-content-curation-daos/1312).

Testnet coming soon.

## Install

Stakebird is a high performance sovereign chain that interoperates with the Cosmos Hub via [IBC](https://cosmos.network/ibc), and eventually Bitcoin and Ethereum as well. 

### Run a local, single-node chain

```sh
# install binaries
make install

# create keys
make create-wallet

# initialize chain
make init

# run
staked start
```

### Run a local testnet with a connection to Gaia (Cosmos Hub)

Stakebird requires [Gaia](https://github.com/cosmos/gaia) (Cosmos Hub), and an IBC [relayer](https://github.com/iqlusioninc/relayer) to facilitate connections.

To setup a local testnet running Gaia and Stakebird, run:
```
./contrib/ibc/gaia-stakebird.sh
```

To setup the relayer and do a token transfer between chains, run:
```
./contrib/ibc/stakebird-xfer.sh
```

### Ethereum Interoperability (Peggy)

#### Setup
```shell script
cd testnet-contracts/

# Create .env with environment variables required for contract deployment
cp .env.example .env

# Download dependencies
yarn
```

#### Start local Ethereum blockchain simulator (Ganache)
```shell script
# Open a new terminal window

# Start local blockchain
yarn develop
```

#### Compile and deploy bridge contracts
```shell script
# Open a new terminal window (in project root)

# Deploy and set up contracts, mint ERC20 TEST tokens and approve some to bank contract
yarn peggy:all
```

## CLI
The curating module can be accessed via CLI and REST API.

```sh
stakecli tx curating post -h
stakecli tx curating upvote -h
```
