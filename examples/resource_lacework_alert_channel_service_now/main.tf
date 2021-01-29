provider "lacework" {}

resource "lacework_alert_channel_service_now" "example" {
  name         = "Service Now Channel Alert Example"
  instance_url = "snow-lacework.com"
  username     = "snow-user"
  password     = "snow-pass"
}