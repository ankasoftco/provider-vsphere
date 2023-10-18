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
	vspherecustomattribute "github.com/ankasoftco/provider-vsphere/internal/controller/inventory/vspherecustomattribute"
	vspheredatacenter "github.com/ankasoftco/provider-vsphere/internal/controller/inventory/vspheredatacenter"
	vspherefolder "github.com/ankasoftco/provider-vsphere/internal/controller/inventory/vspherefolder"
	vspheretag "github.com/ankasoftco/provider-vsphere/internal/controller/inventory/vspheretag"
	vspheretagcategory "github.com/ankasoftco/provider-vsphere/internal/controller/inventory/vspheretagcategory"
	vspheredistributedportgroup "github.com/ankasoftco/provider-vsphere/internal/controller/networking/vspheredistributedportgroup"
	vspheredistributedvirtualswitch "github.com/ankasoftco/provider-vsphere/internal/controller/networking/vspheredistributedvirtualswitch"
	vspherehostportgroup "github.com/ankasoftco/provider-vsphere/internal/controller/networking/vspherehostportgroup"
	vspherehostvirtualswitch "github.com/ankasoftco/provider-vsphere/internal/controller/networking/vspherehostvirtualswitch"
	providerconfig "github.com/ankasoftco/provider-vsphere/internal/controller/providerconfig"
	vsphereentitypermissions "github.com/ankasoftco/provider-vsphere/internal/controller/security/vsphereentitypermissions"
	vsphererole "github.com/ankasoftco/provider-vsphere/internal/controller/security/vsphererole"
	vspheredatastorecluster "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspheredatastorecluster"
	vspheredatastoreclustervmantiaffinityrule "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspheredatastoreclustervmantiaffinityrule"
	vspherefile "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspherefile"
	vspherenasdatastore "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspherenasdatastore"
	vspherestoragedrsvmoverride "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspherestoragedrsvmoverride"
	vspherevmfsdatastore "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspherevmfsdatastore"
	vspherevmstoragepolicy "github.com/ankasoftco/provider-vsphere/internal/controller/storage/vspherevmstoragepolicy"
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
		vspherecustomattribute.Setup,
		vspheredatacenter.Setup,
		vspherefolder.Setup,
		vspheretag.Setup,
		vspheretagcategory.Setup,
		vspheredistributedportgroup.Setup,
		vspheredistributedvirtualswitch.Setup,
		vspherehostportgroup.Setup,
		vspherehostvirtualswitch.Setup,
		providerconfig.Setup,
		vsphereentitypermissions.Setup,
		vsphererole.Setup,
		vspheredatastorecluster.Setup,
		vspheredatastoreclustervmantiaffinityrule.Setup,
		vspherefile.Setup,
		vspherenasdatastore.Setup,
		vspherestoragedrsvmoverride.Setup,
		vspherevmfsdatastore.Setup,
		vspherevmstoragepolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
