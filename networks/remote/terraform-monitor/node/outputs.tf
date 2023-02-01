
// The cluster name
output "name" {
  value = var.name
}

// The list of cluster instance IDs
output "instances" {
  value = [ aws_instance.node.*.id ]
}

// The list of cluster instance public IPs
output "public_ip" {
  value = [ aws_eip.node.*.public_ip ]
}
