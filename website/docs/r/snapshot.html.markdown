---
layout: "outscale"
page_title: "3DS OUTSCALE: outscale_snapshot"
sidebar_current: "outscale-snapshot"
description: |-
  [Manages a snapshot.]
---

# outscale_snapshot Resource

Manages a snapshot.
For more information on this resource, see the [User Guide](https://wiki.outscale.net/display/EN/About+Snapshots).
For more information on this resource actions, see the [API documentation](https://docs.outscale.com/api#3ds-outscale-api-snapshot).

## Example Usage

```hcl

# Create a snapshot

#resource "outscale_volume" "volume01" {
#  subregion_name = "${var.region}a"
#  size           = 40
#}

resource "outscale_snapshot" "snapshot01" {
  volume_id = outscale_volume.volume01.volume_id
}

# Copy a snapshot

resource "outscale_snapshot" "snapshot02" {
  description        = "Terraform snapshot copy"
  source_snapshot_id = "snap-12345678"
  source_region_name = "eu-west-2"
}


```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) A description for the snapshot.
* `file_location` - (Optional) The pre-signed URL of the snapshot you want to import from the OSU bucket.
* `snapshot_size` - (Optional) The size of the snapshot you want to create in your account, in bytes. This size must be greater than or equal to the size of the original, uncompressed snapshot.
* `source_region_name` - (Optional) The name of the source Region, which must be the same as the Region of your account.
* `source_snapshot_id` - (Optional) The ID of the snapshot you want to copy.
* `volume_id` - (Optional) The ID of the volume you want to create a snapshot of.
* `tags` - One or more tags to add to this resource.
      * `key` - The key of the tag, with a minimum of 1 character.
      * `value` - The value of the tag, between 0 and 255 characters.
    
## Attribute Reference

The following attributes are exported:

* `snapshot` - Information about the snapshot.
  * `account_alias` - The account alias of the owner of the snapshot.
  * `account_id` - The account ID of the owner of the snapshot.
  * `description` - The description of the snapshot.
  * `permissions_to_create_volume` - Information about the users who have permissions for the resource.
      * `account_ids` - The account ID of one or more users who have permissions for the resource.
      * `global_permission` - If `true`, the resource is public. If `false`, the resource is private.
  * `progress` - The progress of the snapshot, as a percentage.
  * `snapshot_id` - The ID of the snapshot.
  * `state` - The state of the snapshot (`in-queue` \| `completed` \| `error`).
  * `tags` - One or more tags associated with the snapshot.
      * `key` - The key of the tag, with a minimum of 1 character.
      * `value` - The value of the tag, between 0 and 255 characters.
  * `volume_id` - The ID of the volume used to create the snapshot.
  * `volume_size` - The size of the volume used to create the snapshot, in gibibytes (GiB).