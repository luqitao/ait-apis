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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	// Subscriptions for this notification
	Subscriptions []NotificationSubscription `json:"subscriptions,omitempty"`
}

type NotificationSubscription struct {
	// Method of recipient, such as email/dingding/sms/webhook
	Method string `json:"method"`
	// Recipients, multiple recipients will be separated by comma. This field will be deprecated
	// +optional
	Recipient string `json:"recipient,omitempty"`
	// Receivers, at least one notification receiver should be added
	// +optional
	Receivers []NamespacedName `json:"receivers,omitempty"`
	// Sender for this receiver
	// +optional
	Sender string `json:"sender,omitempty"`
	// Template to render notification message
	// +optional
	Template string `json:"template,omitempty"`
}

type NamespacedName struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Notification is the Schema for the notifications API
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=nf
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationList contains a list of Notification
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
