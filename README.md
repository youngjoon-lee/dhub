# DHub (DataHub) Chain

**dhub** is a blockchain built using Cosmos SDK and Tendermint and co-works with the [decentralized oracle](https://github.com/youngjoon-lee/doracle-poc).

## Get started

Build the binary.
```bash
make install
```

Configure a new chain.
```bash
dhubd init node1 --chain-id dhub-1
dhubd keys add alice
dhubd keys add bob
dhubd add-genesis-account $(dhubd keys show alice -a) 100000000000000uhub
dhubd add-genesis-account $(dhubd keys show bob -a) 100000000000000uhub
dhubd gentx validator 1000000000000uhub --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01  --min-self-delegation 1000000 --chain-id testing
dhubd collect-gentxs
```

Start the chain.
```bash
dhubd start
```
