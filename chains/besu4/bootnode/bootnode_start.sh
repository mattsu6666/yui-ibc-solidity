#!/usr/bin/env bash

set -eu

P2P_HOST=$(awk 'END{print $1}' /etc/hosts)

mkdir -p ${PWD}/data
ls -d1 ${PWD}/networkFiles/keys/* | \
    sed -n ${NODE_INDEX}p | \
    xargs -I '{}' sh -c 'cp {}/* ./data'

besu --data-path=./data --genesis-file=./networkFiles/genesis.json --p2p-host=${P2P_HOST} --p2p-port=${P2P_PORT} --rpc-http-enabled --rpc-http-api=ADMIN,ETH,NET,IBFT --rpc-ws-enabled --rpc-ws-api=ADMIN,ETH,NET,IBFT --host-allowlist="*" --rpc-http-cors-origins="all" --rpc-http-port=${RPC_PORT} --rpc-ws-port=${RPC_WS_PORT} --revert-reason-enabled
