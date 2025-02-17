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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type ConditionType string

// These are valid conditions of notification message. Currently, ready is the only condition type
// for notification message, it means that whether this message has been sent successfully.
const (
	// MessageReady means this message has been sent successfully
	MessageReady ConditionType = "Ready"
)

// NotificationMessageSpec defines the desired state of NotificationMessage
type NotificationMessageSpec struct {
	// Notifications associated by this message
	Notifications []NotificationName `json:"notifications"`
	// Message body, json in any form
	Body runtime.RawExtension `json:"body"`
}

type NotificationName struct {
	Name string `json:"name"`
}

// NotificationMessageStatus defines the observed state of NotificationMessage
type NotificationMessageStatus struct {
	Conditions []NotificationMessageCondition `json:"conditions,omitempty"`
	// +optional
	Phase NotificationMessagePhase `json:"phase,omitempty"`
}

// NotificationMessageCondition contains condition information for a notification message.
type NotificationMessageCondition struct {
	// Set condition for every receiver. Use Notification/Method/Namespace/Name to identify a receiver.
	// Notification
	// +optional
	Notification string `json:"notification"`
	// Method of receiver
	// +optional
	Method string `json:"method"`
	// Namespace of receiver
	// +optional
	ReceiverNamespace string `json:"receiverNamespace"`
	// Receiver name
	// +optional
	ReceiverName string `json:"receiverName"`
	// Type of notification message condition.
	Type ConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status"`
	// Retry times for this notification message.
	// +optional
	RetryTimes int `json:"retryTimes,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Brief reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

type NotificationMessagePhase string

// These are the valid phases of NotificationMessage.
const (
	// Pending means the message has been created, but not processed.
	NotificationMessagePending NotificationMessagePhase = "Pending"
	// Processing means the message is  in processing.
	NotificationMessageProcessing NotificationMessagePhase = "Processing"
	// Failed means the message has been processed and  failed for some reason.
	NotificationMessageFailed NotificationMessagePhase = "Failed"
	// Done means the message has been processed successfully.
	NotificationMessageDone NotificationMessagePhase = "Done"
)

// NoticationMessage is the Schema for the templates API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=nfm
type NotificationMessage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationMessageSpec   `json:"spec"`
	Status            NotificationMessageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationMessageList contains a list of NotificationMessage
type NotificationMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationMessage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationMessage{}, &NotificationMessageList{})
}
