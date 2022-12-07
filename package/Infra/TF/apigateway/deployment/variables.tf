variable "rest_api_id" {
  type = string
  description = "[required] - rest api id"
}

variable "api_gateway_resource_ids" {
  type = list(string)
  description = "[required] - id of resources to be updated"
}


