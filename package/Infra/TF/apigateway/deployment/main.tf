resource "aws_api_gateway_deployment" "deployment" {
  rest_api_id = var.rest_api_id
  # stage_name = var.stage_id

  triggers = {
    redeployment = sha1(jsonencode([
      var.api_gateway_resource_ids
    ]))
  }

  lifecycle {
    create_before_destroy = true
  }
}