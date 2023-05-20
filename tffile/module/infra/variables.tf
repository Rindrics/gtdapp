variable "environment" {
  type        = string
  description = "Name of environment such as 'production'"
}

variable "host_ingress_port_http" {
  type        = number
  description = "HTTP port number to use ingress"
}

variable "host_ingress_port_ssh" {
  type        = number
  description = "SSH port number to use ingress"
}
