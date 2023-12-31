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

type VSphereCustomAttributeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object type for which the custom attribute is valid. If not specified, the attribute is valid for all managed object types.
	ManagedObjectType *string `json:"managedObjectType,omitempty" tf:"managed_object_type,omitempty"`

	// The display name of the custom attribute.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VSphereCustomAttributeParameters struct {

	// Object type for which the custom attribute is valid. If not specified, the attribute is valid for all managed object types.
	// +kubebuilder:validation:Optional
	ManagedObjectType *string `json:"managedObjectType,omitempty" tf:"managed_object_type,omitempty"`

	// The display name of the custom attribute.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// VSphereCustomAttributeSpec defines the desired state of VSphereCustomAttribute
type VSphereCustomAttributeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereCustomAttributeParameters `json:"forProvider"`
}

// VSphereCustomAttributeStatus defines the observed state of VSphereCustomAttribute.
type VSphereCustomAttributeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereCustomAttributeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereCustomAttribute is the Schema for the VSphereCustomAttributes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereCustomAttribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   VSphereCustomAttributeSpec   `json:"spec"`
	Status VSphereCustomAttributeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereCustomAttributeList contains a list of VSphereCustomAttributes
type VSphereCustomAttributeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereCustomAttribute `json:"items"`
}

// Repository type metadata.
var (
	VSphereCustomAttribute_Kind             = "VSphereCustomAttribute"
	VSphereCustomAttribute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereCustomAttribute_Kind}.String()
	VSphereCustomAttribute_KindAPIVersion   = VSphereCustomAttribute_Kind + "." + CRDGroupVersion.String()
	VSphereCustomAttribute_GroupVersionKind = CRDGroupVersion.WithKind(VSphereCustomAttribute_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereCustomAttribute{}, &VSphereCustomAttributeList{})
}
