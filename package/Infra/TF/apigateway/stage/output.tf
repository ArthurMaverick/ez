output "stage_arn" {
  value =  aws_api_gateway_stage.stage.arn
}

output "stage_id" {
  value =  aws_api_gateway_stage.stage.id
}

output "stage_name" {
  value =  aws_api_gateway_stage.stage.stage_name
}