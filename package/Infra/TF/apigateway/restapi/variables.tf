variable "rest_api_name" {
  type = string
  default = "demo"
  description = "[optional] - name of rest api"
}

variable "endpoint" {
  type = string
  description = "[required] - choose REGIONAL or PRIVATE "
}

variable "vpc_endpoint_id" {
  type = string
  description = "[required] - vpc endpoint id of api gateway"
}

