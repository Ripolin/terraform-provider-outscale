resource "outscale_volume" "outscale_volume_snap" {
    subregion_name = format("%s%s", var.region, "a")
    size           = 40
}

resource "outscale_snapshot" "outscale_snapshot" {
    volume_id = outscale_volume.outscale_volume_snap.volume_id
    tags {
    key = "name"
    value = "Snapsho_Creation_test"
    }
}

### Tes tCopy snapshot
resource "outscale_snapshot" "outscale_snapshot_2" {
description = "Test-copy-snapshot"
source_snapshot_id = var.snapshot_id
source_region_name = var.region
tags {
    key = "name"
    value = "Snapsho_Copy_test"
    }
}
