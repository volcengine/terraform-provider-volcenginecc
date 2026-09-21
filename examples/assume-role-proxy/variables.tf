variable "role_trn" {
  type        = string
  description = "Temporary test role TRN; credentials come exclusively from the environment."
}
variable "vpc_name" {
  type        = string
  description = "Unique name for the VPC created by this run."
}
variable "assume_role_proxy_url" {
  type        = string
  default     = null
  sensitive   = true
  description = "null inherits HTTP_PROXY; an explicit URL overrides it."
}
