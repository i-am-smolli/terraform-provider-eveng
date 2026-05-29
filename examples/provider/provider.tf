terraform {
  required_providers {
    eveng = {
      source  = "i-am-smolli/eveng"
      version = "0.1.8"
    }
  }
}

provider "eveng" {
  host     = "http://localhost"
  username = "admin"
  password = "eve"
}
