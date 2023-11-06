#!/bin/bash

set -x

if [ $IMPORT_GENESIS ];then 
    rollup_relayer --config $1 --import-genesis
else
    rollup_relayer --config $1
fi