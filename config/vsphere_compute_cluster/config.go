package vSphereComputeCluster

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_compute_cluster", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeCluster"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_host_group", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterHostGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterVmAffinityRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_anti_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterVmAntiAffinityRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_dependency_rule", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterVmDependencyRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_group", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterVmGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vsphere_compute_cluster_vm_host_rule", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereComputeClusterVmHostRule"
		r.Version = "v1alpha1"
	})
}