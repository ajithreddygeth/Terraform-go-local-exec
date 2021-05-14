resource "null_resource" "go" {
  provisioner "local-exec" {
    environment = {
      AWS_REGION           = "${var.aws_region}"
      AWS_ACCOUNT     = "${var.aws_account}"
    }
    command = "./utils/bin/utils"
  }
}