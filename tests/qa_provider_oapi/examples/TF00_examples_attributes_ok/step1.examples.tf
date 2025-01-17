#---001------------------------------------------------------------------
resource "outscale_vm" "vm001" {
    image_id                 = var.image_id
    vm_type                  = var.vm_type
    keypair_name             = var.keypair_name
    security_group_ids       = [var.security_group_id]
    placement_subregion_name = format("%s%s", var.region, "a")
    placement_tenancy        = "default"
}
#------------------------------------------------------------------------

#---002------------------------------------------------------------------
resource "outscale_vm" "vm002" {
    image_id                 = var.image_id
    vm_type                  = var.vm_type
    keypair_name             = var.keypair_name
    security_group_names     = [var.security_group_name]
    placement_subregion_name = format("%s%s", var.region, "a")
    placement_tenancy        = "default"
}
#------------------------------------------------------------------------

#---003------------------------------------------------------------------
resource "outscale_vm" "vm003" {
    image_id           = var.image_id
    vm_type            = var.vm_type
    keypair_name       = var.keypair_name
    security_group_ids = [var.security_group_id]
}

data "outscale_vm" "vm_003d" {
	filter {
        name   = "vm_ids"
        values = [outscale_vm.vm003.vm_id]
    }
}
#------------------------------------------------------------------------

#---004------------------------------------------------------------------
resource "outscale_vm" "vm004" {
    count = 2

    image_id           = var.image_id
    vm_type            = var.vm_type
    keypair_name       = var.keypair_name
    security_group_ids = [var.security_group_id]
}

data "outscale_vm" "outscale_vm004_0d" {
	filter {
        name   = "vm_ids"
        values = [outscale_vm.vm004.*.vm_id[0]]
    }
}

data "outscale_vm" "outscale_vm_004_1d" {
	filter {
        name   = "vm_ids"
        values = [outscale_vm.vm004.*.vm_id[1]]
    }
}
#------------------------------------------------------------------------

#---005------------------------------------------------------------------
resource "outscale_net" "net005" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_subnet" "subnet005" {
    subregion_name = format("%s%s", var.region, "a")
    ip_range       = "10.0.0.0/16"
    net_id         = outscale_net.net005.net_id
}

resource "outscale_security_group" "security_group005" {
    count = 1

    description         = "test group005"
    security_group_name = "sg1-test-group_test-net005"
    net_id              = outscale_net.net005.net_id
}

resource "outscale_vm" "outscale_vm005" {
    image_id                 = var.image_id
    vm_type                  = var.vm_type
    keypair_name             = var.keypair_name
    security_group_ids       = [outscale_security_group.security_group005.security_group_id]
    placement_subregion_name = format("%s%s", var.region, "a")
    placement_tenancy        = "default"
    #is_source_dest_checked   = true
    subnet_id                = outscale_subnet.subnet005.subnet_id
}
#------------------------------------------------------------------------

#---014------------------------------------------------------------------
#TODO outscale_vm_attributes (merged in outscale_vm)

#------------------------------------------------------------------------



#------------------------------------------------------------------------

#---006------------------------------------------------------------------
resource "outscale_vm" "vm006" {
    image_id                 = var.image_id
    vm_type                  = var.vm_type
    keypair_name             = var.keypair_name
    security_group_ids       = [var.security_group_id]
    #placement_subregion_name = format("%s%s", var.region, "a")
    #placement_tenancy        = "default"
}

data "outscale_vm_state" "vm_state006d" {
     filter {
         name   = "vm_ids"
         values = [outscale_vm.vm006.vm_id]
     }
}
#------------------------------------------------------------------------

#---007------------------------------------------------------------------
resource "outscale_public_ip" "public_ip007" {
}
#------------------------------------------------------------------------

#---008------------------------------------------------------------------
resource "outscale_public_ip" "public_ip008" {
}

#data "outscale_public_ip" "public_ip008d" {
#   filter {
#      name  = "public_ips"
#      values = [outscale_public_ip.public_ip008.public_ip]
#   }    
#}
#------------------------------------------------------------------------

#---009------------------------------------------------------------------
resource "outscale_public_ip" "public_ip009" {
}

resource "outscale_vm" "vm009" {
    image_id           = var.image_id
    vm_type            = var.vm_type
    keypair_name       = var.keypair_name
    security_group_ids = [var.security_group_id]
}

resource "outscale_public_ip_link" "public_ip_link009" {
    vm_id     = outscale_vm.vm009.vm_id
    public_ip = outscale_public_ip.public_ip009.public_ip
}
#------------------------------------------------------------------------

#---010------------------------------------------------------------------
resource "outscale_volume" "volume010" {
    subregion_name = format("%s%s", var.region, "a")
    size           = 10
}
#------------------------------------------------------------------------

#---011------------------------------------------------------------------
resource "outscale_volume" "volume011" {
    subregion_name = format("%s%s", var.region, "a")
    size            = 10
    iops            = 100
    volume_type     = "io1"
}
#------------------------------------------------------------------------

