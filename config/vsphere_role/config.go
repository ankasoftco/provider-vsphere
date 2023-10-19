package vsphererole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_role", func(r *config.Resource) {
		r.ShortGroup = "Security"
		r.Kind = "VSphereRole"
		r.Version = "v1alpha1"
	})
}
