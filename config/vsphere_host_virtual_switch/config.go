package vspherehostvirtualswitch

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_host_virtual_switch", func(r *config.Resource) {
		r.ShortGroup = "Networking"
		r.Kind = "VSphereHostVirtualSwitch"
		r.Version = "v1alpha1"
	})
}
