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

type SecurityGroupRuleInitParameters struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// direction of the rule. Can be ingress (inbound) or egress (outbound).
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets such as "self_security_group". See docs for possible options.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the security group this rule belongs to.
	SecurityGroupBinding *string `json:"securityGroupBinding,omitempty" tf:"security_group_binding,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupRuleObservation struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// direction of the rule. Can be ingress (inbound) or egress (outbound).
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this rule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets such as "self_security_group". See docs for possible options.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the security group this rule belongs to.
	SecurityGroupBinding *string `json:"securityGroupBinding,omitempty" tf:"security_group_binding,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupRuleParameters struct {

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// direction of the rule. Can be ingress (inbound) or egress (outbound).
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Minimum port number.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets such as "self_security_group". See docs for possible options.
	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the security group this rule belongs to.
	// +kubebuilder:validation:Optional
	SecurityGroupBinding *string `json:"securityGroupBinding,omitempty" tf:"security_group_binding,omitempty"`

	// Target security group ID for this rule.
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Maximum port number.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

// SecurityGroupRuleSpec defines the desired state of SecurityGroupRule
type SecurityGroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupRuleParameters `json:"forProvider"`
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
	InitProvider SecurityGroupRuleInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupRuleStatus defines the observed state of SecurityGroupRule.
type SecurityGroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule is the Schema for the SecurityGroupRules API. Yandex VPC Security Group Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.direction) || has(self.initProvider.direction)",message="direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.securityGroupBinding) || has(self.initProvider.securityGroupBinding)",message="securityGroupBinding is a required parameter"
	Spec   SecurityGroupRuleSpec   `json:"spec"`
	Status SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRuleList contains a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupRule_Kind             = "SecurityGroupRule"
	SecurityGroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupRule_Kind}.String()
	SecurityGroupRule_KindAPIVersion   = SecurityGroupRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupRule{}, &SecurityGroupRuleList{})
}
