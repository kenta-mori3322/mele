sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
sudo apt install make

wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.15.3
source ~/.profile

# Install executables
cd /home/ubuntu/mele-validator-dev
make install

# Initialize the validator, where "test-mel" is a moniker name
meled init validator --chain-id test

# Set up your keys, where "mele-test-wallet" is a wallet name
#meled keys add mele-test-wallet --keyring-backend test
#meled keys add mele-test-wallet --recover --keyring-backend test  answer should be "auction cruel base speak drip orphan develop egg dream retreat theory trash fan circle audit"

# Create the upgrade manager directory structure
mkdir -p ~/.meled/upgrade_manager/genesis/bin && mkdir -p ~/.meled/upgrade_manager/upgrades

# Copy the genesis binary to the upgrade manager
cp $(which meled) ~/.meled/upgrade_manager/genesis/bin

# Fetch the genesis file
#sudo apt install jq -y
#curl http://3.19.27.59:26657/genesis? | jq ".result.genesis" > ~/.meled/config/genesis.json

# Copy the meled-manager binary to /usr/bin
cp $(which meled-manager) /usr/bin

# Create the service file with the following content
# sudo nano /etc/systemd/system/meled.service
sudo echo "[Unit]
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
ExecStart=/usr/bin/meled-manager start --pruning=\"nothing\" --rpc.laddr \"tcp://0.0.0.0:26657\"
StandardOutput=file:/var/log/meled/meled.log
StandardError=file:/var/log/meled/meled_error.log
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" > "/etc/systemd/system/meled.service"

# Create log files for meled
make log-files
sudo systemctl enable meled
#sed -i '184s/.*/seeds=\"2ef75d8cd40586f5081206bf5dfe7fc414c5607a@127.0.0.1:26656\"/' ~/.meled/config/config.toml
sudo systemctl start meled

# Create the validator, where "mele-test-wallet" is a wallet name and the "test-mel" is a moniker name
#meled tx mstaking create-validator --from mele-test-wallet --moniker test-mel --pubkey $(meled tendermint show-validator) --chain-id mainnet --keyring-backend test --amount 1umelg --commission-max-change-rate 0.01 --commission-max-rate 0.2 --commission-rate 0.1 --min-self-delegation 1

#meled q mstaking validator $(meled keys show mele-test-wallet --keyring-backend test --bech val -a) --chain-id mainnet
