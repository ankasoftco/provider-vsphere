package vspheredatastorecluster

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_datastore_cluster", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereDatastoreCluster"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_datastore_cluster_vm_anti_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereDatastoreClusterVmAntiAffinityRule"
		r.Version = "v1alpha1"
	})
}
