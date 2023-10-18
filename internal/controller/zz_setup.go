/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	vspherelicense "github.com/ankasoftco/provider-vsphere/internal/controller/administration/vspherelicense"
	vspherecomputecluster "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputecluster"
	vspherecomputeclusterhostgroup "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclusterhostgroup"
	vspherecomputeclustervmaffinityrule "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclustervmaffinityrule"
	vspherecomputeclustervmantiaffinityrule "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclustervmantiaffinityrule"
	vspherecomputeclustervmdependencyrule "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclustervmdependencyrule"
	vspherecomputeclustervmgroup "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclustervmgroup"
	vspherecomputeclustervmhostrule "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherecomputeclustervmhostrule"
	vspheredpmhostoverride "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspheredpmhostoverride"
	vspheredrsvmoverride "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspheredrsvmoverride"
	vspherehavmoverride "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherehavmoverride"
	vspherehost "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherehost"
	vsphereresourcepool "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vsphereresourcepool"
	vspherevnic "github.com/ankasoftco/provider-vsphere/internal/controller/host_and_cluster_management/vspherevnic"
	providerconfig "github.com/ankasoftco/provider-vsphere/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vspherelicense.Setup,
		vspherecomputecluster.Setup,
		vspherecomputeclusterhostgroup.Setup,
		vspherecomputeclustervmaffinityrule.Setup,
		vspherecomputeclustervmantiaffinityrule.Setup,
		vspherecomputeclustervmdependencyrule.Setup,
		vspherecomputeclustervmgroup.Setup,
		vspherecomputeclustervmhostrule.Setup,
		vspheredpmhostoverride.Setup,
		vspheredrsvmoverride.Setup,
		vspherehavmoverride.Setup,
		vspherehost.Setup,
		vsphereresourcepool.Setup,
		vspherevnic.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
