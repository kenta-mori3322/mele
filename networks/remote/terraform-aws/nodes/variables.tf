variable "name" {
  description = "The testnet name, e.g testnet"
}

variable "image_name" {
  description = "Image name"
  default = "ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"
}

variable "instance_type" {
  description = "The instance size to use"
  default = "t3.large"
}

variable "region" {
  description = "AWS region to use"
}

variable "multiplier" {
  description = "Multiplier for node identification"
}

variable "execute" {
  description = "Set to false to disable the module"
  default = true
}

variable "SERVERS" {
  description = "Number of servers in an availability zone"
  default = "1"
}

variable "ssh_private_file" {
  description = "SSH private key file to be used to connect to the nodes"
  type = string
}

variable "ssh_public_file" {
  description = "SSH public key file to be used on the nodes"
  type = string
}