variable "do_token" {}

variable "region" {
  description = "The region to deploy to"
  default     = "sfo3"

}

variable "do_space_name" {
  description = "The name of your DigitalOcean Space"
  default     = "nicklesseos-com-space"
}


variable "do_spaces_access_key" {
  description = "DigitalOcean Spaces Access Key"
  type        = string
}

variable "do_spaces_secret_key" {
  description = "DigitalOcean Spaces Secret Key"
  type        = string
}

variable "google_client_id" {
  description = "Google Client ID"
  type        = string
}

variable "google_client_secret" {
  description = "Google Client Secret"
  type        = string
}

variable "google_oauth_redirect_url" {
  description = "Google OAuth Redirect URL"
  type        = string
}

variable "session_secret" {
  description = "Session Secret"
  type        = string
}
