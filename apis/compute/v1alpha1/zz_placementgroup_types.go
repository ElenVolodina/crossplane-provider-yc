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

type PlacementGroupObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
}

type PlacementGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// PlacementGroupSpec defines the desired state of PlacementGroup
type PlacementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlacementGroupParameters `json:"forProvider"`
}

// PlacementGroupStatus defines the observed state of PlacementGroup.
type PlacementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlacementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroup is the Schema for the PlacementGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type PlacementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlacementGroupSpec   `json:"spec"`
	Status            PlacementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroupList contains a list of PlacementGroups
type PlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementGroup `json:"items"`
}

// Repository type metadata.
var (
	PlacementGroup_Kind             = "PlacementGroup"
	PlacementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlacementGroup_Kind}.String()
	PlacementGroup_KindAPIVersion   = PlacementGroup_Kind + "." + CRDGroupVersion.String()
	PlacementGroup_GroupVersionKind = CRDGroupVersion.WithKind(PlacementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PlacementGroup{}, &PlacementGroupList{})
}