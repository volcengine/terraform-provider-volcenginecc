resource "volcenginecc_cloudmonitor_object_group" "example" {
  name = "ccapi_object_group_dynamic_meta"
  objects = [{
    dimension_conditions = {
      meta_condition = {
        all_dimensions = false
        condition      = "or"
        metas = [{
          comparator = "contain"
          key        = "ResourceName"
          values     = ["test"]
        }]
      }
      type = "meta"
    }
    namespace = "VCM_ECS"
    region    = "cn-beijing"
    type      = "dynamic"
  }]
}
