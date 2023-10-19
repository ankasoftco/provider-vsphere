package vspherecustomattribute

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_custom_attribute", func(r *config.Resource) {
		r.ShortGroup = "Inventory"
		r.Kind = "VSphereCustomAttribute"
		r.Version = "v1alpha1"
	})
}
