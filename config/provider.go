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
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
