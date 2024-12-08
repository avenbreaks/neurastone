ARG from_image

FROM $from_image

ENV CUSTOM_MONIKER="testedge2_seed_node"

RUN [[ ! -f $HOME/.neurd/config/genesis.json ]] && \
    neurd config chain-id neura_54211-3 && \
    neurd init $CUSTOM_MONIKER --chain-id neura_54211-3 && \
    curl -OL https://raw.githubusercontent.com/neura-network/testnets/main/TestEdge2/genesis.tar.bz2 &&\
    bzip2 -d genesis.tar.bz2 && tar -xvf genesis.tar && \
    mv genesis.json $HOME/.neurd/config/genesis.json && \
    curl -OL https://raw.githubusercontent.com/neura-network/testnets/main/TestEdge2/addrbook.json && \
    mv addrbook.json $HOME/.neurd/config/addrbook.json