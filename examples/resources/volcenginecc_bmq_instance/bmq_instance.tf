resource "volcenginecc_bmq_instance" "BMQInstanceDemo" {
  name              = "BMQInstanceDemo"
  billing_type      = "POST"
  project_name      = "default"
  specification     = "bmq.standard"
  vpc_id            = "vpc-miltj87lh2ww5smt1bxxxxx"
  message_retention = 1
  endpoints = {
    public = {
      eip_id = "eip-3nriu2y2ufwu8931exxxxx"
    }
    overlay = {
      vpc_ids = ["vpc-miltj87lh2ww5smt1bxxxxx"]
    }
  }
  security_group_id_list = ["sg-3nqnz9en1ucxs931eaxxxxx"]
  subnet_id_list         = ["subnet-w02wsq25fitc865ykaxxxxx"]
  zone_id_list           = ["cn-beijing-a"]
  tags = [
    {
      key   = "env"
      type  = "CUSTOM"
      value = "test"
    }
  ]
}