terraform {
  required_providers {
    ns = {
      source  = "netskope/ns"
      version = "0.7.123"
    }
  }
}

provider "ns" {
  # Configuration options
}