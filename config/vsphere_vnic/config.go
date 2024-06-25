package vspherevnic

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_vnic", func(r *config.Resource) {
		r.ShortGroup = "hostandclustermanagement"
		r.Kind = "VSphereVnic"
		r.Version = "v1alpha1"
	})
}
