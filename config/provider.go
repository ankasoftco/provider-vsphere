/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	vSphereLicense "github.com/ankasoftco/provider-vsphere/config/vsphere_license"
	vSphereComputeCluster "github.com/ankasoftco/provider-vsphere/config/vsphere_compute_cluster"
	vSphereDpmHostOverride "github.com/ankasoftco/provider-vsphere/config/vsphere_dpm_host_override"
	vSphereDrsVmOverride "github.com/ankasoftco/provider-vsphere/config/vsphere_drs_vm_override"
	vSphereHaVmOverride "github.com/ankasoftco/provider-vsphere/config/vsphere_ha_vm_override"
	vSphereHost "github.com/ankasoftco/provider-vsphere/config/vsphere_host"
	vSphereResourcePool "github.com/ankasoftco/provider-vsphere/config/vsphere_resource_pool"
	vSphereVnic "github.com/ankasoftco/provider-vsphere/config/vsphere_vnic"
	vSphereCustomAttribute "github.com/ankasoftco/provider-vsphere/config/vsphere_custom_attribute"
	vSphereDatacenter "github.com/ankasoftco/provider-vsphere/config/vsphere_datacenter"
	vSphereFolder "github.com/ankasoftco/provider-vsphere/config/vsphere_folder"
	vSphereTag "github.com/ankasoftco/provider-vsphere/config/vsphere_tag"
	vSphereDistributedPortGroup "github.com/ankasoftco/provider-vsphere/config/vsphere_distributed_port_group"
	vSphereDistributedVirtualSwitch "github.com/ankasoftco/provider-vsphere/config/vsphere_distributed_virtual_switch"
	vSphereHostPortGroup "github.com/ankasoftco/provider-vsphere/config/vsphere_host_port_group"
	vSphereHostVirtualSwitch "github.com/ankasoftco/provider-vsphere/config/vsphere_host_virtual_switch"
	vSphereEntityPermissions "github.com/ankasoftco/provider-vsphere/config/vsphere_entity_permissions"
	vSphereRole "github.com/ankasoftco/provider-vsphere/config/vsphere_role"
	vSphereDatastoreCluster "github.com/ankasoftco/provider-vsphere/config/vsphere_datastore_cluster"
	vSphereFile "github.com/ankasoftco/provider-vsphere/config/vsphere_file"
	vSphereNasDatastore "github.com/ankasoftco/provider-vsphere/config/vsphere_nas_datastore"
	vSphereStorageDrsVmOverride "github.com/ankasoftco/provider-vsphere/config/vsphere_storage_drs_vm_override"
	vSphereVmStoragePolicy "github.com/ankasoftco/provider-vsphere/config/vsphere_vm_storage_policy"
	vSphereVmfsDatastore "github.com/ankasoftco/provider-vsphere/config/vsphere_vmfs_datastore"
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
		vSphereLicense.Configure,
		vSphereComputeCluster.Configure,
		vSphereDpmHostOverride.Configure,
		vSphereDrsVmOverride.Configure,
		vSphereHaVmOverride.Configure,
		vSphereHost.Configure,
		vSphereResourcePool.Configure,
		vSphereVnic.Configure,
		vSphereCustomAttribute.Configure,
		vSphereDatacenter.Configure,
		vSphereFolder.Configure,
		vSphereTag.Configure,
		vSphereDistributedPortGroup.Configure,
		vSphereDistributedVirtualSwitch.Configure,
		vSphereHostPortGroup.Configure,
		vSphereHostVirtualSwitch.Configure,
		vSphereEntityPermissions.Configure,
		vSphereRole.Configure,
		vSphereDatastoreCluster.Configure,
		vSphereFile.Configure,
		vSphereNasDatastore.Configure,
		vSphereStorageDrsVmOverride.Configure,
		vSphereVmStoragePolicy.Configure,
		vSphereVmfsDatastore.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
