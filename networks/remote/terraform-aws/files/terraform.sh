#!/bin/bash
# Script to initialize a testnet settings on a server

until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 1
done

# Usage: terraform.sh <testnet_name> <testnet_region_number> <testnet_node_number> <influxdb_password>

# Add meled node number for remote identification
REGION_ID="$(($2 + 1))"
REGION_NODE_ID="$(($3 + 1))"
ID="$((${REGION_ID} * 100 + ${REGION_NODE_ID}))"
INFLUX_PASSWORD=$4

echo "$ID" > /etc/nodeid
echo "$INFLUX_PASSWORD" > /etc/influx_password

# Install jq JSON helper and mele binaries
echo "deb [trusted=yes] http://deb.melechain.dev/ubuntu ./" | sudo tee -a /etc/apt/sources.list.d/mele.list > /dev/null

sudo apt update
sudo apt install -y mele

sudo apt update
sudo apt install -y jq

# Install InfluxDB
sudo curl -sL https://repos.influxdata.com/influxdb.key | sudo apt-key add -
sudo echo "deb https://repos.influxdata.com/ubuntu bionic stable" | sudo tee /etc/apt/sources.list.d/influxdb.list
sudo apt update
sudo apt install -y influxdb
sudo systemctl enable --now influxdb
sed -i 's#\# auth-enabled = false#auth-enabled = true#g' /etc/influxdb/influxdb.conf
sudo systemctl restart influxdb
sleep 10
curl -XPOST "http://localhost:8086/query" --data-urlencode "q=CREATE USER admin WITH PASSWORD '${INFLUX_PASSWORD}' WITH ALL PRIVILEGES"
curl -XPOST "http://localhost:8086/query" -u "admin:${INFLUX_PASSWORD}" --data-urlencode "q=CREATE DATABASE vcf"
curl -XPOST "http://localhost:8086/query" -u "admin:${INFLUX_PASSWORD}" --data-urlencode "q=CREATE DATABASE telegraf"

# Install telegraf
sudo apt-get update
sudo apt-get -y install telegraf
sed -i "s#^  \# username = \"telegraf\"#  username = \"admin\"#g" /etc/telegraf/telegraf.conf
sed -i "s#^  \# password = \"metricsmetricsmetricsmetrics\"#  password = \"${INFLUX_PASSWORD}\"#g" /etc/telegraf/telegraf.conf
sudo systemctl enable --now telegraf
sleep 10
sudo systemctl restart telegraf
