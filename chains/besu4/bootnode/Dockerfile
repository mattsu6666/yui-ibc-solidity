# TODO: Update version
ARG BESU_VER=21.1
FROM hyperledger/besu:${BESU_VER}

USER root

WORKDIR /tmp/besu
ADD ./networkFiles /tmp/besu/networkFiles/
ADD ./bootnode/bootnode_start.sh /tmp/besu/

EXPOSE $RPC_PORT $P2P_PORT

ENTRYPOINT ./bootnode_start.sh
