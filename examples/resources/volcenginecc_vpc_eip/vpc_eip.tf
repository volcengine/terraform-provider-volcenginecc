resource "volcenginecc_vpc_eip" "EipDemo" {
  name                 = "EipDemo"
  description          = "EipDemo description"
  isp                  = "BGP"
  billing_type         = 2
  bandwidth            = 3
  period               = 5
  project_name         = "default"
  bandwidth_package_id = "bwp-ij5gz1lf66m874o8cth*****"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  instance_id   = "i-ye48ymyy9s5i3z4*****"
  instance_type = "EcsInstance"
  direct_mode   = true
}