provider "aws" {
  region = var.region
  version = "~> 2.70"
}

resource "aws_key_pair" "monitor" {
  count = 1
  key_name   = "monitor"
  public_key = file(var.ssh_public_file)
}

data "aws_ami" "linux" {
  most_recent = true

  filter {
    name   = "name"
    values = [var.image_name]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"]
}

data "aws_availability_zones" "zones" {
  state = "available"
}

resource "aws_security_group" "secgroup" {
  count = 1
  name = var.name
  description = "Security group for testnet monitor"
  tags = {
    Name = "secgroup-${var.name}"
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 3000
    to_port = 3000
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "node" {
  count = 1
  ami = data.aws_ami.linux.image_id
  instance_type = var.instance_type
  key_name = aws_key_pair.monitor[0].key_name
  associate_public_ip_address = true
  security_groups = [ aws_security_group.secgroup[0].name ]
  availability_zone = element(data.aws_availability_zones.zones.names, count.index)

  tags = {
    Environment = var.name
    Name = "${var.name}-${element(data.aws_availability_zones.zones.names,count.index)}"
  }

  volume_tags = {
    Environment = var.name
    Name = "${var.name}-${element(data.aws_availability_zones.zones.names,count.index)}-VOLUME"
  }

  root_block_device {
    volume_size = 40
  }

  connection {
    host = self.public_ip
    user = "ubuntu"
    private_key = file(var.ssh_private_file)
    timeout = "600s"
  }

  provisioner "file" {
    source = "files/terraform.sh"
    destination = "/tmp/terraform.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/terraform.sh",
      "sudo /tmp/terraform.sh",
    ]
  }
}

resource "aws_eip" "node" {
  count     = 1
  instance = element(aws_instance.node.*.id, count.index)
  vpc      = true
}