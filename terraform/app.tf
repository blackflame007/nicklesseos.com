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

    env {
      key   = "JWT_SECRET"
      value = var.jwt_secret
    }

    env {
      key   = "DB_URL"
      value = var.db_url
    }

    env {
      key   = "DB_AUTH_TOKEN"
      value = var.db_auth_token
    }

    // Define the domain
    domain {
      name     = digitalocean_domain.default.name
      type     = "PRIMARY"
      wildcard = false
    }

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
