terraform {
  required_providers {
    devto = {
      source = "devto/devto"
    }
  }
}

provider "devto" {
  api_key = var.api_key
 # api_url = "http://localhost:6666"
}

variable "api_key" {
  type = string
}

resource "devto_article" "one"{
  title = "dp-test-two"
  markdown_body = <<EOT
## Subtitle
This is the body of the post

EOT
  published = false
}