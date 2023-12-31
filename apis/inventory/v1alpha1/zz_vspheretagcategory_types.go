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

type VSphereTagCategoryObservation struct {

	// Object types to which this category's tags can be attached. Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	AssociableTypes []*string `json:"associableTypes,omitempty" tf:"associable_types,omitempty"`

	// The associated cardinality of the category. Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	Cardinality *string `json:"cardinality,omitempty" tf:"cardinality,omitempty"`

	// The description of the category.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of the category.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VSphereTagCategoryParameters struct {

	// Object types to which this category's tags can be attached. Valid types include: Folder, ClusterComputeResource, Datacenter, Datastore, StoragePod, DistributedVirtualPortgroup, DistributedVirtualSwitch, VmwareDistributedVirtualSwitch, HostSystem, com.vmware.content.Library, com.vmware.content.library.Item, HostNetwork, Network, OpaqueNetwork, ResourcePool, VirtualApp, VirtualMachine.
	// +kubebuilder:validation:Optional
	AssociableTypes []*string `json:"associableTypes,omitempty" tf:"associable_types,omitempty"`

	// The associated cardinality of the category. Can be one of SINGLE (object can only be assigned one tag in this category) or MULTIPLE (object can be assigned multiple tags in this category).
	// +kubebuilder:validation:Optional
	Cardinality *string `json:"cardinality,omitempty" tf:"cardinality,omitempty"`

	// The description of the category.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the category.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// VSphereTagCategorySpec defines the desired state of VSphereTagCategory
type VSphereTagCategorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereTagCategoryParameters `json:"forProvider"`
}

// VSphereTagCategoryStatus defines the observed state of VSphereTagCategory.
type VSphereTagCategoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereTagCategoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereTagCategory is the Schema for the VSphereTagCategorys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereTagCategory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.associableTypes)",message="associableTypes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.cardinality)",message="cardinality is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   VSphereTagCategorySpec   `json:"spec"`
	Status VSphereTagCategoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereTagCategoryList contains a list of VSphereTagCategorys
type VSphereTagCategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereTagCategory `json:"items"`
}

// Repository type metadata.
var (
	VSphereTagCategory_Kind             = "VSphereTagCategory"
	VSphereTagCategory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereTagCategory_Kind}.String()
	VSphereTagCategory_KindAPIVersion   = VSphereTagCategory_Kind + "." + CRDGroupVersion.String()
	VSphereTagCategory_GroupVersionKind = CRDGroupVersion.WithKind(VSphereTagCategory_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereTagCategory{}, &VSphereTagCategoryList{})
}
