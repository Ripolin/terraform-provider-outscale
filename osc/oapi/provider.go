// GENERATED FILE: DO NOT EDIT!

package oapi

// To create a server, first write a class that implements this interface.
// Then pass an instance of it to Initialize().
type Provider interface {

	//
	POST_AcceptNetPeering(parameters *POST_AcceptNetPeeringParameters, responses *POST_AcceptNetPeeringResponses) (err error)

	//
	POST_CopyImage(parameters *POST_CopyImageParameters, responses *POST_CopyImageResponses) (err error)

	//
	POST_CopySnapshot(parameters *POST_CopySnapshotParameters, responses *POST_CopySnapshotResponses) (err error)

	//
	POST_CreateImage(parameters *POST_CreateImageParameters, responses *POST_CreateImageResponses) (err error)

	//
	POST_CreateImageExportTask(parameters *POST_CreateImageExportTaskParameters, responses *POST_CreateImageExportTaskResponses) (err error)

	//
	POST_CreateKeypair(parameters *POST_CreateKeypairParameters, responses *POST_CreateKeypairResponses) (err error)

	//
	POST_CreateLoadBalancer(parameters *POST_CreateLoadBalancerParameters, responses *POST_CreateLoadBalancerResponses) (err error)

	//
	POST_CreateLoadBalancerListeners(parameters *POST_CreateLoadBalancerListenersParameters, responses *POST_CreateLoadBalancerListenersResponses) (err error)

	//
	POST_CreateNatService(parameters *POST_CreateNatServiceParameters, responses *POST_CreateNatServiceResponses) (err error)

	//
	POST_CreateNet(parameters *POST_CreateNetParameters, responses *POST_CreateNetResponses) (err error)

	//
	POST_CreateNetPeering(parameters *POST_CreateNetPeeringParameters, responses *POST_CreateNetPeeringResponses) (err error)

	//
	POST_CreateNic(parameters *POST_CreateNicParameters, responses *POST_CreateNicResponses) (err error)

	//
	POST_CreatePublicIp(parameters *POST_CreatePublicIpParameters, responses *POST_CreatePublicIpResponses) (err error)

	//
	POST_CreateRoute(parameters *POST_CreateRouteParameters, responses *POST_CreateRouteResponses) (err error)

	//
	POST_CreateRouteTable(parameters *POST_CreateRouteTableParameters, responses *POST_CreateRouteTableResponses) (err error)

	//
	POST_CreateSnapshot(parameters *POST_CreateSnapshotParameters, responses *POST_CreateSnapshotResponses) (err error)

	//
	POST_CreateSnapshotExportTask(parameters *POST_CreateSnapshotExportTaskParameters, responses *POST_CreateSnapshotExportTaskResponses) (err error)

	//
	POST_CreateStickyCookiePolicy(parameters *POST_CreateStickyCookiePolicyParameters, responses *POST_CreateStickyCookiePolicyResponses) (err error)

	//
	POST_CreateSubnet(parameters *POST_CreateSubnetParameters, responses *POST_CreateSubnetResponses) (err error)

	//
	POST_CreateTags(parameters *POST_CreateTagsParameters, responses *POST_CreateTagsResponses) (err error)

	//
	POST_CreateVms(parameters *POST_CreateVmsParameters, responses *POST_CreateVmsResponses) (err error)

	//
	POST_CreateVolume(parameters *POST_CreateVolumeParameters, responses *POST_CreateVolumeResponses) (err error)

	//
	POST_DeleteKeypair(parameters *POST_DeleteKeypairParameters, responses *POST_DeleteKeypairResponses) (err error)

	//
	POST_DeleteLoadBalancer(parameters *POST_DeleteLoadBalancerParameters, responses *POST_DeleteLoadBalancerResponses) (err error)

	//
	POST_DeleteLoadBalancerListeners(parameters *POST_DeleteLoadBalancerListenersParameters, responses *POST_DeleteLoadBalancerListenersResponses) (err error)

	//
	POST_DeleteLoadBalancerPolicy(parameters *POST_DeleteLoadBalancerPolicyParameters, responses *POST_DeleteLoadBalancerPolicyResponses) (err error)

	//
	POST_DeleteNatService(parameters *POST_DeleteNatServiceParameters, responses *POST_DeleteNatServiceResponses) (err error)

	//
	POST_DeleteNet(parameters *POST_DeleteNetParameters, responses *POST_DeleteNetResponses) (err error)

	//
	POST_DeleteNetPeering(parameters *POST_DeleteNetPeeringParameters, responses *POST_DeleteNetPeeringResponses) (err error)

	//
	POST_DeleteNic(parameters *POST_DeleteNicParameters, responses *POST_DeleteNicResponses) (err error)

	//
	POST_DeletePublicIp(parameters *POST_DeletePublicIpParameters, responses *POST_DeletePublicIpResponses) (err error)

	//
	POST_DeleteRoute(parameters *POST_DeleteRouteParameters, responses *POST_DeleteRouteResponses) (err error)

	//
	POST_DeleteRouteTable(parameters *POST_DeleteRouteTableParameters, responses *POST_DeleteRouteTableResponses) (err error)

	//
	POST_DeleteSnapshot(parameters *POST_DeleteSnapshotParameters, responses *POST_DeleteSnapshotResponses) (err error)

	//
	POST_DeleteSubnet(parameters *POST_DeleteSubnetParameters, responses *POST_DeleteSubnetResponses) (err error)

	//
	POST_DeleteTags(parameters *POST_DeleteTagsParameters, responses *POST_DeleteTagsResponses) (err error)

	//
	POST_DeleteVms(parameters *POST_DeleteVmsParameters, responses *POST_DeleteVmsResponses) (err error)

	//
	POST_DeleteVolume(parameters *POST_DeleteVolumeParameters, responses *POST_DeleteVolumeResponses) (err error)

	//
	POST_DeregisterVmsInLoadBalancer(parameters *POST_DeregisterVmsInLoadBalancerParameters, responses *POST_DeregisterVmsInLoadBalancerResponses) (err error)

	//
	POST_ImportSnapshot(parameters *POST_ImportSnapshotParameters, responses *POST_ImportSnapshotResponses) (err error)

	//
	POST_LinkLoadBalancerServerCertificate(parameters *POST_LinkLoadBalancerServerCertificateParameters, responses *POST_LinkLoadBalancerServerCertificateResponses) (err error)

	//
	POST_LinkNic(parameters *POST_LinkNicParameters, responses *POST_LinkNicResponses) (err error)

	//
	POST_LinkPrivateIps(parameters *POST_LinkPrivateIpsParameters, responses *POST_LinkPrivateIpsResponses) (err error)

	//
	POST_LinkPublicIp(parameters *POST_LinkPublicIpParameters, responses *POST_LinkPublicIpResponses) (err error)

	//
	POST_LinkRouteTable(parameters *POST_LinkRouteTableParameters, responses *POST_LinkRouteTableResponses) (err error)

	//
	POST_LinkVolume(parameters *POST_LinkVolumeParameters, responses *POST_LinkVolumeResponses) (err error)

	//
	POST_ReadApiLogs(parameters *POST_ReadApiLogsParameters, responses *POST_ReadApiLogsResponses) (err error)

	//
	POST_ReadImageExportTasks(parameters *POST_ReadImageExportTasksParameters, responses *POST_ReadImageExportTasksResponses) (err error)

	//
	POST_ReadImages(parameters *POST_ReadImagesParameters, responses *POST_ReadImagesResponses) (err error)

	//
	POST_ReadKeypairs(parameters *POST_ReadKeypairsParameters, responses *POST_ReadKeypairsResponses) (err error)

	//
	POST_ReadLoadBalancerAttributes(parameters *POST_ReadLoadBalancerAttributesParameters, responses *POST_ReadLoadBalancerAttributesResponses) (err error)

	//
	POST_ReadLoadBalancers(parameters *POST_ReadLoadBalancersParameters, responses *POST_ReadLoadBalancersResponses) (err error)

	//
	POST_ReadNatServices(parameters *POST_ReadNatServicesParameters, responses *POST_ReadNatServicesResponses) (err error)

	//
	POST_ReadNetPeerings(parameters *POST_ReadNetPeeringsParameters, responses *POST_ReadNetPeeringsResponses) (err error)

	//
	POST_ReadNets(parameters *POST_ReadNetsParameters, responses *POST_ReadNetsResponses) (err error)

	//
	POST_ReadNics(parameters *POST_ReadNicsParameters, responses *POST_ReadNicsResponses) (err error)

	//
	POST_ReadPublicIps(parameters *POST_ReadPublicIpsParameters, responses *POST_ReadPublicIpsResponses) (err error)

	//
	POST_ReadRouteTables(parameters *POST_ReadRouteTablesParameters, responses *POST_ReadRouteTablesResponses) (err error)

	//
	POST_ReadSnapshotExportTasks(parameters *POST_ReadSnapshotExportTasksParameters, responses *POST_ReadSnapshotExportTasksResponses) (err error)

	//
	POST_ReadSnapshots(parameters *POST_ReadSnapshotsParameters, responses *POST_ReadSnapshotsResponses) (err error)

	//
	POST_ReadSubnets(parameters *POST_ReadSubnetsParameters, responses *POST_ReadSubnetsResponses) (err error)

	//
	POST_ReadTags(parameters *POST_ReadTagsParameters, responses *POST_ReadTagsResponses) (err error)

	//
	POST_ReadVmAttribute(parameters *POST_ReadVmAttributeParameters, responses *POST_ReadVmAttributeResponses) (err error)

	//
	POST_ReadVms(parameters *POST_ReadVmsParameters, responses *POST_ReadVmsResponses) (err error)

	//
	POST_ReadVmsState(parameters *POST_ReadVmsStateParameters, responses *POST_ReadVmsStateResponses) (err error)

	//
	POST_ReadVolumes(parameters *POST_ReadVolumesParameters, responses *POST_ReadVolumesResponses) (err error)

	//
	POST_RebootVms(parameters *POST_RebootVmsParameters, responses *POST_RebootVmsResponses) (err error)

	//
	POST_RegisterImage(parameters *POST_RegisterImageParameters, responses *POST_RegisterImageResponses) (err error)

	//
	POST_RegisterVmsInLoadBalancer(parameters *POST_RegisterVmsInLoadBalancerParameters, responses *POST_RegisterVmsInLoadBalancerResponses) (err error)

	//
	POST_RejectNetPeering(parameters *POST_RejectNetPeeringParameters, responses *POST_RejectNetPeeringResponses) (err error)

	//
	POST_StartVms(parameters *POST_StartVmsParameters, responses *POST_StartVmsResponses) (err error)

	//
	POST_StopVms(parameters *POST_StopVmsParameters, responses *POST_StopVmsResponses) (err error)

	//
	POST_UnlinkNic(parameters *POST_UnlinkNicParameters, responses *POST_UnlinkNicResponses) (err error)

	//
	POST_UnlinkPrivateIps(parameters *POST_UnlinkPrivateIpsParameters, responses *POST_UnlinkPrivateIpsResponses) (err error)

	//
	POST_UnlinkPublicIp(parameters *POST_UnlinkPublicIpParameters, responses *POST_UnlinkPublicIpResponses) (err error)

	//
	POST_UnlinkRouteTable(parameters *POST_UnlinkRouteTableParameters, responses *POST_UnlinkRouteTableResponses) (err error)

	//
	POST_UnlinkVolume(parameters *POST_UnlinkVolumeParameters, responses *POST_UnlinkVolumeResponses) (err error)

	//
	POST_UpdateImage(parameters *POST_UpdateImageParameters, responses *POST_UpdateImageResponses) (err error)

	//
	POST_UpdateLoadBalancerPolicies(parameters *POST_UpdateLoadBalancerPoliciesParameters, responses *POST_UpdateLoadBalancerPoliciesResponses) (err error)

	//
	POST_UpdateNic(parameters *POST_UpdateNicParameters, responses *POST_UpdateNicResponses) (err error)

	//
	POST_UpdateRoute(parameters *POST_UpdateRouteParameters, responses *POST_UpdateRouteResponses) (err error)

	//
	POST_UpdateRoutePropagation(parameters *POST_UpdateRoutePropagationParameters, responses *POST_UpdateRoutePropagationResponses) (err error)

	//
	POST_UpdateRouteTableLink(parameters *POST_UpdateRouteTableLinkParameters, responses *POST_UpdateRouteTableLinkResponses) (err error)

	//
	POST_UpdateSnapshot(parameters *POST_UpdateSnapshotParameters, responses *POST_UpdateSnapshotResponses) (err error)
}
