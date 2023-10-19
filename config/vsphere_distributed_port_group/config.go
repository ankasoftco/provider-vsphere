package vspheredistributedportgroup

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_distributed_port_group", func(r *config.Resource) {
		r.ShortGroup = "Networking"
		r.Kind = "VSphereDistributedPortGroup"
		r.Version = "v1alpha1"
	})
}
