variable "rest_api_id" {
  type = string
  description = "[required] - ID of the associated REST API"
}

variable "resource_id" {
  type = string
  description = "[required] - API resource ID"
}

variable "http_method" {
  type = string
  description = "[required] - HTTP Method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)"
  default = "ANY"
}

variable "authorization" {
  type = string
  description = "[required] - Type of authorization used for the method (NONE, CUSTOM, AWS_IAM, COGNITO_USER_POOLS)"
  default = "NONE"
}