package vspherenasdatastore

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_nas_datastore", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereNasDatastore"
		r.Version = "v1alpha1"
	})
}
