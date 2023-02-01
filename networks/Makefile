# Name of the testnet. Used in chain-id.
TESTNET_NAME?=remotenet

# Name of the servers grouped together for management purposes. Used in tagging the servers in the cloud.
CLUSTER_NAME?=$(TESTNET_NAME)

# Number of servers to put in one availability zone in AWS.
SERVERS?=1

# Number of regions to use in AWS. One region usually contains 2-3 availability zones.
REGION_LIMIT?=1

# Region to deploy VPC and application in AWS
REGION ?= eu-central-1

# Local binaries path
DAEMON_BINARY?=$(CURDIR)/../build/meled
CLI_BINARY?=$(CURDIR)/../build/melecli

# Path to the genesis.json and config.toml files to deploy on full nodes.
GENESISFILE?=$(CURDIR)/../build/genesis.json
CONFIGFILE?=$(CURDIR)/../build/config.toml

install-ansible-requirements:
	cd remote/ansible && ansible-galaxy install -r requirements.yml

validators-bootstrap:
	# Make sure you have AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY or your IAM roles set for AWS API access.
	@if ! [ -f $(SSH_PUBLIC_FILE) ]; then echo "Error: keys file doesn't exist" && exit 1 ; fi

	cd remote/terraform-aws && \
		terraform init && \
		(terraform workspace new "$(CLUSTER_NAME)" || terraform workspace select "$(CLUSTER_NAME)") && \
	  	terraform apply \
			-auto-approve \
			-var SSH_PUBLIC_FILE="$(SSH_PUBLIC_FILE)" \
			-var SSH_PRIVATE_FILE="$(SSH_PRIVATE_FILE)" \
			-var TESTNET_NAME="$(CLUSTER_NAME)" \
			-var SERVERS="$(SERVERS)" \
			-var REGION_LIMIT="$(REGION_LIMIT)"

	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
		-i inventory/ec2.py \
		-l "tag_Environment_$(CLUSTER_NAME)" \
		-u ubuntu \
		-b \
		-e TESTNET_NAME="$(TESTNET_NAME)" \
		--private-key="$(SSH_PRIVATE_FILE)" \
		setup-validators.yml

	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
		-i inventory/ec2.py \
		-l "tag_Environment_$(CLUSTER_NAME)" \
		-u ubuntu \
		--private-key="$(SSH_PRIVATE_FILE)" \
		-b \
		start.yml

	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
		-i inventory/ec2.py \
		-l "tag_Environment_$(CLUSTER_NAME)" \
		-u ubuntu \
		--private-key="$(SSH_PRIVATE_FILE)" \
		-b \
		mele-monitor.yml

	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False OBJC_DISABLE_INITIALIZE_FORK_SAFETY=YES ansible-playbook \
		-i inventory/ec2.py \
		-l "tag_Environment_$(CLUSTER_NAME)" \
		-u ubuntu \
		--private-key="$(SSH_PRIVATE_FILE)" \
		-b \
		setup-prometheus.yml

validators-start:
	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
    	-i inventory/ec2.py \
    	-l "tag_Environment_$(CLUSTER_NAME)" \
    	-u ubuntu \
    	--private-key="$(SSH_PRIVATE_FILE)" \
    	-b \
    	start.yml

validators-stop:
	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
    	-i inventory/ec2.py \
    	-l "tag_Environment_$(CLUSTER_NAME)" \
    	-u ubuntu \
    	--private-key="$(SSH_PRIVATE_FILE)" \
    	-b \
    	stop.yml

validators-status:
	cd remote/ansible && ansible-playbook \
		-i inventory/ec2.py \
		-l "tag_Environment_$(CLUSTER_NAME)" \
		status.yml

ship-local-binaries:
	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
    	-i inventory/ec2.py \
    	-l "tag_Environment_$(CLUSTER_NAME)" \
    	-u ubuntu \
    	-e DAEMON_BINARY=$(DAEMON_BINARY) \
    	-e CLI_BINARY=$(CLI_BINARY) \
    	--private-key="$(SSH_PRIVATE_FILE)" \
    	-b \
    	ship-local-binaries.yml

validators-destroy:
	cd remote/terraform-aws && \
		terraform workspace select "$(CLUSTER_NAME)" && \
		terraform destroy \
			-force \
			-var SSH_PUBLIC_FILE="$(SSH_PUBLIC_FILE)" \
			-var SSH_PRIVATE_FILE="$(SSH_PRIVATE_FILE)" && \
		terraform workspace select default && \
		terraform workspace delete "$(CLUSTER_NAME)"
	rm -rf remote/ansible/keys/ remote/ansible/files/

monitor-bootstrap:
	# Make sure you have AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY or your IAM roles set for AWS API access.
	@if ! [ -f $(SSH_PUBLIC_FILE) ]; then echo "Error: keys file doesn't exist" && exit 1 ; fi

	cd remote/terraform-monitor && \
		terraform init && \
		(terraform workspace new "$(MONITOR_NAME)" || terraform workspace select "$(MONITOR_NAME)") && \
	  	terraform apply \
			-auto-approve \
			-var SSH_PUBLIC_FILE="$(SSH_PUBLIC_FILE)" \
			-var SSH_PRIVATE_FILE="$(SSH_PRIVATE_FILE)" \

	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
        -i inventory/ec2.py \
        -l "tag_Environment_$(MONITOR_NAME)" \
        -u ubuntu \
        -e CLUSTER_NAME=$(CLUSTER_NAME) \
        --private-key="$(SSH_PRIVATE_FILE)" \
        -b \
        setup-grafana.yml

monitor-destroy:
	cd remote/terraform-monitor && \
		terraform workspace select "$(MONITOR_NAME)" && \
		terraform destroy \
			-force \
			-var SSH_PUBLIC_FILE="$(SSH_PUBLIC_FILE)" \
			-var SSH_PRIVATE_FILE="$(SSH_PRIVATE_FILE)" && \
		terraform workspace select default && \
		terraform workspace delete "$(MONITOR_NAME)"

extract-config:
	# Make sure you have AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY or your IAM roles set for AWS API access.
	@if ! [ -f $(SSH_PUBLIC_FILE) ]; then echo "Error: keys file doesn't exist" && exit 1 ; fi
	cd remote/ansible && \
		ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
			-i inventory/ec2.py \
			-l "tag_Environment_$(CLUSTER_NAME)" \
			-b \
			-u ubuntu \
			-e GENESISFILE="$(GENESISFILE)" \
			-e CONFIGFILE="$(CONFIGFILE)" \
			extract-config.yml
