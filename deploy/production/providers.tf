terraform {
  backend "s3" {
    bucket  = "di-terraform"
    key     = "di-notebook-connect"
    encrypt = true
  }
}