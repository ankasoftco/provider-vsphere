package vspheretag

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_tag", func(r *config.Resource) {
		r.ShortGroup = "Inventory"
		r.Kind = "VSphereTag"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_tag_category", func(r *config.Resource) {
		r.ShortGroup = "Inventory"
		r.Kind = "VSphereTagCategory"
		r.Version = "v1alpha1"
	})
}
