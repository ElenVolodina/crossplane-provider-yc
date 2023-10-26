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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SymmetricKeyInitParameters struct {

	// Encryption algorithm to be used with a new key version,
	// generated with the next rotation. The default value is AES_128.
	DefaultAlgorithm *string `json:"defaultAlgorithm,omitempty" tf:"default_algorithm,omitempty"`

	// An optional description of the key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the key.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Interval between automatic rotations. To disable automatic rotation, omit this parameter.
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`
}

type SymmetricKeyObservation struct {

	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Encryption algorithm to be used with a new key version,
	// generated with the next rotation. The default value is AES_128.
	DefaultAlgorithm *string `json:"defaultAlgorithm,omitempty" tf:"default_algorithm,omitempty"`

	// An optional description of the key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the key.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Last rotation timestamp of the key.
	RotatedAt *string `json:"rotatedAt,omitempty" tf:"rotated_at,omitempty"`

	// Interval between automatic rotations. To disable automatic rotation, omit this parameter.
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// The status of the key.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SymmetricKeyParameters struct {

	// Encryption algorithm to be used with a new key version,
	// generated with the next rotation. The default value is AES_128.
	// +kubebuilder:validation:Optional
	DefaultAlgorithm *string `json:"defaultAlgorithm,omitempty" tf:"default_algorithm,omitempty"`

	// An optional description of the key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the key.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Interval between automatic rotations. To disable automatic rotation, omit this parameter.
	// +kubebuilder:validation:Optional
	RotationPeriod *string `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`
}

// SymmetricKeySpec defines the desired state of SymmetricKey
type SymmetricKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SymmetricKeyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SymmetricKeyInitParameters `json:"initProvider,omitempty"`
}

// SymmetricKeyStatus defines the observed state of SymmetricKey.
type SymmetricKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SymmetricKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SymmetricKey is the Schema for the SymmetricKeys API. Creates a Yandex KMS symmetric key that can be used for cryptographic operation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SymmetricKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SymmetricKeySpec   `json:"spec"`
	Status            SymmetricKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SymmetricKeyList contains a list of SymmetricKeys
type SymmetricKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SymmetricKey `json:"items"`
}

// Repository type metadata.
var (
	SymmetricKey_Kind             = "SymmetricKey"
	SymmetricKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SymmetricKey_Kind}.String()
	SymmetricKey_KindAPIVersion   = SymmetricKey_Kind + "." + CRDGroupVersion.String()
	SymmetricKey_GroupVersionKind = CRDGroupVersion.WithKind(SymmetricKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SymmetricKey{}, &SymmetricKeyList{})
}
