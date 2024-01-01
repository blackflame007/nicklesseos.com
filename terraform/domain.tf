# Create a new domain
resource "digitalocean_domain" "default" {
  name = "nicklesseos.com"
}

# Create a www domain
resource "digitalocean_domain" "wildcard" {
  name = "*.nicklesseos.com"
}
