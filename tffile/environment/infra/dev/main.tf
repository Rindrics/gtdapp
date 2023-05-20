terraform {
  required_providers {
    kind = {
      source  = "justenwalker/kind"
      version = "0.17.0"
    }
  }
}

module "kind_dev_cluster" {
  source                 = "../../../module/infra/"
  environment            = "dev"
  host_ingress_port_http = 8080
  host_ingress_port_ssh  = 8443
}
