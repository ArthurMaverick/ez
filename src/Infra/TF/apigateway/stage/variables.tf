variable "rest_api_id" {
  type = string
  description = "[required] - ID of the associated REST API"
}

variable "deployment_id" {
  type = string
  description = "[required] - ID of the deployment that the stage points to"
}

variable "stage_name" {
  type = string
  description = "[required] - Name of the stage"
}



