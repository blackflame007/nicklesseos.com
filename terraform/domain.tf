# Create a new domain
resource "digitalocean_domain" "default" {
  name = "nicklesseos.com"
}

# Create a CNAME record for the domain
resource "digitalocean_record" "www" {
  domain = digitalocean_domain.default.name
  type   = "CNAME"
  name   = "www"
  value  = "@"
}
