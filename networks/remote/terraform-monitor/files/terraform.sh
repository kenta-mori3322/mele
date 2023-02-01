#!/bin/bash
# Script to initialize a testnet settings on a server

until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 1
done

# Usage: terraform.sh

