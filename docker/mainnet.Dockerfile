ARG from_image

FROM $from_image

ENV CUSTOM_MONIKER="mainnet_seed_node"

RUN [[ ! -f $HOME/.neurad/config/genesis.json ]] && \
    neurad config chain-id neura_11235-1 && \
    neurad init $CUSTOM_MONIKER --chain-id neura_11235-1 && \
    curl -OL https://raw.githubusercontent.com/neura-network/mainnet/master/genesis.json && \
    mv genesis.json $HOME/.neurad/config/genesis.json && \
    curl -OL https://raw.githubusercontent.com/neura-network/mainnet/master/addrbook.json && \
    mv addrbook.json $HOME/.neurad/config/addrbook.json