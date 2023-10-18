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

type VSphereDatacenterObservation struct {

	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Managed object ID of the datacenter.
	Moid *string `json:"moid,omitempty" tf:"moid,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of tag IDs to apply to this object.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VSphereDatacenterParameters struct {

	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VSphereDatacenterSpec defines the desired state of VSphereDatacenter
type VSphereDatacenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereDatacenterParameters `json:"forProvider"`
}

// VSphereDatacenterStatus defines the observed state of VSphereDatacenter.
type VSphereDatacenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereDatacenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereDatacenter is the Schema for the VSphereDatacenters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereDatacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   VSphereDatacenterSpec   `json:"spec"`
	Status VSphereDatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereDatacenterList contains a list of VSphereDatacenters
type VSphereDatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereDatacenter `json:"items"`
}

// Repository type metadata.
var (
	VSphereDatacenter_Kind             = "VSphereDatacenter"
	VSphereDatacenter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereDatacenter_Kind}.String()
	VSphereDatacenter_KindAPIVersion   = VSphereDatacenter_Kind + "." + CRDGroupVersion.String()
	VSphereDatacenter_GroupVersionKind = CRDGroupVersion.WithKind(VSphereDatacenter_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereDatacenter{}, &VSphereDatacenterList{})
}