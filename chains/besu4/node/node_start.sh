#!/usr/bin/env bash

set -eu

BOOTNODE_HOST=$(getent hosts ${BOOTNODE_SERVICE} | awk '{ print $1 }')
P2P_HOST=$(awk 'END{print $1}' /etc/hosts)

mkdir -p ${PWD}/data
ls -d1 ${PWD}/networkFiles/keys/* | \
    sed -n ${NODE_INDEX}p | \
    xargs -I '{}' sh -c 'cp {}/* ./data'

BOOTNODE=$(ls -d1 $PWD/networkFiles/keys/* | \
    sed -n ${BOOTNODE_INDEX}p | \
    xargs -I '{}' cat {}/key.pub | \
    awk -v host=${BOOTNODE_HOST} -v port=${BOOTNODE_PORT} \
      '{node_id=substr($0, 3); print "enode://" node_id "@" host ":" port}')

besu --data-path=./data --genesis-file=./networkFiles/genesis.json --bootnodes=${BOOTNODE}  --p2p-host=${P2P_HOST} --p2p-port=${P2P_PORT} --rpc-http-enabled --rpc-http-api=ADMIN,ETH,NET,IBFT --rpc-ws-enabled --rpc-ws-api=ADMIN,ETH,NET,IBFT --host-allowlist="*" --rpc-http-cors-origins="all" --rpc-http-port=${RPC_PORT} --rpc-ws-port=${RPC_WS_PORT} --revert-reason-enabled
