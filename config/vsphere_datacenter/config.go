package vspheredatacenter

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_datacenter", func(r *config.Resource) {
		r.ShortGroup = "Inventory"
		r.Kind = "VSphereDatacenter"
		r.Version = "v1alpha1"
	})
}
