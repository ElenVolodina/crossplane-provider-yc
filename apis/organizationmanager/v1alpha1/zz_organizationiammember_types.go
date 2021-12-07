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

type OrganizationIamMemberObservation struct {
}

type OrganizationIamMemberParameters struct {

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	OrganizationID *string `json:"organizationId" tf:"organization_id,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// OrganizationIamMemberSpec defines the desired state of OrganizationIamMember
type OrganizationIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationIamMemberParameters `json:"forProvider"`
}

// OrganizationIamMemberStatus defines the observed state of OrganizationIamMember.
type OrganizationIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIamMember is the Schema for the OrganizationIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type OrganizationIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIamMemberSpec   `json:"spec"`
	Status            OrganizationIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIamMemberList contains a list of OrganizationIamMembers
type OrganizationIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationIamMember `json:"items"`
}

// Repository type metadata.
var (
	OrganizationIamMember_Kind             = "OrganizationIamMember"
	OrganizationIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationIamMember_Kind}.String()
	OrganizationIamMember_KindAPIVersion   = OrganizationIamMember_Kind + "." + CRDGroupVersion.String()
	OrganizationIamMember_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationIamMember{}, &OrganizationIamMemberList{})
}