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

type FolderIAMMemberObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FolderIAMMemberParameters struct {

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Required) ID of the folder to attach a policy to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/provider-jet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	// (Required) The identity that will be granted the privilege that is specified in the `role` field.
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The role that should be assigned.
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// FolderIAMMemberSpec defines the desired state of FolderIAMMember
type FolderIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIAMMemberParameters `json:"forProvider"`
}

// FolderIAMMemberStatus defines the observed state of FolderIAMMember.
type FolderIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMMember is the Schema for the FolderIAMMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type FolderIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderIAMMemberSpec   `json:"spec"`
	Status            FolderIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMMemberList contains a list of FolderIAMMembers
type FolderIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIAMMember `json:"items"`
}

// Repository type metadata.
var (
	FolderIAMMember_Kind             = "FolderIAMMember"
	FolderIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIAMMember_Kind}.String()
	FolderIAMMember_KindAPIVersion   = FolderIAMMember_Kind + "." + CRDGroupVersion.String()
	FolderIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(FolderIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIAMMember{}, &FolderIAMMemberList{})
}
