#!/usr/bin/env bash

for ((i = 0 ; i <= 3 ; i++)); do
    meled keys delete "validator$i" --keyring-backend test -y
    cat build/node$i/meled/key_seed.json | jq '.secret' -r | meled keys add "validator$i" --keyring-backend test --recover
done