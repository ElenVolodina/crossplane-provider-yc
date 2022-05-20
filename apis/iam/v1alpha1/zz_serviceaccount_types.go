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

type ServiceAccountObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceAccountParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Description of the service account.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) ID of the folder that the service account will be created in.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) Name of the service account.
	Name *string `json:"name" tf:"name,omitempty"`
}

// ServiceAccountSpec defines the desired state of ServiceAccount
type ServiceAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountParameters `json:"forProvider"`
}

// ServiceAccountStatus defines the observed state of ServiceAccount.
type ServiceAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccount is the Schema for the ServiceAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountSpec   `json:"spec"`
	Status            ServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountList contains a list of ServiceAccounts
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccount_Kind             = "ServiceAccount"
	ServiceAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccount_Kind}.String()
	ServiceAccount_KindAPIVersion   = ServiceAccount_Kind + "." + CRDGroupVersion.String()
	ServiceAccount_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}
