#!/bin/bash

set -x

geth --networkid $networkid --mine --datadir $datadir --password $password --allow-insecure-unlock --gcmode archive   --unlock $unlockAccount   --http --http.addr "0.0.0.0" --http.port 8545 --http.vhosts "*"  --http.api "eth,scroll,net,web3,debug"   --ws --ws.addr "0.0.0.0" --ws.port 8546 --ws.api "eth,scroll,net,web3,debug" --l1.endpoint $l1_endpoint --l1.sync.startblock $startblock --l1.confirmations $confirmations