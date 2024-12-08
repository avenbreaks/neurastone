#!/bin/bash

KEY="mykey"
CHAINID="neura_121799-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t neura-datadir.XXXXX)

echo "create and add new keys"
./neurad keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Evmos with moniker=$MONIKER and chain-id=$CHAINID"
./neurad init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./neurad add-genesis-account \
"$(./neurad keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aphoton,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./neurad gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./neurad collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./neurad validate-genesis --home $DATA_DIR

echo "starting neura node $i in background ..."
./neurad start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started neura node"
tail -f /dev/null