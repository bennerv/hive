// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// SelectorSyncSetSpecApplyConfiguration represents an declarative configuration of the SelectorSyncSetSpec type for use
// with apply.
type SelectorSyncSetSpecApplyConfiguration struct {
	SyncSetCommonSpecApplyConfiguration `json:",inline"`
	ClusterDeploymentSelector           *metav1.LabelSelector `json:"clusterDeploymentSelector,omitempty"`
}

// SelectorSyncSetSpecApplyConfiguration constructs an declarative configuration of the SelectorSyncSetSpec type for use with
// apply.
func SelectorSyncSetSpec() *SelectorSyncSetSpecApplyConfiguration {
	return &SelectorSyncSetSpecApplyConfiguration{}
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *SelectorSyncSetSpecApplyConfiguration) WithResources(values ...runtime.RawExtension) *SelectorSyncSetSpecApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithResourceApplyMode sets the ResourceApplyMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceApplyMode field is set to the value of the last call.
func (b *SelectorSyncSetSpecApplyConfiguration) WithResourceApplyMode(value hivev1.SyncSetResourceApplyMode) *SelectorSyncSetSpecApplyConfiguration {
	b.ResourceApplyMode = &value
	return b
}

// WithPatches adds the given value to the Patches field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Patches field.
func (b *SelectorSyncSetSpecApplyConfiguration) WithPatches(values ...*SyncObjectPatchApplyConfiguration) *SelectorSyncSetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPatches")
		}
		b.Patches = append(b.Patches, *values[i])
	}
	return b
}

// WithSecrets adds the given value to the Secrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Secrets field.
func (b *SelectorSyncSetSpecApplyConfiguration) WithSecrets(values ...*SecretMappingApplyConfiguration) *SelectorSyncSetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSecrets")
		}
		b.Secrets = append(b.Secrets, *values[i])
	}
	return b
}

// WithApplyBehavior sets the ApplyBehavior field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ApplyBehavior field is set to the value of the last call.
func (b *SelectorSyncSetSpecApplyConfiguration) WithApplyBehavior(value hivev1.SyncSetApplyBehavior) *SelectorSyncSetSpecApplyConfiguration {
	b.ApplyBehavior = &value
	return b
}

// WithEnableResourceTemplates sets the EnableResourceTemplates field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableResourceTemplates field is set to the value of the last call.
func (b *SelectorSyncSetSpecApplyConfiguration) WithEnableResourceTemplates(value bool) *SelectorSyncSetSpecApplyConfiguration {
	b.EnableResourceTemplates = &value
	return b
}

// WithClusterDeploymentSelector sets the ClusterDeploymentSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterDeploymentSelector field is set to the value of the last call.
func (b *SelectorSyncSetSpecApplyConfiguration) WithClusterDeploymentSelector(value metav1.LabelSelector) *SelectorSyncSetSpecApplyConfiguration {
	b.ClusterDeploymentSelector = &value
	return b
}
