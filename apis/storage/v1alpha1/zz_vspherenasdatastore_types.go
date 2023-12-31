/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VSphereNasDatastoreObservation struct {

	// Access mode for the mount point. Can be one of readOnly or readWrite.
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// The connectivity status of the datastore. If this is false, some other computed attributes may be out of date.
	Accessible *bool `json:"accessible,omitempty" tf:"accessible,omitempty"`

	// Maximum capacity of the datastore, in MB.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The managed object ID of the datastore cluster to place the datastore in.
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	// The path to the datastore folder to put the datastore in.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// Available space of this datastore, in MB.
	FreeSpace *float64 `json:"freeSpace,omitempty" tf:"free_space,omitempty"`

	// The managed object IDs of the hosts to mount the datastore on.
	HostSystemIds []*string `json:"hostSystemIds,omitempty" tf:"host_system_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current maintenance mode state of the datastore.
	MaintenanceMode *string `json:"maintenanceMode,omitempty" tf:"maintenance_mode,omitempty"`

	// If true, more than one host in the datacenter has been configured with access to the datastore.
	MultipleHostAccess *bool `json:"multipleHostAccess,omitempty" tf:"multiple_host_access,omitempty"`

	// The name of the datastore.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates that this NAS volume is a protocol endpoint. This field is only populated if the host supports virtual datastores.
	ProtocolEndpoint *bool `json:"protocolEndpoint,omitempty" tf:"protocol_endpoint,omitempty"`

	// The hostnames or IP addresses of the remote server or servers. Only one element should be present for NFS v3 but multiple can be present for NFS v4.1.
	RemoteHosts []*string `json:"remoteHosts,omitempty" tf:"remote_hosts,omitempty"`

	// The remote path of the mount point.
	RemotePath *string `json:"remotePath,omitempty" tf:"remote_path,omitempty"`

	// The security type to use.
	SecurityType *string `json:"securityType,omitempty" tf:"security_type,omitempty"`

	// A list of tag IDs to apply to this object.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of NAS volume. Can be one of NFS (to denote v3) or NFS41 (to denote NFS v4.1).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The unique locator for the datastore.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Total additional storage space, in MB, potentially used by all virtual machines on this datastore.
	UncommittedSpace *float64 `json:"uncommittedSpace,omitempty" tf:"uncommitted_space,omitempty"`
}

type VSphereNasDatastoreParameters struct {

	// Access mode for the mount point. Can be one of readOnly or readWrite.
	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The managed object ID of the datastore cluster to place the datastore in.
	// +kubebuilder:validation:Optional
	DatastoreClusterID *string `json:"datastoreClusterId,omitempty" tf:"datastore_cluster_id,omitempty"`

	// The path to the datastore folder to put the datastore in.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The managed object IDs of the hosts to mount the datastore on.
	// +kubebuilder:validation:Optional
	HostSystemIds []*string `json:"hostSystemIds,omitempty" tf:"host_system_ids,omitempty"`

	// The name of the datastore.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The hostnames or IP addresses of the remote server or servers. Only one element should be present for NFS v3 but multiple can be present for NFS v4.1.
	// +kubebuilder:validation:Optional
	RemoteHosts []*string `json:"remoteHosts,omitempty" tf:"remote_hosts,omitempty"`

	// The remote path of the mount point.
	// +kubebuilder:validation:Optional
	RemotePath *string `json:"remotePath,omitempty" tf:"remote_path,omitempty"`

	// The security type to use.
	// +kubebuilder:validation:Optional
	SecurityType *string `json:"securityType,omitempty" tf:"security_type,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of NAS volume. Can be one of NFS (to denote v3) or NFS41 (to denote NFS v4.1).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// VSphereNasDatastoreSpec defines the desired state of VSphereNasDatastore
type VSphereNasDatastoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereNasDatastoreParameters `json:"forProvider"`
}

// VSphereNasDatastoreStatus defines the observed state of VSphereNasDatastore.
type VSphereNasDatastoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereNasDatastoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereNasDatastore is the Schema for the VSphereNasDatastores API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereNasDatastore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.hostSystemIds)",message="hostSystemIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.remoteHosts)",message="remoteHosts is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.remotePath)",message="remotePath is a required parameter"
	Spec   VSphereNasDatastoreSpec   `json:"spec"`
	Status VSphereNasDatastoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereNasDatastoreList contains a list of VSphereNasDatastores
type VSphereNasDatastoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereNasDatastore `json:"items"`
}

// Repository type metadata.
var (
	VSphereNasDatastore_Kind             = "VSphereNasDatastore"
	VSphereNasDatastore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereNasDatastore_Kind}.String()
	VSphereNasDatastore_KindAPIVersion   = VSphereNasDatastore_Kind + "." + CRDGroupVersion.String()
	VSphereNasDatastore_GroupVersionKind = CRDGroupVersion.WithKind(VSphereNasDatastore_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereNasDatastore{}, &VSphereNasDatastoreList{})
}
