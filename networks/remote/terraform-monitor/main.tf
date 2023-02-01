#Terraform Configuration

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

module "node" {
  source           = "./node"
  name             = "monitor"
  image_name       = "${var.image}"
  instance_type    = "${var.instance_type}"
  region           = "eu-central-1"
  ssh_public_file  = "${var.SSH_PUBLIC_FILE}"
  ssh_private_file = "${var.SSH_PRIVATE_FILE}"
}

output "public_ip" {
  value = module.node.public_ip
}