#---012------------------------------------------------------------------
resource "outscale_volume" "volume012" {
    subregion_name = format("%s%s", var.region, "a")
    size            = 10
    snapshot_id     = var.snapshot_id
}
#------------------------------------------------------------------------

#---013------------------------------------------------------------------
resource "outscale_volume" "volume013" {
    subregion_name = format("%s%s", var.region, "a")
    size           = 10
    iops           = 100
    volume_type    = "io1"

}

data "outscale_volume" "outscale_volume013d" {
    filter {
        name   = "volume_ids"
        values = [outscale_volume.volume013.volume_id]
    }    
}
#------------------------------------------------------------------------

#---015------------------------------------------------------------------
resource "outscale_volume" "volume015" {
    subregion_name = format("%s%s", var.region, "a")
    size           = 40
}

resource "outscale_vm" "vm015" {
    image_id           = var.image_id
    vm_type            = var.vm_type
    keypair_name       = var.keypair_name
    security_group_ids = [var.security_group_id]
}

resource "outscale_volumes_link" "volumes_link015" {
    device_name = "/dev/xvdc"
    volume_id   = outscale_volume.volume015.id
    vm_id       = outscale_vm.vm015.id
}
#------------------------------------------------------------------------

#---016------------------------------------------------------------------
resource "outscale_internet_service" "internet_service016" {
} 
#------------------------------------------------------------------------

#---017------------------------------------------------------------------
resource "outscale_internet_service" "internet_service017" {
}

data "outscale_internet_service" "internet_service017d" {
    filter {
        name   = "internet_service_ids"
        values = [outscale_internet_service.internet_service017.internet_service_id]
    }
}
#------------------------------------------------------------------------

#---018------------------------------------------------------------------
resource "outscale_internet_service_link" "internet_service_link018" {
    internet_service_id = outscale_internet_service.internet_service018.internet_service_id 
    net_id              = outscale_net.net018.net_id
}

resource "outscale_net" "net018" {
    ip_range = "10.0.0.0/18"
}

resource "outscale_internet_service" "internet_service018" {
} 
#------------------------------------------------------------------------

#---019------------------------------------------------------------------
resource "outscale_net" "net019" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_subnet" "subnet019" {
    net_id     = outscale_net.net019.net_id
    ip_range = "10.0.0.0/18"
}

resource "outscale_public_ip" "public_ip019" {
}

resource "outscale_nat_service" "nat_service019" {
    depends_on   = ["outscale_route.route019"]
    subnet_id    = outscale_subnet.subnet019.subnet_id
    public_ip_id = outscale_public_ip.public_ip019.public_ip_id
}

resource "outscale_route_table" "route_table019" {
    net_id = outscale_net.net019.net_id
}

resource "outscale_route" "route019" {
    destination_ip_range = "0.0.0.0/0"
    gateway_id           = outscale_internet_service.internet_service019.internet_service_id
    route_table_id       = outscale_route_table.route_table019.route_table_id
}

resource "outscale_route_table_link" "route_table_link019" {
    subnet_id      = outscale_subnet.subnet019.subnet_id
    route_table_id = outscale_route_table.route_table019.id
}

resource "outscale_internet_service" "internet_service019" {
}

resource "outscale_internet_service_link" "internet_service_link019" {
    net_id              = outscale_net.net019.net_id
    internet_service_id = outscale_internet_service.internet_service019.id
}

data "outscale_nat_service" "nat_service019" {
    filter {
        name   = "nat_service_ids"
        values = [outscale_nat_service.nat_service019.nat_service_id]
    }    
}
#------------------------------------------------------------------------

#---020------------------------------------------------------------------
resource "outscale_net" "net020" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_subnet" "subnet020" {
    net_id   = outscale_net.net020.net_id
    ip_range = "10.0.0.0/18"
}
#------------------------------------------------------------------------





#------------------------------------------------------------------------
#TODO outscale_route_table
#------------------------------------------------------------------------
#TODO outscale_route_table_link (1)
#------------------------------------------------------------------------
#TODO outscale_route
#------------------------------------------------------------------------


#---024------------------------------------------------------------------
resource "outscale_net" "net024" {
    ip_range = "10.0.0.0/16"
    
    tags = [{
        key = "Name"
        value = "outscale_net_resource"
    },
    {
        key = "VerSion"
        value = "1Q84"
    }]
}

data "outscale_net" "net024d" {
    filter {
        name   = "net_ids"
        values = [outscale_net.net024.net_id]
    }
} 
#------------------------------------------------------------------------

#---025------------------------------------------------------------------
resource "outscale_net" "net025" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_net_attributes" "net_attributes025" {
     net_id              = outscale_net.net025.net_id
     dhcp_options_set_id = var.dhcp_options_set_id
}

data "outscale_net_attributes" "net_attributes025d" {
     net_id             = outscale_net.net025.net_id
}
#------------------------------------------------------------------------



