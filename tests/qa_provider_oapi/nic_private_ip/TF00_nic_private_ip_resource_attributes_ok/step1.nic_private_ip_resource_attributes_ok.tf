resource "outscale_net" "outscale_net" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_subnet" "outscale_subnet" {
    subregion_name = format("%s%s", var.region, "a")
    ip_range       = "10.0.0.0/16"
    net_id         = outscale_net.outscale_net.net_id
}

resource "outscale_nic" "outscale_nic" {
    subnet_id = outscale_subnet.outscale_subnet.subnet_id
}

resource "outscale_nic_private_ip" "outscale_nic_private_ip" {
    nic_id      = outscale_nic.outscale_nic.nic_id
    #secondary_private_ip_address_count = 1
    private_ips = ["10.0.45.67"]
} 
