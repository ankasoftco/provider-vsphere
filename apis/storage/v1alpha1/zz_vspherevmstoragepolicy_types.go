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

type TagRulesObservation struct {

	// Whether to include or exclude datastores tagged with the provided tags
	IncludeDatastoresWithTags *bool `json:"includeDatastoresWithTags,omitempty" tf:"include_datastores_with_tags,omitempty"`

	// The tag category to select the tags from.
	TagCategory *string `json:"tagCategory,omitempty" tf:"tag_category,omitempty"`

	// The tags to use for creating a tag-based vm placement rule.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TagRulesParameters struct {

	// Whether to include or exclude datastores tagged with the provided tags
	// +kubebuilder:validation:Optional
	IncludeDatastoresWithTags *bool `json:"includeDatastoresWithTags,omitempty" tf:"include_datastores_with_tags,omitempty"`

	// The tag category to select the tags from.
	// +kubebuilder:validation:Required
	TagCategory *string `json:"tagCategory" tf:"tag_category,omitempty"`

	// The tags to use for creating a tag-based vm placement rule.
	// +kubebuilder:validation:Required
	Tags []*string `json:"tags" tf:"tags,omitempty"`
}

type VSphereVmStoragePolicyObservation struct {

	// Description of the storage policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the storage policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tag rules to filter datastores to be used for placement of VMs.
	TagRules []TagRulesObservation `json:"tagRules,omitempty" tf:"tag_rules,omitempty"`
}

type VSphereVmStoragePolicyParameters struct {

	// Description of the storage policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the storage policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tag rules to filter datastores to be used for placement of VMs.
	// +kubebuilder:validation:Optional
	TagRules []TagRulesParameters `json:"tagRules,omitempty" tf:"tag_rules,omitempty"`
}

// VSphereVmStoragePolicySpec defines the desired state of VSphereVmStoragePolicy
type VSphereVmStoragePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereVmStoragePolicyParameters `json:"forProvider"`
}

// VSphereVmStoragePolicyStatus defines the observed state of VSphereVmStoragePolicy.
type VSphereVmStoragePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereVmStoragePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereVmStoragePolicy is the Schema for the VSphereVmStoragePolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereVmStoragePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.tagRules)",message="tagRules is a required parameter"
	Spec   VSphereVmStoragePolicySpec   `json:"spec"`
	Status VSphereVmStoragePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereVmStoragePolicyList contains a list of VSphereVmStoragePolicys
type VSphereVmStoragePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereVmStoragePolicy `json:"items"`
}

// Repository type metadata.
var (
	VSphereVmStoragePolicy_Kind             = "VSphereVmStoragePolicy"
	VSphereVmStoragePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereVmStoragePolicy_Kind}.String()
	VSphereVmStoragePolicy_KindAPIVersion   = VSphereVmStoragePolicy_Kind + "." + CRDGroupVersion.String()
	VSphereVmStoragePolicy_GroupVersionKind = CRDGroupVersion.WithKind(VSphereVmStoragePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereVmStoragePolicy{}, &VSphereVmStoragePolicyList{})
}
