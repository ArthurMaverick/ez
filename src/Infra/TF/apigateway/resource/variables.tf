variable "rest_api_id" {
  type = string
  description = "[required] - ID of the associated REST API"
}

variable "root_resource_id" {
  type = string
  description = "[required] - ID of the parent API resource"
}

variable "path_part" {
  type = string
  description = "[required] - Last path segment of this API resource."
  default = "demo"
}


