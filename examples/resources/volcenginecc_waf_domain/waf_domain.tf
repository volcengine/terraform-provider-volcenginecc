resource "volcenginecc_waf_domain" "WafDomainDemo" {
  access_mode        = 10
  domain             = "www.test.com"
  lb_algorithm       = "wrr"
  public_real_server = 1
  vpc_id             = ""
  protocol_ports = {
    http = [80]
  }
  protocols = ["HTTP"]
  backend_groups = [
    {
      access_port = [80]
      backends = [
        {
          protocol = "HTTP"
          port     = 80
          ip       = "1.1.1.1"
          weight   = 50
        }
      ]
      name = "default"
    }
  ]
}




