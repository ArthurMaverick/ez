resource "aws_api_gateway_stage" "stage" {
  deployment_id = var.deployment_id
  rest_api_id   = var.rest_api_id
  stage_name    = var.stage_name
}
