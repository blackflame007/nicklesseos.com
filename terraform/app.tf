resource "digitalocean_app" "nicklesseos-com" {
  spec {
    name   = "nicklesseos-com"
    region = var.region

    // Define the domain
    domain {
      name = digitalocean_domain.default.name
      type = "PRIMARY"
    }

    # domain {
    #   name = "www.nicklesseos.com"
    #   type = "ALIAS"
    # }

    service {
      name      = "nicklesseos-com"
      http_port = 3000

      github {
        repo           = "blackflame007/nicklesseos.com"
        branch         = "main"
        deploy_on_push = true
      }

      health_check {
        http_path             = "/health"
        initial_delay_seconds = 30
        failure_threshold     = 4
      }
    }
  }
}
