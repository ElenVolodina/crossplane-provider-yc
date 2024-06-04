/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

type KafkaTopicInitParameters struct {

	// +crossplane:generate:reference:type=KafkaCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of the topic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of the topic's partitions.
	Partitions *float64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// User-defined settings for the topic. The structure is documented below.
	TopicConfig []KafkaTopicTopicConfigInitParameters `json:"topicConfig,omitempty" tf:"topic_config,omitempty"`
}

type KafkaTopicObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the topic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of the topic's partitions.
	Partitions *float64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Amount of data copies (replicas) for the topic in the cluster.
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// User-defined settings for the topic. The structure is documented below.
	TopicConfig []KafkaTopicTopicConfigObservation `json:"topicConfig,omitempty" tf:"topic_config,omitempty"`
}

type KafkaTopicParameters struct {

	// +crossplane:generate:reference:type=KafkaCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of the topic.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of the topic's partitions.
	// +kubebuilder:validation:Optional
	Partitions *float64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Amount of data copies (replicas) for the topic in the cluster.
	// +kubebuilder:validation:Optional
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// User-defined settings for the topic. The structure is documented below.
	// +kubebuilder:validation:Optional
	TopicConfig []KafkaTopicTopicConfigParameters `json:"topicConfig,omitempty" tf:"topic_config,omitempty"`
}

type KafkaTopicTopicConfigInitParameters struct {

	// Kafka topic settings. For more information, see
	// the official documentation
	// and the Kafka documentation.
	CleanupPolicy *string `json:"cleanupPolicy,omitempty" tf:"cleanup_policy,omitempty"`

	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	DeleteRetentionMs *string `json:"deleteRetentionMs,omitempty" tf:"delete_retention_ms,omitempty"`

	FileDeleteDelayMs *string `json:"fileDeleteDelayMs,omitempty" tf:"file_delete_delay_ms,omitempty"`

	FlushMessages *string `json:"flushMessages,omitempty" tf:"flush_messages,omitempty"`

	FlushMs *string `json:"flushMs,omitempty" tf:"flush_ms,omitempty"`

	MaxMessageBytes *string `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	MinCompactionLagMs *string `json:"minCompactionLagMs,omitempty" tf:"min_compaction_lag_ms,omitempty"`

	MinInsyncReplicas *string `json:"minInsyncReplicas,omitempty" tf:"min_insync_replicas,omitempty"`

	Preallocate *bool `json:"preallocate,omitempty" tf:"preallocate,omitempty"`

	RetentionBytes *string `json:"retentionBytes,omitempty" tf:"retention_bytes,omitempty"`

	RetentionMs *string `json:"retentionMs,omitempty" tf:"retention_ms,omitempty"`

	SegmentBytes *string `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

type KafkaTopicTopicConfigObservation struct {

	// Kafka topic settings. For more information, see
	// the official documentation
	// and the Kafka documentation.
	CleanupPolicy *string `json:"cleanupPolicy,omitempty" tf:"cleanup_policy,omitempty"`

	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	DeleteRetentionMs *string `json:"deleteRetentionMs,omitempty" tf:"delete_retention_ms,omitempty"`

	FileDeleteDelayMs *string `json:"fileDeleteDelayMs,omitempty" tf:"file_delete_delay_ms,omitempty"`

	FlushMessages *string `json:"flushMessages,omitempty" tf:"flush_messages,omitempty"`

	FlushMs *string `json:"flushMs,omitempty" tf:"flush_ms,omitempty"`

	MaxMessageBytes *string `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	MinCompactionLagMs *string `json:"minCompactionLagMs,omitempty" tf:"min_compaction_lag_ms,omitempty"`

	MinInsyncReplicas *string `json:"minInsyncReplicas,omitempty" tf:"min_insync_replicas,omitempty"`

	Preallocate *bool `json:"preallocate,omitempty" tf:"preallocate,omitempty"`

	RetentionBytes *string `json:"retentionBytes,omitempty" tf:"retention_bytes,omitempty"`

	RetentionMs *string `json:"retentionMs,omitempty" tf:"retention_ms,omitempty"`

	SegmentBytes *string `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

type KafkaTopicTopicConfigParameters struct {

	// Kafka topic settings. For more information, see
	// the official documentation
	// and the Kafka documentation.
	// +kubebuilder:validation:Optional
	CleanupPolicy *string `json:"cleanupPolicy,omitempty" tf:"cleanup_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteRetentionMs *string `json:"deleteRetentionMs,omitempty" tf:"delete_retention_ms,omitempty"`

	// +kubebuilder:validation:Optional
	FileDeleteDelayMs *string `json:"fileDeleteDelayMs,omitempty" tf:"file_delete_delay_ms,omitempty"`

	// +kubebuilder:validation:Optional
	FlushMessages *string `json:"flushMessages,omitempty" tf:"flush_messages,omitempty"`

	// +kubebuilder:validation:Optional
	FlushMs *string `json:"flushMs,omitempty" tf:"flush_ms,omitempty"`

	// +kubebuilder:validation:Optional
	MaxMessageBytes *string `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	MinCompactionLagMs *string `json:"minCompactionLagMs,omitempty" tf:"min_compaction_lag_ms,omitempty"`

	// +kubebuilder:validation:Optional
	MinInsyncReplicas *string `json:"minInsyncReplicas,omitempty" tf:"min_insync_replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Preallocate *bool `json:"preallocate,omitempty" tf:"preallocate,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionBytes *string `json:"retentionBytes,omitempty" tf:"retention_bytes,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionMs *string `json:"retentionMs,omitempty" tf:"retention_ms,omitempty"`

	// +kubebuilder:validation:Optional
	SegmentBytes *string `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

// KafkaTopicSpec defines the desired state of KafkaTopic
type KafkaTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaTopicParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider KafkaTopicInitParameters `json:"initProvider,omitempty"`
}

// KafkaTopicStatus defines the observed state of KafkaTopic.
type KafkaTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KafkaTopic is the Schema for the KafkaTopics API. Manages a topic of a Kafka cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitions) || (has(self.initProvider) && has(self.initProvider.partitions))",message="spec.forProvider.partitions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicationFactor) || (has(self.initProvider) && has(self.initProvider.replicationFactor))",message="spec.forProvider.replicationFactor is a required parameter"
	Spec   KafkaTopicSpec   `json:"spec"`
	Status KafkaTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopicList contains a list of KafkaTopics
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}

// Repository type metadata.
var (
	KafkaTopic_Kind             = "KafkaTopic"
	KafkaTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaTopic_Kind}.String()
	KafkaTopic_KindAPIVersion   = KafkaTopic_Kind + "." + CRDGroupVersion.String()
	KafkaTopic_GroupVersionKind = CRDGroupVersion.WithKind(KafkaTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
}