#---026------------------------------------------------------------------
#TODO outscale_net_peering
#------------------------------------------------------------------------

#---027------------------------------------------------------------------
#TODO outscale_net_peering_acceptation (6)
#------------------------------------------------------------------------




#---021------------------------------------------------------------------
data "outscale_image" "image021" {
    filter {
        name   = "image_ids"
        values = [var.image_id]
    }
} 
#------------------------------------------------------------------------


#---022------------------------------------------------------------------
resource "outscale_image" "image022" {
    image_name = "terraform test for image attributes"
    vm_id      = var.vm_id
    no_reboot  = "true"
} 
#------------------------------------------------------------------------


#---023------------------------------------------------------------------
resource "outscale_vm" "vm023" {
    image_id           = var.image_id
    vm_type            = var.vm_type
    keypair_name       = var.keypair_name
    security_group_ids = [var.security_group_id]
}

resource "outscale_image" "image023" {
    image_name = "terraform test image launch permission"
    vm_id      = outscale_vm.vm023.vm_id
    no_reboot  = "true"
}

resource "outscale_image_launch_permission" "image_launch_permission023" {
    image_id = outscale_image.image023.image_id

    permission_additions = {
		account_ids = [var.account_id]
	}
} 
#------------------------------------------------------------------------






#---028------------------------------------------------------------------
resource "outscale_keypair" "keypair028" {
    keypair_name = "keyname_test_ford028"
}

data "outscale_keypair" "keypair028d" {
    filter {
        name   = "keypair_names"
        values = [outscale_keypair.keypair028.keypair_name]
    }    
}
#------------------------------------------------------------------------

#---029------------------------------------------------------------------
resource "outscale_keypair" "keypair029" {
    keypair_name = "keyname_test_import029"
    public_key   = file("keypair_public_test.pub")
} 
#------------------------------------------------------------------------



#---030------------------------------------------------------------------
resource "outscale_net" "net030" {
    ip_range = "10.0.0.0/16"

    tags = {
        key   = "Name"
        value = "outscale_net_resource030"
    }
}

data "outscale_security_group" "security_group030d" {     
    filter {
        name   = "security_group_ids"
        values = [outscale_security_group.security_group030.security_group_id]
    }
}

resource "outscale_security_group" "security_group030" {
    description         = "test group"
    security_group_name = "sg1-test-group_test-d030"
    net_id              = outscale_net.net030.net_id
}
#------------------------------------------------------------------------

#---031------------------------------------------------------------------
resource "outscale_security_group" "security_group031" {
    description         = "test group031"
    security_group_name = "sg1-test-group_test031"
}

resource "outscale_security_group_rule" "security_group_rule031" {
    flow              = "Inbound"
    security_group_id = outscale_security_group.security_group031.id

    from_port_range   = "0"
    to_port_range     = "0"
    #ip_protocol       = "-1"
    ip_protocol       = "tcp"
    ip_range          = "0.0.0.0/0"
} 
#------------------------------------------------------------------------

#---032------------------------------------------------------------------
resource "outscale_net" "net032" {
    ip_range = "10.0.0.0/16"
}

resource "outscale_subnet" "subnet032" {
    subregion_name = format("%s%s", var.region, "a")
    ip_range       = "10.0.0.0/16"
    net_id         = outscale_net.net032.net_id
}

resource "outscale_nic" "nic032" {
    subnet_id = outscale_subnet.subnet032.subnet_id
} 
#------------------------------------------------------------------------

#---033------------------------------------------------------------------
resource "outscale_volume" "volume033" {
    subregion_name = format("%s%s", var.region, "a")
    size           = 40
}

resource "outscale_snapshot" "snapshot033" {
    volume_id = outscale_volume.volume033.volume_id
}

data "outscale_snapshot" "snapshot033d" {
    filter {
        name   = "snapshot_ids"
        values = [outscale_snapshot.snapshot033.snapshot_id]
    }
} 
#------------------------------------------------------------------------


#---034------------------------------------------------------------------
resource "outscale_volume" "volume034" {
    subregion_name = format("%s%s", var.region, "a")
    size            = 40
}

resource "outscale_snapshot" "snapshot034" {
    volume_id = outscale_volume.volume034.volume_id
}

resource "outscale_snapshot_attributes" "snapshot_attributes034" {
    snapshot_id = outscale_snapshot.snapshot034.snapshot_id
    permissions_to_create_volume_additions = {
			account_ids = [var.account_id]
	}
}

data "outscale_snapshot" "snapshot034d" {
    depends_on  = ["outscale_snapshot_attributes.snapshot_attributes034"]
    snapshot_id = outscale_snapshot.snapshot034.snapshot_id
} 
#------------------------------------------------------------------------

#------------------------------------------------------------------------
#TODO outscale_tag (2)
#------------------------------------------------------------------------


#------------------------------------------------------------------------
#TODO outscale_nic_link
#------------------------------------------------------------------------
#TODO outscale_nic_private_ip
#------------------------------------------------------------------------

