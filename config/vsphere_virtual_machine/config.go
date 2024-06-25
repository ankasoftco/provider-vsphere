package vspherevirtualmachine

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_virtual_machine", func(r *config.Resource) {
		r.ShortGroup = "virtualmachine"
		r.Kind = "VSphereVirtualMachine"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_virtual_machine_snapshot", func(r *config.Resource) {
		r.ShortGroup = "virtualmachine"
		r.Kind = "VSphereVirtualMachineSnapshot"
		r.Version = "v1alpha1"
	})
}
