resource "digitalocean_app" "golang-docker" {
  spec {
    name   = "golang-sample"
    region = "ams"

    service {
      name               = "go-service"
      http_port          = 3000

    image {
        registry_type = "DOCKER_HUB"
        registry      = "blackflame007"
        repository    = "nicklesseos.com"
        tag           = "latest"
      }

    health_check {
        http_path     = "/health"
        initial_delay_seconds = 30
        failure_threshold = 4
    }
    }
  }
}