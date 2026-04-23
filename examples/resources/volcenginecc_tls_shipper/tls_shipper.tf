resource "volcenginecc_tls_shipper" "TLSShipperDemo" {
  content_info = {
    format = "json"
    json_info = {
      enable = true
      escape = true
      keys = [
        "__content__",
        "__source__",
        "__path__",
        "__time__",
        "__time_ns_part__",
        "__image_name__",
        "__container_name__",
        "__container_ip__",
        "__container_source__",
        "__pod_ip__",
        "__pod_name__",
        "__pod_uid__",
        "__namespace__",
        "__tag____client_ip__",
        "__tag____receive_time__"
      ]
    }
  }
  kafka_shipper_info = {
    compress    = "gzip"
    end_time    = 1777538948000
    instance    = "kafka-cnngsl83xxxxxxxx"
    kafka_topic = "topic1"
  }
  shipper_end_time = 1777538948000
  shipper_name     = "ccapi-test-kaf-1001"
  shipper_type     = "kafka"
  topic_id         = "b75fffd8-1986-460c-9cca-5axxxxxxx"
  status           = false
}