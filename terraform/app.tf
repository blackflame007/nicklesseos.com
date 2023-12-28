resource "digitalocean_app" "golang-docker" {
  spec {
    name   = "golang-sample"
    region = var.region

    service {
      name      = "go-service"
      http_port = 3000

      github {
        repo        = "blackflame007/nicklesseos.com"
        branch      = "main"
        deploy_on_push = true
      }

      health_check {
        http_path              = "/health"
        initial_delay_seconds  = 30
        failure_threshold      = 4
      }
    }
  }
}
