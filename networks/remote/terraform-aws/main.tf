#Terraform Configuration

variable "REGIONS" {
  description = "AWS Regions"
  type = list(string)
  default = ["eu-central-1", "eu-west-1", "eu-west-2", "us-east-2", "us-west-1", "us-west-2", "ap-south-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ca-central-1", "sa-east-1"]
}

variable "TESTNET_NAME" {
  description = "Name of the testnet"
  default = "remotenet"
}

variable "REGION_LIMIT" {
  description = "Number of regions to populate"
  default = "1"
}

variable "SERVERS" {
  description = "Number of servers in a region"
  default = "1"
}

variable "SSH_PRIVATE_FILE" {
  description = "SSH private key file to be used to connect to the nodes"
  type = string
}

variable "SSH_PUBLIC_FILE" {
  description = "SSH public key file to be used on the nodes"
  type = string
}

variable "image" {
  description = "AWS image name"
  default = "ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"
}

variable "instance_type" {
  description = "AWS instance type"
  default = "t3.large"
}

module "nodes-0" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,0)}"
  multiplier       = "0"
  execute          = "${var.REGION_LIMIT > 0}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-1" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,1)}"
  multiplier       = "1"
  execute          = "${var.REGION_LIMIT > 1}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-2" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,2)}"
  multiplier       = "2"
  execute          = "${var.REGION_LIMIT > 2}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-3" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,3)}"
  multiplier       = "3"
  execute          = "${var.REGION_LIMIT > 3}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-4" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,4)}"
  multiplier       = "4"
  execute          = "${var.REGION_LIMIT > 4}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-5" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,5)}"
  multiplier       = "5"
  execute          = "${var.REGION_LIMIT > 5}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-6" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,6)}"
  multiplier       = "6"
  execute          = "${var.REGION_LIMIT > 6}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-7" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,7)}"
  multiplier       = "7"
  execute          = "${var.REGION_LIMIT > 7}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-8" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,8)}"
  multiplier       = "8"
  execute          = "${var.REGION_LIMIT > 8}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-9" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,9)}"
  multiplier       = "9"
  execute          = "${var.REGION_LIMIT > 9}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-10" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,10)}"
  multiplier       = "10"
  execute          = "${var.REGION_LIMIT > 10}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-11" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,11)}"
  multiplier       = "11"
  execute          = "${var.REGION_LIMIT > 11}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

module "nodes-12" {
  source           = "./nodes"
  name             = "${var.TESTNET_NAME}"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "${element(var.REGIONS,12)}"
  multiplier       = "12"
  execute          = "${var.REGION_LIMIT > 12}"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
  SERVERS          = "${var.SERVERS}"
}

output "public_ips" {
  value = "${concat(
		module.nodes-0.public_ips,
		module.nodes-1.public_ips,
		module.nodes-2.public_ips,
		module.nodes-3.public_ips,
		module.nodes-4.public_ips,
		module.nodes-5.public_ips,
		module.nodes-6.public_ips,
		module.nodes-7.public_ips,
		module.nodes-8.public_ips,
		module.nodes-9.public_ips,
		module.nodes-10.public_ips,
		module.nodes-11.public_ips,
		module.nodes-12.public_ips
		)}"
}

output "influx_db_passwords" {
  value = "${concat(
        module.nodes-0.influx_db_passwords,
		module.nodes-1.influx_db_passwords,
		module.nodes-2.influx_db_passwords,
		module.nodes-3.influx_db_passwords,
		module.nodes-4.influx_db_passwords,
		module.nodes-5.influx_db_passwords,
		module.nodes-6.influx_db_passwords,
		module.nodes-7.influx_db_passwords,
		module.nodes-8.influx_db_passwords,
		module.nodes-9.influx_db_passwords,
		module.nodes-10.influx_db_passwords,
		module.nodes-11.influx_db_passwords,
		module.nodes-12.influx_db_passwords
  )}"
}