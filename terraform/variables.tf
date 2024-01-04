variable "do_token" {}
variable "ssh_fingerprint" {
  description = "The SSH fingerprint of your public key"
}
variable "pvt_key" {
  description = "The path to your private key"
}

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
