resource "digitalocean_app" "nicklesseos-com" {
  spec {
    name   = "nicklesseos-com"
    region = var.region

    env {
      key   = "AWS_ACCESS_KEY_ID"
      value = var.do_spaces_access_key
    }

    env {
      key   = "AWS_SECRET_ACCESS_KEY"
      value = var.do_spaces_secret_key
    }

    env {
      key   = "DO_SPACE_NAME"
      value = var.do_space_name
    }

    env {
      key   = "GOOGLE_CLIENT_ID"
      value = var.google_client_id
    }

    env {
      key   = "GOOGLE_CLIENT_SECRET"
      value = var.google_client_secret
    }

    env {
      key   = "GOOGLE_OAUTH_REDIRECT_URL"
      value = var.google_oauth_redirect_url
    }

    env {
      key   = "SESSION_KEY"
      value = var.session_secret
    }

    // Define the domain
    domain {
      name = digitalocean_domain.default.name
      type = "PRIMARY"
    }

    // Add www Alias
    # domain {
    #   name = "www.${digitalocean_domain.default.name}"
    #   type = "ALIAS"
    # }

    service {
      name            = "nicklesseos-com"
      http_port       = 3000
      dockerfile_path = "./Dockerfile"

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
