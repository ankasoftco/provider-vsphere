package vspherefolder

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_folder", func(r *config.Resource) {
		r.ShortGroup = "Inventory"
		r.Kind = "VSphereFolder"
		r.Version = "v1alpha1"
	})
}
