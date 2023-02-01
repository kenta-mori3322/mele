#!/bin/sh
# new-testnet - example make call to create a new set of validator nodes in AWS
# WARNING: Run it from the current directory - it uses relative paths to ship the binary

if [ $# -ne 2 ]; then
  echo "Usage: ./new-testnet.sh <region-limit> <number-of-nodes-per-availability-zone>"
  exit 1
fi

if [ -f "keys.sh" ]; then
    source keys.sh
else
    echo "Error: keys.sh doesn't exist"
    exit 1
fi

set -eux

# The testnet name is the same on all nodes
export REGION_LIMIT=$1
export SERVERS=$2

# Build the AWS validator nodes and extract the genesis.json and config.toml from one of them
make validators-bootstrap
