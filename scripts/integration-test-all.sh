#!/bin/bash

# "stable" mode tests assume data is static
# "live" mode tests assume data dynamic

SCRIPT=$(basename ${BASH_SOURCE[0]})
TEST="rpc"
QTD=1
SLEEP_TIMEOUT=5
TEST_QTD=1

## remove test data dir after
REMOVE_DATA_DIR=false

#PORT AND RPC_PORT 3 initial digits, to be concat with a suffix later when node is initialized
RPC_PORT="854"
IP_ADDR="0.0.0.0"

KEY="mykey"
CHAINID="neura_1337-1"
MONIKER="mymoniker"
MODE="rpc"

## default port prefixes for neurad
NODE_P2P_PORT="2660"
NODE_PORT="2663"
NODE_RPC_PORT="2666"

usage() {
    echo "Usage: $SCRIPT"
    echo "Optional command line arguments"
    echo "-t <string>  -- Test to run. eg: rpc"
    echo "-q <number>  -- Quantity of nodes to run. eg: 3"
    echo "-z <number>  -- Quantity of nodes to run tests against eg: 3"
    echo "-s <number>  -- Sleep between operations in secs. eg: 5"
    echo "-r <string>  -- Remove test dir after, eg: true, default is false"
    exit 1
}

while getopts "h?t:q:z:s:m:r:" args; do
    case $args in
        h|\?)
            usage;
        exit;;
        t ) TEST=${OPTARG};;
        q ) QTD=${OPTARG};;
        z ) TEST_QTD=${OPTARG};;
        s ) SLEEP_TIMEOUT=${OPTARG};;
        m ) MODE=${OPTARG};;
        r ) REMOVE_DATA_DIR=${OPTARG};;
    esac
done

set -euxo pipefail

DATA_DIR=$(mktemp -d -t neura-datadir.XXXXX)

if [[ ! "$DATA_DIR" ]]; then
    echo "Could not create $DATA_DIR"
    exit 1
fi

# Compile evmos
echo "compiling evmos"
make build

# PID array declaration
arr=()

init_func() {
    # validate dependencies are installed
    command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

    HOME_DIR="$DATA_DIR$i"
    KEY_NAME=$KEY"$i"

    "$PWD"/build/neurad keys add $KEY_NAME --keyring-backend test --home $HOME_DIR --no-backup --algo "eth_secp256k1"
    "$PWD"/build/neurad init $MONIKER --chain-id $CHAINID --home $HOME_DIR

    # Change parameter token denominations to aISLM
    cat $HOME_DIR/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aISLM"' > $HOME_DIR/config/tmp_genesis.json && \
    mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json

    cat $HOME_DIR/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aISLM"' > $HOME_DIR/config/tmp_genesis.json && \
    mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json

    # cat $HOME_DIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aISLM"' > $HOME_DIR/config/tmp_genesis.json && \ 
    # mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json

    cat $HOME_DIR/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="aISLM"' > $HOME_DIR/config/tmp_genesis.json && \
    mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json

    cat $HOME_DIR/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="aISLM"' > $HOME_DIR/config/tmp_genesis.json && \
    mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json

    "$PWD"/build/neurad add-genesis-account \
    "$("$PWD"/build/neurad keys show "$KEY$i" --keyring-backend test -a --home "$HOME_DIR")" 1000000000000000000aISLM \
    --keyring-backend test --home $HOME_DIR
    "$PWD"/build/neurad gentx "$KEY$i" 1000000000000000000aISLM --chain-id $CHAINID --keyring-backend test --home $HOME_DIR
    "$PWD"/build/neurad collect-gentxs --home $HOME_DIR
    "$PWD"/build/neurad validate-genesis --home $HOME_DIR

    if [[ $MODE == "pending" ]]; then
      ls $DATA_DIR$i
      if [[ "$OSTYPE" == "darwin"* ]]; then
        # sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "1s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "120s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $DATA_DIR$i/config/config.toml
        sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $DATA_DIR$i/config/config.toml
      else
        # sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "1s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_prevote = "1s"/timeout_prevote = "120s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "2s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $DATA_DIR$i/config/config.toml
        sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $DATA_DIR$i/config/config.toml
      fi
    fi
}

start_func() {
    echo "starting neura node $i in background ..."
    "$PWD"/build/neurad start --pruning=nothing --rpc.unsafe \
    --p2p.laddr tcp://$IP_ADDR:$NODE_P2P_PORT"$i" --address tcp://$IP_ADDR:$NODE_PORT"$i" --rpc.laddr tcp://$IP_ADDR:$NODE_RPC_PORT"$i" \
    --json-rpc.address=$IP_ADDR:$RPC_PORT"$i" \
    --keyring-backend test --home "$DATA_DIR$i" \
    >"$DATA_DIR"/node"$i".log 2>&1 & disown

    neura_PID=$!
    echo "started neura node, pid=$neura_PID"
    # add PID to array
    arr+=("$neura_PID")

    if [[ $MODE == "pending" ]]; then
      echo "waiting for the first block..."
      sleep 300
    fi
}

# Run node with static blockchain database
# For loop N times
for i in $(seq 1 "$QTD"); do
    init_func "$i"
    start_func "$i"
    sleep 1
    echo "sleeping $SLEEP_TIMEOUT seconds for startup"
    sleep "$SLEEP_TIMEOUT"
    echo "done sleeping"
done

echo "sleeping $SLEEP_TIMEOUT seconds before running tests ... "
sleep "$SLEEP_TIMEOUT"
echo "done sleeping"

set +e

if [[ -z $TEST || $TEST == "rpc" ||  $TEST == "pending" ]]; then

    time_out=300s
    if [[ $TEST == "pending" ]]; then
      time_out=60m0s
    fi

    for i in $(seq 1 "$TEST_QTD"); do
        echo "\n### Priv key for Metamask ####"
        HOME_DIR="$DATA_DIR$i"
        PRIV_KEY=$(./build/neurad keys unsafe-export-eth-key $KEY_NAME --home $HOME_DIR --keyring-backend test)
        echo $PRIV_KEY

        sleep 5

        HOST_RPC=http://$IP_ADDR:$RPC_PORT"$i"
        echo "going to test neura node $HOST_RPC ..."
        MODE=$MODE HOST=$HOST_RPC PRIV_KEY=$PRIV_KEY go test ./tests/... -timeout=$time_out -v -short

        RPC_FAIL=$?
    done

fi

stop_func() {
    neura_PID=$i
    echo "shutting down node, pid=$neura_PID ..."

    # Shutdown evmos node
    kill -9 "$neura_PID"
    wait "$neura_PID"

    if [ $REMOVE_DATA_DIR == "true" ]
    then
        rm -rf $DATA_DIR*
    fi
}

for i in "${arr[@]}"; do
    stop_func "$i"
done

if [[ (-z $TEST || $TEST == "rpc") && $RPC_FAIL -ne 0 ]]; then
    exit $RPC_FAIL
else
    exit 0
fi