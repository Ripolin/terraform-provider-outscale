resource "outscale_net" "outscale_net" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_net_attributes" "outscale_net_attributes" {
     net_id              = outscale_net.outscale_net.net_id
     dhcp_options_set_id = var.dhcp_options_set_id
}

data "outscale_net_attributes" "outscale_net_attributes" {
     net_id             = outscale_net.outscale_net.net_id
}
