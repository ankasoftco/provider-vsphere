package vspherecomputecluster

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "hostandclustermanagement"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_compute_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeCluster"
		r.Version = version
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_host_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterHostGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterVmAffinityRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_anti_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterVmAntiAffinityRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_dependency_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterVmDependencyRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterVmGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_host_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VSphereComputeClusterVmHostRule"
		r.Version = version
	})
}
