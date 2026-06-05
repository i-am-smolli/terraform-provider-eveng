terraform {
  required_providers {
    eveng = {
      source = "i-am-smolli/eveng"
    }
  }
}

provider "eveng" {}

resource "eveng_lab" "example" {
  name   = "StartNodesExample"
  author = "Corentin"
}

resource "eveng_node" "vpc" {
  lab_path = eveng_lab.example.path
  name     = "vpc1"
  template = "vpcs"
  type     = "vpcs"
}

resource "eveng_start_nodes" "start" {
  lab_path   = eveng_lab.example.path
  depends_on = [eveng_node.vpc]
}
