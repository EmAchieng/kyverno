/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

import (
	v1 "github.com/kyverno/kyverno/api/kyverno/v1"
	kyvernov1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1"
)

// SpecApplyConfiguration represents an declarative configuration of the Spec type for use
// with apply.
type SpecApplyConfiguration struct {
	Rules                            []RuleApplyConfiguration                                      `json:"rules,omitempty"`
	ApplyRules                       *v1.ApplyRulesType                                            `json:"applyRules,omitempty"`
	FailurePolicy                    *v1.FailurePolicyType                                         `json:"failurePolicy,omitempty"`
	ValidationFailureAction          *v1.ValidationFailureAction                                   `json:"validationFailureAction,omitempty"`
	ValidationFailureActionOverrides []kyvernov1.ValidationFailureActionOverrideApplyConfiguration `json:"validationFailureActionOverrides,omitempty"`
	Admission                        *bool                                                         `json:"admission,omitempty"`
	Background                       *bool                                                         `json:"background,omitempty"`
	SchemaValidation                 *bool                                                         `json:"schemaValidation,omitempty"`
	WebhookTimeoutSeconds            *int32                                                        `json:"webhookTimeoutSeconds,omitempty"`
	MutateExistingOnPolicyUpdate     *bool                                                         `json:"mutateExistingOnPolicyUpdate,omitempty"`
	GenerateExistingOnPolicyUpdate   *bool                                                         `json:"generateExistingOnPolicyUpdate,omitempty"`
	GenerateExisting                 *bool                                                         `json:"generateExisting,omitempty"`
	UseServerSideApply               *bool                                                         `json:"useServerSideApply,omitempty"`
	WebhookConfiguration             *WebhookConfigurationApplyConfiguration                       `json:"webhookConfiguration,omitempty"`
}

// SpecApplyConfiguration constructs an declarative configuration of the Spec type for use with
// apply.
func Spec() *SpecApplyConfiguration {
	return &SpecApplyConfiguration{}
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *SpecApplyConfiguration) WithRules(values ...*RuleApplyConfiguration) *SpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}

// WithApplyRules sets the ApplyRules field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ApplyRules field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithApplyRules(value v1.ApplyRulesType) *SpecApplyConfiguration {
	b.ApplyRules = &value
	return b
}

// WithFailurePolicy sets the FailurePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailurePolicy field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithFailurePolicy(value v1.FailurePolicyType) *SpecApplyConfiguration {
	b.FailurePolicy = &value
	return b
}

// WithValidationFailureAction sets the ValidationFailureAction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ValidationFailureAction field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithValidationFailureAction(value v1.ValidationFailureAction) *SpecApplyConfiguration {
	b.ValidationFailureAction = &value
	return b
}

// WithValidationFailureActionOverrides adds the given value to the ValidationFailureActionOverrides field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ValidationFailureActionOverrides field.
func (b *SpecApplyConfiguration) WithValidationFailureActionOverrides(values ...*kyvernov1.ValidationFailureActionOverrideApplyConfiguration) *SpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithValidationFailureActionOverrides")
		}
		b.ValidationFailureActionOverrides = append(b.ValidationFailureActionOverrides, *values[i])
	}
	return b
}

// WithAdmission sets the Admission field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Admission field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithAdmission(value bool) *SpecApplyConfiguration {
	b.Admission = &value
	return b
}

// WithBackground sets the Background field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Background field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithBackground(value bool) *SpecApplyConfiguration {
	b.Background = &value
	return b
}

// WithSchemaValidation sets the SchemaValidation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchemaValidation field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithSchemaValidation(value bool) *SpecApplyConfiguration {
	b.SchemaValidation = &value
	return b
}

// WithWebhookTimeoutSeconds sets the WebhookTimeoutSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookTimeoutSeconds field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithWebhookTimeoutSeconds(value int32) *SpecApplyConfiguration {
	b.WebhookTimeoutSeconds = &value
	return b
}

// WithMutateExistingOnPolicyUpdate sets the MutateExistingOnPolicyUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MutateExistingOnPolicyUpdate field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithMutateExistingOnPolicyUpdate(value bool) *SpecApplyConfiguration {
	b.MutateExistingOnPolicyUpdate = &value
	return b
}

// WithGenerateExistingOnPolicyUpdate sets the GenerateExistingOnPolicyUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateExistingOnPolicyUpdate field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithGenerateExistingOnPolicyUpdate(value bool) *SpecApplyConfiguration {
	b.GenerateExistingOnPolicyUpdate = &value
	return b
}

// WithGenerateExisting sets the GenerateExisting field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateExisting field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithGenerateExisting(value bool) *SpecApplyConfiguration {
	b.GenerateExisting = &value
	return b
}

// WithUseServerSideApply sets the UseServerSideApply field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UseServerSideApply field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithUseServerSideApply(value bool) *SpecApplyConfiguration {
	b.UseServerSideApply = &value
	return b
}

// WithWebhookConfiguration sets the WebhookConfiguration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookConfiguration field is set to the value of the last call.
func (b *SpecApplyConfiguration) WithWebhookConfiguration(value *WebhookConfigurationApplyConfiguration) *SpecApplyConfiguration {
	b.WebhookConfiguration = value
	return b
}
