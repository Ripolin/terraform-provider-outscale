resource "outscale_volume" "outscale_volume" {
    subregion_name = format("%s%s", var.region, "a")
    size = 10
    iops           = 100
    volume_type    = "io1"
    tags {
        key = "Name"
        value = "volume-io1-1"
    }
}

resource "outscale_volume" "outscale_volume2" {
    subregion_name = format("%s%s", var.region, "a")
    size = 10
    iops           = 100
    volume_type    = "io1"
    tags {
        key = "Name"
        value = "volume-io1-2"
    }
}

data "outscale_volumes" "outscale_volumes" {
    filter {
        name = "volume_ids"
        values = [outscale_volume.outscale_volume.volume_id, outscale_volume.outscale_volume2.volume_id]
    }
}
