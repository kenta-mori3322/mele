variable "name" {
  description = "The monitor name, e.g monitor"
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

variable "ssh_private_file" {
  description = "SSH private key file to be used to connect to the nodes"
  type = string
}

variable "ssh_public_file" {
  description = "SSH public key file to be used on the nodes"
  type = string
}