/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DefaultSecurityGroupObservation struct {
	// Creation timestamp of this security group.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Egress []EgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Ingress []IngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Name of this security group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Status of this security group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DefaultSecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

type EgressObservation struct {
	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EgressParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type IngressObservation struct {
	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IngressParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

// DefaultSecurityGroupSpec defines the desired state of DefaultSecurityGroup
type DefaultSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSecurityGroupParameters `json:"forProvider"`
}

// DefaultSecurityGroupStatus defines the observed state of DefaultSecurityGroup.
type DefaultSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroup is the Schema for the DefaultSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroupList contains a list of DefaultSecurityGroups
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DefaultSecurityGroup_Kind             = "DefaultSecurityGroup"
	DefaultSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSecurityGroup_Kind}.String()
	DefaultSecurityGroup_KindAPIVersion   = DefaultSecurityGroup_Kind + "." + CRDGroupVersion.String()
	DefaultSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSecurityGroup{}, &DefaultSecurityGroupList{})
}
