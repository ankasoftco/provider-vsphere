/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vsphere_license":                                 config.IdentifierFromProvider,
	"vsphere_compute_cluster":                         config.IdentifierFromProvider,
	"vsphere_compute_cluster_host_group":              config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_affinity_rule":        config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_anti_affinity_rule":   config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_dependency_rule":      config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_group":                config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_host_rule":            config.IdentifierFromProvider,
	"vsphere_dpm_host_override":                       config.IdentifierFromProvider,
	"vsphere_drs_vm_override":                         config.IdentifierFromProvider,
	"vsphere_ha_vm_override":                          config.IdentifierFromProvider,
	"vsphere_host":                                    config.IdentifierFromProvider,
	"vsphere_resource_pool":                           config.IdentifierFromProvider,
	"vsphere_vnic":                                    config.IdentifierFromProvider,
	"vsphere_custom_attribute":                        config.IdentifierFromProvider,
	"vsphere_datacenter":                              config.IdentifierFromProvider,
	"vsphere_folder":                                  config.IdentifierFromProvider,
	"vsphere_tag":                                     config.IdentifierFromProvider,
	"vsphere_tag_category":                            config.IdentifierFromProvider,
	"vsphere_distributed_port_group":                  config.IdentifierFromProvider,
	"vsphere_distributed_virtual_switch":              config.IdentifierFromProvider,
	"vsphere_host_port_group":                         config.IdentifierFromProvider,
	"vsphere_host_virtual_switch":                     config.IdentifierFromProvider,
	"vsphere_entity_permissions":                      config.IdentifierFromProvider,
	"vsphere_role":                                    config.IdentifierFromProvider,
	"vsphere_datastore_cluster":                       config.IdentifierFromProvider,
	"vsphere_datastore_cluster_vm_anti_affinity_rule": config.IdentifierFromProvider,
	"vsphere_file":                                    config.IdentifierFromProvider,
	"vsphere_nas_datastore":                           config.IdentifierFromProvider,
	"vsphere_storage_drs_vm_override":                 config.IdentifierFromProvider,
	"vsphere_vm_storage_policy":                       config.IdentifierFromProvider,
	"vsphere_vmfs_datastore":                          config.IdentifierFromProvider,
	"vsphere_content_library":                         config.IdentifierFromProvider,
	"vsphere_content_library_item":                    config.IdentifierFromProvider,
	"vsphere_vapp_container":                          config.IdentifierFromProvider,
	"vsphere_vapp_entity":                             config.IdentifierFromProvider,
	"vsphere_virtual_disk":                            config.IdentifierFromProvider,
	"vsphere_virtual_machine":                         config.IdentifierFromProvider,
	"vsphere_virtual_machine_snapshot":                config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
