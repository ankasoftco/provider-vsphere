package vsphereresourcepool

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_resource_pool", func(r *config.Resource) {
		r.ShortGroup = "hostandclustermanagement"
		r.Kind = "VSphereResourcePool"
		r.Version = "v1alpha1"
	})
}
