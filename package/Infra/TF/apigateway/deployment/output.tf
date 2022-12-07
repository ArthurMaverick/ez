output "aws_api_gateway_deployment_id" {
  value = aws_api_gateway_deployment.deployment.id
}

output "aws_api_gateway_deployment_triggers" {
  value = aws_api_gateway_deployment.deployment.triggers
}