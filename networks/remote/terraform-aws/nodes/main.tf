provider "aws" {
  region = var.region
  version = "~> 2.70"
}

resource "aws_key_pair" "testnets" {
  count = var.execute ? 1 : 0
  key_name   = "testnets-${var.name}"
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
  count = var.execute ? 1 : 0
  name = var.name
  description = "Automated security group for performance testing testnets"
  tags = {
    Name = "testnets-${var.name}"
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 26656
    to_port = 26657
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 1317
    to_port = 1317
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 9090
    to_port = 9090
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 8086
    to_port = 8086
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

resource "random_password" "password" {
  count = var.execute ? var.SERVERS : 0
  length = 16
  special = true
  override_special = "_%@"
}

resource "aws_instance" "node" {
  count = var.execute ? var.SERVERS : 0
  ami = data.aws_ami.linux.image_id
  instance_type = var.instance_type
  key_name = aws_key_pair.testnets[0].key_name
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
      "sudo /tmp/terraform.sh ${var.name} ${var.multiplier} ${count.index} ${element(random_password.password.*.result, count.index)} ",
    ]
  }
}

resource "aws_eip" "node" {
  count     = var.execute ? var.SERVERS : 0
  instance = element(aws_instance.node.*.id, count.index)
  vpc      = true
}