resource "digitalocean_droplet" "web_server" {
  image  = "docker-20-04"  # Use an image with Docker pre-installed
  name   = "nicklesseos.com"
  region = var.region
  size   = "s-1vcpu-1gb"

  # SSH key for accessing the Droplet
  ssh_keys = [var.ssh_fingerprint]

  # Provisivar.regioncker image
  provisioner "remote-exec" {
    inline = [
      "docker pull ${var.docker_image}:${var.docker_image_tag}",
      "docker run -d -p 80:3000 ${var.docker_image}:${var.docker_image_tag}"
    ]
  }

  connection {
    type        = "ssh"
    user        = "root"
    private_key = var.pvt_key
    host        = self.ipv4_address
  }

}