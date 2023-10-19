package vspherelicense

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_license", func(r *config.Resource) {
		r.ShortGroup = "Administration"
		r.Kind = "VSphereLicense"
		r.Version = "v1alpha1"
	})
}
