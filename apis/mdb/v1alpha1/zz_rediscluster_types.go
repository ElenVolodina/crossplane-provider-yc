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

type RedisClusterConfigObservation struct {
}

type RedisClusterConfigParameters struct {

	// +kubebuilder:validation:Optional
	Databases *int64 `json:"databases,omitempty" tf:"databases,omitempty"`

	// +kubebuilder:validation:Optional
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// +kubebuilder:validation:Optional
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	SlowlogLogSlowerThan *int64 `json:"slowlogLogSlowerThan,omitempty" tf:"slowlog_log_slower_than,omitempty"`

	// +kubebuilder:validation:Optional
	SlowlogMaxLen *int64 `json:"slowlogMaxLen,omitempty" tf:"slowlog_max_len,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type RedisClusterHostObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type RedisClusterHostParameters struct {

	// +kubebuilder:validation:Optional
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type RedisClusterMaintenanceWindowObservation struct {
}

type RedisClusterMaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	Hour *int64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedisClusterObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RedisClusterParameters struct {

	// +kubebuilder:validation:Required
	Config []RedisClusterConfigParameters `json:"config" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +crossplane:generate:reference:type=bb.yandex-team.ru/crossplane/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Host []RedisClusterHostParameters `json:"host" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []RedisClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=bb.yandex-team.ru/crossplane/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Resources []RedisClusterResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// +kubebuilder:validation:Optional
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type RedisClusterResourcesObservation struct {
}

type RedisClusterResourcesParameters struct {

	// +kubebuilder:validation:Required
	DiskSize *int64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster
type RedisClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisClusterParameters `json:"forProvider"`
}

// RedisClusterStatus defines the observed state of RedisCluster.
type RedisClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCluster is the Schema for the RedisClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisClusterSpec   `json:"spec"`
	Status            RedisClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisClusterList contains a list of RedisClusters
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisCluster_Kind             = "RedisCluster"
	RedisCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCluster_Kind}.String()
	RedisCluster_KindAPIVersion   = RedisCluster_Kind + "." + CRDGroupVersion.String()
	RedisCluster_GroupVersionKind = CRDGroupVersion.WithKind(RedisCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
}
