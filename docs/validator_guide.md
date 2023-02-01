# Mele Validator Guide

## Infrastructure

Recommended configuration:

- Number of CPUs: **2**
- Memory: **8GB**
- Disk: **250GB SSD**
- OS: **Ubuntu 18.04 LTS**
- Allow all incoming connections from TCP port 26656 and 26657
- Static IP address

The recommended configuration from AWS is the equivalent of a *t3.large* machine with **250GB** EBS attached storage.

## Prerequisites

Update the system and install dependencies:
```bash
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```

Install Golang:
```bash
# Install latest go version https://golang.org/doc/install
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.15.3
source ~/.profile
```

To verify that Golang is installed:
```bash
go version
# Should return go version go1.15.3 linux/amd64
```

## Running a Validator Node

Install the executables
```bash
git clone https://github.com/melechain/mele.git
cd mele

# install executables
make install
```

Initialize the validator
```bash
# Replace <your-moniker> with the publicly viewable name for your validator.
meled init --chain-id mainnet <your-moniker>
```

Set up your keys
```bash
# Create a wallet for your node. <your-wallet-name> is just a human-readable name you can use to remember your wallet. It can be the same or different than your moniker.
meled keys add <your_wallet_name> --keyring-backend test

# If you have a wallet with a balance assigned to it, you need to import it by. You will be prompted to enter your bip39 mnemonic
meled keys add <your_wallet_name> --recover --keyring-backend test
```

Create the upgrade manager directory structure
```bash
mkdir -p ~/.meled/upgrade_manager/genesis/bin && mkdir -p ~/.meled/upgrade_manager/upgrades
```

Copy the genesis binary to the upgrade manager
```bash
cp $(which meled) ~/.meled/upgrade_manager/genesis/bin
```

Verify that the binary has been copied
```bash
ls -las  ~/.meled/upgrade_manager/genesis/bin

# Should return the meled binary in the correct location
```

Fetch the genesis file
```bash
# Copy the genesis file to the meled directory (i.e. TBD = http://3.19.27.59:26657/genesis? )
curl <TBD> | jq ".result.genesis" > ~/.meled/config/genesis.json

# The genesis file will be available when the mainnet is up and running
```

Copy the meled-manager binary to /usr/bin
```bash
sudo cp $(which meled-manager) /usr/bin
```

Create the service file with the following content
```bash
sudo nano /etc/systemd/system/meled.service
```

```
[Unit]
Description=meled
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=ubuntu
Group=ubuntu
Environment=DAEMON_NAME=meled
Environment=DAEMON_HOME=/home/ubuntu/.meled
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/meled-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657" --state-sync.snapshot-interval 100 --state-sync.snapshot-keep-recent 2
StandardOutput=file:/var/log/meled/meled.log
StandardError=file:/var/log/meled/meled_error.log
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

*If you are not logged in as the `ubuntu` user and/or if your home directory is not `/home/ubuntu`, please change the `User`, `Group`, `Environment`, and `ExecStart` variables in the service config above appropriately.*

Start the Daemon service
```bash
# Create log files for meled
make log-files

# Enable the meled service
sudo systemctl enable meled

# Add seed nodes to ~/.meled/config/config.toml, replace line 184 with
# you can get the seed id from `http://3.19.27.59:26657/status?`
# Seeds will be available when the mainnet is up and running
seeds = "26ddec184a4e489074215a96135bf6b859bd8c9b@3.19.27.59:26656" 


# Start the node
sudo systemctl start meled
```

To check on the status of the node use:
```
meled status
# or
sudo journalctl -u meled -f
```

To view the logs use:
```
# Standard output of meled
tail -f /var/log/meled/meled.log

# Standard error of meled
tail -f /var/log/meled/meled_error.log
```

## Applying for being a validator

Create the validator
```bash
# Be sure to replace <your-wallet-name>, <your-moniker>, <amount_of_umelg>
meled tx mstaking create-validator --from <your-wallet-name> --moniker <your-moniker> --pubkey $(meled tendermint show-validator) --chain-id mainnet --keyring-backend test --amount <amount_of_umelg>umelg
```

Verify the node is in the validator list
```bash
# Check all validators
meled q mstaking validators

# Check current validator
# Be sure to replace <your-wallet-name>
meled q mstaking validator $(meled keys show <your-wallet-name> --keyring-backend test --bech val -a) --chain-id mainnet
```

Recovering From a Slashing Infraction

First, you need to verify the state of your validator by running:
```bash
# Check current validator
# Be sure to replace <your-wallet-name>
meled q mstaking validator $(meled keys show <your-wallet-name> --keyring-backend test --bech val -a) --chain-id mainnet
```

The response would be similar to the following:
```
  commission:
    commission_rates:
      max_change_rate: "0.010000000000000000"
      max_rate: "0.200000000000000000"
      rate: "0.100000000000000000"
    update_time: "2021-07-13T12:52:17.660400Z"
  consensus_pubkey:
    '@type': /cosmos.crypto.ed25519.PubKey
    key: 3CCzld9echQOp9Hk0ydu6jbY3Mrcw2lt/V0AiVaTWOE=
  delegator_shares: "90000000000000.000000000000000000"
  description:
    details: ""
    identity: ""
    moniker: validator
    security_contact: ""
    website: ""
  jailed: false
  min_self_delegation: "1"
  operator_address: melevaloper17zc58s96rxj79jtqqsnzt3wtx3tern6ah30k4e
  status: BOND_STATUS_BONDED
  tokens: "90000000000000"
  unbonding_height: "0"
  unbonding_time: "1970-01-01T00:00:00Z"
```

As you can see the jailed status is set to true. You need to send an unjail transaction if you want to continue being a validator.

To check when is the earliest time the validator can be unjailed run:
```bash
meled q mslashing signing-info $(meled tendermint show-validator) --chain-id mainnet
```

The response will return the jailed_until parameter in the UTC time zone:
```
address: melevalcons163q2272thhxq70rfw7g6sldeuce3hmgrysshy2
index_offset: "12"
jailed_until: "1970-01-01T00:00:00Z"
missed_blocks_counter: "0"
start_height: "0"
tombstoned: false
```

To unjail your validator run:
```
# Be sure to replace <your-wallet-name>
meled tx mslashing unjail --from <your-wallet-name> --keyring-backend test --chain-id mainnet
```

Stopping a Validator Node

To gracefully shutdown a validator node which is in the active validator set, the operator must first unbond their tokens before being able to shut down the node and withdraw their stake.
```bash
# Be sure to replace <your-wallet-name>
meled tx mstaking unbond --from <your-wallet-name> --chain-id mainnet --keyring-backend test
```

After running the unbonding transaction, you need to check the length of the unbonding time by running:
```bash
meled q mstaking params
```

After that period of time, your stake will be returned to your account. During the unbounding period, you can be slashed for any infraction that happened before the unbonding transaction.
As soon as you run the unbond transaction, you are free to shut down your validator node.
```bash
sudo systemctl stop meled
```
