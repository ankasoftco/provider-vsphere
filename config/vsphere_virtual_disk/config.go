package vspherevirtualdisk

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_virtual_disk", func(r *config.Resource) {
		r.ShortGroup = "virtualmachine"
		r.Kind = "VSphereVirtualDisk"
		r.Version = "v1alpha1"
	})
}
