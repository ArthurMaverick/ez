output "rest_api_id" {
  value = aws_api_gateway_rest_api.rest_api.id
}

output "rest_api_name" {
  value = aws_api_gateway_rest_api.rest_api.name
}

output "root_resource_id" {
  value = aws_api_gateway_rest_api.rest_api.root_resource_id
}