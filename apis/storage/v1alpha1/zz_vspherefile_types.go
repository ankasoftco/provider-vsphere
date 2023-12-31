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

type VSphereFileObservation struct {
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	DestinationFile *string `json:"destinationFile,omitempty" tf:"destination_file,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SourceDatacenter *string `json:"sourceDatacenter,omitempty" tf:"source_datacenter,omitempty"`

	SourceDatastore *string `json:"sourceDatastore,omitempty" tf:"source_datastore,omitempty"`

	SourceFile *string `json:"sourceFile,omitempty" tf:"source_file,omitempty"`
}

type VSphereFileParameters struct {

	// +kubebuilder:validation:Optional
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// +kubebuilder:validation:Optional
	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationFile *string `json:"destinationFile,omitempty" tf:"destination_file,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDatacenter *string `json:"sourceDatacenter,omitempty" tf:"source_datacenter,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDatastore *string `json:"sourceDatastore,omitempty" tf:"source_datastore,omitempty"`

	// +kubebuilder:validation:Optional
	SourceFile *string `json:"sourceFile,omitempty" tf:"source_file,omitempty"`
}

// VSphereFileSpec defines the desired state of VSphereFile
type VSphereFileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereFileParameters `json:"forProvider"`
}

// VSphereFileStatus defines the observed state of VSphereFile.
type VSphereFileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereFileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereFile is the Schema for the VSphereFiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereFile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.datastore)",message="datastore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.destinationFile)",message="destinationFile is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sourceFile)",message="sourceFile is a required parameter"
	Spec   VSphereFileSpec   `json:"spec"`
	Status VSphereFileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereFileList contains a list of VSphereFiles
type VSphereFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereFile `json:"items"`
}

// Repository type metadata.
var (
	VSphereFile_Kind             = "VSphereFile"
	VSphereFile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereFile_Kind}.String()
	VSphereFile_KindAPIVersion   = VSphereFile_Kind + "." + CRDGroupVersion.String()
	VSphereFile_GroupVersionKind = CRDGroupVersion.WithKind(VSphereFile_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereFile{}, &VSphereFileList{})
}
