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

type VSphereHostVirtualSwitchObservation struct {

	// List of active network adapters used for load balancing.
	ActiveNics []*string `json:"activeNics,omitempty" tf:"active_nics,omitempty"`

	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits,omitempty"`

	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `json:"allowMacChanges,omitempty" tf:"allow_mac_changes,omitempty"`

	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous,omitempty"`

	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval *float64 `json:"beaconInterval,omitempty" tf:"beacon_interval,omitempty"`

	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon,omitempty"`

	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `json:"failback,omitempty" tf:"failback,omitempty"`

	// The managed object ID of the host to set the virtual switch up on.
	HostSystemID *string `json:"hostSystemId,omitempty" tf:"host_system_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation,omitempty"`

	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol,omitempty"`

	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the virtual switch.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []*string `json:"networkAdapters,omitempty" tf:"network_adapters,omitempty"`

	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches,omitempty"`

	// The number of ports that this virtual switch is configured to use.
	NumberOfPorts *float64 `json:"numberOfPorts,omitempty" tf:"number_of_ports,omitempty"`

	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth *float64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth,omitempty"`

	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize *float64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size,omitempty"`

	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled,omitempty"`

	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth *float64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth,omitempty"`

	// List of standby network adapters used for failover.
	StandbyNics []*string `json:"standbyNics,omitempty" tf:"standby_nics,omitempty"`

	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy,omitempty"`
}

type VSphereHostVirtualSwitchParameters struct {

	// List of active network adapters used for load balancing.
	// +kubebuilder:validation:Optional
	ActiveNics []*string `json:"activeNics,omitempty" tf:"active_nics,omitempty"`

	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	// +kubebuilder:validation:Optional
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits,omitempty"`

	// Controls whether or not the Media Access Control (MAC) address can be changed.
	// +kubebuilder:validation:Optional
	AllowMacChanges *bool `json:"allowMacChanges,omitempty" tf:"allow_mac_changes,omitempty"`

	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	// +kubebuilder:validation:Optional
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous,omitempty"`

	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	// +kubebuilder:validation:Optional
	BeaconInterval *float64 `json:"beaconInterval,omitempty" tf:"beacon_interval,omitempty"`

	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	// +kubebuilder:validation:Optional
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon,omitempty"`

	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	// +kubebuilder:validation:Optional
	Failback *bool `json:"failback,omitempty" tf:"failback,omitempty"`

	// The managed object ID of the host to set the virtual switch up on.
	// +kubebuilder:validation:Optional
	HostSystemID *string `json:"hostSystemId,omitempty" tf:"host_system_id,omitempty"`

	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	// +kubebuilder:validation:Optional
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation,omitempty"`

	// The discovery protocol type. Valid values are cdp and lldp.
	// +kubebuilder:validation:Optional
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol,omitempty"`

	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the virtual switch.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The list of network adapters to bind to this virtual switch.
	// +kubebuilder:validation:Optional
	NetworkAdapters []*string `json:"networkAdapters,omitempty" tf:"network_adapters,omitempty"`

	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	// +kubebuilder:validation:Optional
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches,omitempty"`

	// The number of ports that this virtual switch is configured to use.
	// +kubebuilder:validation:Optional
	NumberOfPorts *float64 `json:"numberOfPorts,omitempty" tf:"number_of_ports,omitempty"`

	// The average bandwidth in bits per second if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingAverageBandwidth *float64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth,omitempty"`

	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingBurstSize *float64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size,omitempty"`

	// Enable traffic shaping on this virtual switch or port group.
	// +kubebuilder:validation:Optional
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled,omitempty"`

	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	// +kubebuilder:validation:Optional
	ShapingPeakBandwidth *float64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth,omitempty"`

	// List of standby network adapters used for failover.
	// +kubebuilder:validation:Optional
	StandbyNics []*string `json:"standbyNics,omitempty" tf:"standby_nics,omitempty"`

	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	// +kubebuilder:validation:Optional
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy,omitempty"`
}

// VSphereHostVirtualSwitchSpec defines the desired state of VSphereHostVirtualSwitch
type VSphereHostVirtualSwitchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereHostVirtualSwitchParameters `json:"forProvider"`
}

// VSphereHostVirtualSwitchStatus defines the observed state of VSphereHostVirtualSwitch.
type VSphereHostVirtualSwitchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereHostVirtualSwitchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereHostVirtualSwitch is the Schema for the VSphereHostVirtualSwitchs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereHostVirtualSwitch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.activeNics)",message="activeNics is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.hostSystemId)",message="hostSystemId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkAdapters)",message="networkAdapters is a required parameter"
	Spec   VSphereHostVirtualSwitchSpec   `json:"spec"`
	Status VSphereHostVirtualSwitchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereHostVirtualSwitchList contains a list of VSphereHostVirtualSwitchs
type VSphereHostVirtualSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereHostVirtualSwitch `json:"items"`
}

// Repository type metadata.
var (
	VSphereHostVirtualSwitch_Kind             = "VSphereHostVirtualSwitch"
	VSphereHostVirtualSwitch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereHostVirtualSwitch_Kind}.String()
	VSphereHostVirtualSwitch_KindAPIVersion   = VSphereHostVirtualSwitch_Kind + "." + CRDGroupVersion.String()
	VSphereHostVirtualSwitch_GroupVersionKind = CRDGroupVersion.WithKind(VSphereHostVirtualSwitch_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereHostVirtualSwitch{}, &VSphereHostVirtualSwitchList{})
}
