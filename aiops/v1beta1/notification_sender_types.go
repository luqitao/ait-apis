/*
Copyright 2019 Alauda AIOps Team.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO For compatibility, this resource will be deprecated in the future
// NotificationSenderSpec defines the desired state of NotificationSender
type NotificationSenderSpec struct {
	// Email sender configuration
	// +optional
	Email *Email `json:"email,omitempty"`
	// Sms sender configuration
	// +optional
	Sms *Sms `json:"sms,omitempty"`
	// Webhook sender configuration
	// +optional
	Webhook *Webhook `json:"webhook,omitempty"`
}

type Email struct {
	// Email sender's username
	Username string `json:"username"`
	// Email sender's password
	Password ValueWithSecret `json:"password"`
	// Email vendor's host
	// +optional
	Host string `json:"host"`
	// Email vendor's port
	Port int32 `json:"port"`
	// Enable ssl for connection, default is false
	// +optional
	SslEnabled bool `json:"ssl,omitempty"`
	// Skip insecure verify or not, default is false
	// +optional
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

type Sms struct {
	// AccountSid from sms vendor
	AccountSid ValueWithSecret `json:"accountSid"`
	// AccountToken from sms vendor
	AccountToken ValueWithSecret `json:"accountToken"`
	// Sms vendor's host
	Host string `json:"host"`
	// Sms vendor's port
	Port int32 `json:"port"`
	// Sms vendor's software version
	SoftVersion string `json:"softVersion"`
	// Sms sender's app id
	AppId string `json:"appId"`
	// Sms sender's template id
	// +optional
	TemplateId string `json:"templateId,omitempty"`
}

type Webhook struct {
	// Auth for this webhook
	// +optional
	Auth WebhookAuth `json:"auth,omitempty"`
	// Headers for this webhook
	// +optional
	Headers []KeyValueItem `json:"headers,omitempty"`
}

type WebhookAuth struct {
	// Basic auth configuration
	// +optional
	Basic BasicAuth `json:"basic,omitempty"`
}

type BasicAuth struct {
	// Basic auth username
	Username string `json:"username"`
	// Basic auth password
	Password ValueWithSecret `json:"password"`
}

type KeyValueItem struct {
	// Key for this dict item
	Key string `json:"key"`
	// Value for this dict item
	Value ValueWithSecret `json:"value"`
}

type ValueWithSecret struct {
	// Direct value for this field
	// +optional
	Value string `json:"value,omitempty"`
	// Reference value from secret for this field
	// +optional
	ValueFromSecret SecretInfo `json:"valueFromSecret,omitempty"`
}

type SecretInfo struct {
	// Secret's name
	Name string `json:"name"`
	// Secret's namespace
	Namespace string `json:"namespace"`
	// Key in secret
	Key string `json:"key"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationSender is the Schema for the Senders API
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=nfs,scope=Cluster
type NotificationSender struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationSenderSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationSenderList contains a list of NotificationSender
type NotificationSenderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationSender `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationSender{}, &NotificationSenderList{})
}
