resource "aws_api_gateway_rest_api" "rest_api" {
  name        = var.rest_api_name
  description = var.description

  endpoint_configuration {
    types            = [var.endpoint]
    vpc_endpoint_ids = var.vpc_endpoint_id == "" && var.endpoint != "PRIVATE" ? null : var.vpc_endpoint_id
  }
}
