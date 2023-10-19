/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	vspherecomputecluster "github.com/ankasoftco/provider-vsphere/config/vsphere_compute_cluster"
	vspherecontentlibrary "github.com/ankasoftco/provider-vsphere/config/vsphere_content_library"
	vspherecustomattribute "github.com/ankasoftco/provider-vsphere/config/vsphere_custom_attribute"
	vspheredatacenter "github.com/ankasoftco/provider-vsphere/config/vsphere_datacenter"
	vspheredatastorecluster "github.com/ankasoftco/provider-vsphere/config/vsphere_datastore_cluster"
	vspheredistributedportgroup "github.com/ankasoftco/provider-vsphere/config/vsphere_distributed_port_group"
	vspheredistributedvirtualswitch "github.com/ankasoftco/provider-vsphere/config/vsphere_distributed_virtual_switch"
	vspheredpmhostoverride "github.com/ankasoftco/provider-vsphere/config/vsphere_dpm_host_override"
	vspheredrsvmoverride "github.com/ankasoftco/provider-vsphere/config/vsphere_drs_vm_override"
	vsphereentitypermissions "github.com/ankasoftco/provider-vsphere/config/vsphere_entity_permissions"
	vspherefile "github.com/ankasoftco/provider-vsphere/config/vsphere_file"
	vspherefolder "github.com/ankasoftco/provider-vsphere/config/vsphere_folder"
	vspherehavmoverride "github.com/ankasoftco/provider-vsphere/config/vsphere_ha_vm_override"
	vspherehost "github.com/ankasoftco/provider-vsphere/config/vsphere_host"
	vspherehostportgroup "github.com/ankasoftco/provider-vsphere/config/vsphere_host_port_group"
	vspherehostvirtualswitch "github.com/ankasoftco/provider-vsphere/config/vsphere_host_virtual_switch"
	vspherelicense "github.com/ankasoftco/provider-vsphere/config/vsphere_license"
	vspherenasdatastore "github.com/ankasoftco/provider-vsphere/config/vsphere_nas_datastore"
	vsphereresourcepool "github.com/ankasoftco/provider-vsphere/config/vsphere_resource_pool"
	vsphererole "github.com/ankasoftco/provider-vsphere/config/vsphere_role"
	vspherestoragedrsvmoverride "github.com/ankasoftco/provider-vsphere/config/vsphere_storage_drs_vm_override"
	vspheretag "github.com/ankasoftco/provider-vsphere/config/vsphere_tag"
	vspherevappcontainer "github.com/ankasoftco/provider-vsphere/config/vsphere_vapp_container"
	vspherevappentity "github.com/ankasoftco/provider-vsphere/config/vsphere_vapp_entity"
	vspherevirtualdisk "github.com/ankasoftco/provider-vsphere/config/vsphere_virtual_disk"
	vspherevirtualmachine "github.com/ankasoftco/provider-vsphere/config/vsphere_virtual_machine"
	vspherevmstoragepolicy "github.com/ankasoftco/provider-vsphere/config/vsphere_vm_storage_policy"
	vspherevmfsdatastore "github.com/ankasoftco/provider-vsphere/config/vsphere_vmfs_datastore"
	vspherevnic "github.com/ankasoftco/provider-vsphere/config/vsphere_vnic"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vsphere"
	modulePath     = "github.com/ankasoftco/provider-vsphere"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		vspherelicense.Configure,
		vspherecomputecluster.Configure,
		vspheredpmhostoverride.Configure,
		vspheredrsvmoverride.Configure,
		vspherehavmoverride.Configure,
		vspherehost.Configure,
		vsphereresourcepool.Configure,
		vspherevnic.Configure,
		vspherecustomattribute.Configure,
		vspheredatacenter.Configure,
		vspherefolder.Configure,
		vspheretag.Configure,
		vspheredistributedportgroup.Configure,
		vspheredistributedvirtualswitch.Configure,
		vspherehostportgroup.Configure,
		vspherehostvirtualswitch.Configure,
		vsphereentitypermissions.Configure,
		vsphererole.Configure,
		vspheredatastorecluster.Configure,
		vspherefile.Configure,
		vspherenasdatastore.Configure,
		vspherestoragedrsvmoverride.Configure,
		vspherevmstoragepolicy.Configure,
		vspherevmfsdatastore.Configure,
		vspherecontentlibrary.Configure,
		vspherevappcontainer.Configure,
		vspherevappentity.Configure,
		vspherevirtualdisk.Configure,
		vspherevirtualmachine.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
