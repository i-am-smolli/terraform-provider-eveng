terraform {
  required_providers {
    eveng = {
      source = "i-am-smolli/eveng"
    }
  }
}

provider "eveng" {}

resource "eveng_folder" "example" {
  path = "/example"
}
