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

package v2alpha1

import (
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	v1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExternalAPICallApplyConfiguration represents an declarative configuration of the ExternalAPICall type for use
// with apply.
type ExternalAPICallApplyConfiguration struct {
	v1.APICallApplyConfiguration `json:",omitempty,inline"`
	RefreshInterval              *metav1.Duration `json:"refreshInterval,omitempty"`
}

// ExternalAPICallApplyConfiguration constructs an declarative configuration of the ExternalAPICall type for use with
// apply.
func ExternalAPICall() *ExternalAPICallApplyConfiguration {
	return &ExternalAPICallApplyConfiguration{}
}

// WithURLPath sets the URLPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URLPath field is set to the value of the last call.
func (b *ExternalAPICallApplyConfiguration) WithURLPath(value string) *ExternalAPICallApplyConfiguration {
	b.URLPath = &value
	return b
}

// WithMethod sets the Method field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Method field is set to the value of the last call.
func (b *ExternalAPICallApplyConfiguration) WithMethod(value kyvernov1.Method) *ExternalAPICallApplyConfiguration {
	b.Method = &value
	return b
}

// WithData adds the given value to the Data field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Data field.
func (b *ExternalAPICallApplyConfiguration) WithData(values ...*v1.RequestDataApplyConfiguration) *ExternalAPICallApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithData")
		}
		b.Data = append(b.Data, *values[i])
	}
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *ExternalAPICallApplyConfiguration) WithService(value *v1.ServiceCallApplyConfiguration) *ExternalAPICallApplyConfiguration {
	b.Service = value
	return b
}

// WithJMESPath sets the JMESPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JMESPath field is set to the value of the last call.
func (b *ExternalAPICallApplyConfiguration) WithJMESPath(value string) *ExternalAPICallApplyConfiguration {
	b.JMESPath = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *ExternalAPICallApplyConfiguration) WithRefreshInterval(value metav1.Duration) *ExternalAPICallApplyConfiguration {
	b.RefreshInterval = &value
	return b
}
