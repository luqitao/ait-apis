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

// NotificationServerSpec defines the desired state of NotificationServer
type NotificationServerSpec struct {
	// Email server configuration
	// +optional
	Email *EmailServer `json:"email,omitempty"`
	// Sms server configuration
	// +optional
	Sms *SmsServer `json:"sms,omitempty"`
}

type EmailServer struct {
	// Email vendor's host
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

type SmsServer struct {
	// Sms vendor's host
	Host string `json:"host"`
	// Sms vendor's port
	Port int32 `json:"port"`
	// Sms vendor's software version
	SoftVersion string `json:"softVersion"`
	// Sms server's template id
	// +optional
	TemplateId string `json:"templateId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationServer is the Schema for the Servers API
// +k8s:openapi-gen=true
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
type NotificationServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationServerSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationServerList contains a list of NotificationServer
type NotificationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationServer{}, &NotificationServerList{})
}
