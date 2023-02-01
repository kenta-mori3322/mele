#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/meled/${BINARY:-meled-manager}
ID=${ID:-0}
LOG=${LOG:-meled.log}
LOG_LEVEL=""

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'meled' E.g.: -e BINARY=meled_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Setup upgrade manager
##
export DAEMON_NAME=meled
export DAEMON_HOME="/meled/node${ID}/meled"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=on
export DAEMON_RESTART_AFTER_UPGRADE=on
export CONFIG_FILE_NAME="config.toml"
export GENESIS_FILE_NAME="genesis.json"

mkdir -p ${DAEMON_HOME}/upgrade_manager/genesis/bin
mkdir -p ${DAEMON_HOME}/upgrade_manager/upgrades
mkdir -p ${DAEMON_HOME}/config

cp /meled/${CONFIG_FILE_NAME} ${DAEMON_HOME}/config/${CONFIG_FILE_NAME}
cp /meled/${GENESIS_FILE_NAME} ${DAEMON_HOME}/config/${GENESIS_FILE_NAME}

sed -i "s/validator4/docker-validator${ID}/g" ${DAEMON_HOME}/config/${CONFIG_FILE_NAME}

if ! [ -f "${DAEMON_HOME}/upgrade_manager/genesis/bin/${DAEMON_NAME}" ]; then
	cp /meled/${DAEMON_NAME} ${DAEMON_HOME}/upgrade_manager/genesis/bin/${DAEMON_NAME}
fi

##
## Run binary with all parameters
##
export MELEDHOME="/meled/node${ID}/meled"

if [ -d "$(dirname "${MELEDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${MELEDHOME}" "$@" | tee "${MELEDHOME}/${LOG}"
else
  "${BINARY}" --home "${MELEDHOME}" "$@"
fi