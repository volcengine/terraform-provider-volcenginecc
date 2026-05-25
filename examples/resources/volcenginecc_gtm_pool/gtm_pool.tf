resource "volcenginecc_gtm_pool" "GTMPoolDemo" {
  addresses = [
    {
      rectified_geos = [
        "cn_telecom_shanghai"
      ],
      capacity = 20,
      mode     = "auto",
      value    = "2001:db8::2",
      weight   = 20
    },
    {
      rectified_geos = [
        "cn_telecom_beijing"
      ],
      capacity = 40,
      mode     = "auto",
      value    = "2001:db8::3",
      weight   = 21
    }
  ]
  gtm_id        = "162ca39d-*******"
  addr_type     = "ipv6"
  capacity_mode = "pool"
  capacity      = 90
  name          = "testpool"
  remark        = "testdes"
}