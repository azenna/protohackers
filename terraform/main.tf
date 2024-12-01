terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "zenna_token" {
    type = string
}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = var.zenna_token
}

# Create a new SSH key
resource "digitalocean_ssh_key" "default" {
  name       = "tf-zennas-key"
  public_key = file("~/.ssh/id_ed25519.pub")
}

resource "digitalocean_project" "protohackers" {
  name        = "protohackers"
  description = "Writing servers exercises"
}

# Create a new Droplet using the SSH key
resource "digitalocean_droplet" "zen-droplet" {
  image    = "ubuntu-24-04-x64"
  name     = "zen-droplet"
  region   = "nyc3"
  size     = "s-1vcpu-1gb"
  ssh_keys = [digitalocean_ssh_key.default.fingerprint]
}
