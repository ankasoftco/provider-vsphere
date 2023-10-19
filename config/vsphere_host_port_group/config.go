package vspherehostportgroup

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_host_port_group", func(r *config.Resource) {
		r.ShortGroup = "Networking"
		r.Kind = "VSphereHostPortGroup"
		r.Version = "v1alpha1"
	})
}
