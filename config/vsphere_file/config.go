package vspherefile

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_file", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereFile"
		r.Version = "v1alpha1"
	})
}
