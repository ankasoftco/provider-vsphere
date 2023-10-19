package vspherestoragedrsvmoverride

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_storage_drs_vm_override", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereStorageDrsVmOverride"
		r.Version = "v1alpha1"
	})
}
