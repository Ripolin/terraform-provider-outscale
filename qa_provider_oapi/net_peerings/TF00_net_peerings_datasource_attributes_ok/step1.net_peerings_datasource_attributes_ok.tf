resource "outscale_net" "outscale_net" {
    count = 1

    ip_range = "10.10.0.0/24"
}

resource "outscale_net" "outscale_net2" {
    count = 1

    ip_range = "10.31.0.0/16"
}

resource "outscale_net" "outscale_net3" {
    count = 1

    ip_range = "20.0.0.0/16"
}

resource "outscale_net_peering" "outscale_net_peering" {
    accepter_net_id   = outscale_net.outscale_net.net_id
    source_net_id     = outscale_net.outscale_net2.net_id
}

resource "outscale_net_peering" "outscale_net_peering2" {
    accepter_net_id   = outscale_net.outscale_net.net_id
    source_net_id     = outscale_net.outscale_net3.net_id
}


data "outscale_net_peerings" "outscale_net_peerings" {

    filter {
         name = "net_peering_ids"
         values = [outscale_net_peering.outscale_net_peering.net_peering_id, outscale_net_peering.outscale_net_peering2.net_peering_id]
    }
}
