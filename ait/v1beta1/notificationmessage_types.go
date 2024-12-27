/*


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
	mv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	Group        = "ait.alauda.io"
	Version      = "v1beta1"
	GroupVersion = Group + "/" + Version
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

type Phase string

// These are the valid phases of NotificationMessage.
const (
	// Pending means the message has been created, but not processed.
	PhasePending Phase = "Pending"
	// Processing means the message is  in processing.
	PhaseProcessing Phase = "Processing"
	// Failed means the message has been processed and  failed for some reason.
	PhaseFailed Phase = "Failed"
	// Done means the message has been processed successfully.
	PhaseDone Phase = "Done"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationMessage
// +k8s:openapi-gen=true
type NotificationMessage struct {
	mv1.TypeMeta   `json:",inline"`
	mv1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationMessageSpec   `json:"spec,omitempty"`
	Status NotificationMessageStatus `json:"status,omitempty"`
}

// NotificationMessageList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NotificationMessageList struct {
	mv1.TypeMeta `json:",inline"`
	mv1.ListMeta `json:"metadata,omitempty"`
	Items        []NotificationMessage `json:"items"`
}

// NotificationMessageSpec defines the desired state of NotificationMessage
type NotificationMessageSpec struct {
	// Notifications associated by this message
	Notifications []NotificationItem `json:"notifications,omitempty"`
	// Message body, json in any form
	Body runtime.RawExtension `json:"body"`
}

type NotificationItem struct {
	Name      string                `json:"name"`
	Ephemeral NotificationEphemeral `json:"ephemeral,omitempty"`
}

type NotificationEphemeral struct {
	// Users for this subscription
	// +optional
	Users []string `json:"users,omitempty"`
	// Groups for this subscription
	// +optional
	Groups []string `json:"groups,omitempty"`
	// Methods for this subscription
	Methods []string `json:"methods"`
	// Template to render notification message
	Template string `json:"template"`
}

// NotificationMessageStatus defines the observed state of NotificationMessage
type NotificationMessageStatus struct {
	Conditions []NotificationMessageCondition `json:"conditions,omitempty"`
	// +optional
	Phase Phase `json:"phase,omitempty"`
}

// NotificationMessageCondition contains condition information for a notification message.
type NotificationMessageCondition struct {
	// Uid for this condition
	Uid string `json:"uid,omitempty"`
	// Notification referred by this condition
	Notification string `json:"notification,omitempty"`
	// Method referred by this condition
	Method string `json:"method,omitempty"`
	// Users referred by this condition
	// +optional
	Users []string `json:"users,omitempty"`
	// Groups referred by this condition
	// +optional
	Groups []string `json:"groups,omitempty"`
	// Addresses for this condition
	Addresses []string `json:"addresses,omitempty"`
	// Type of notification message condition.
	Type ConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status"`
	// Retry times for this notification message.
	// +optional
	RetryTimes int `json:"retryTimes,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime mv1.Time `json:"lastTransitionTime,omitempty"`
	// Brief reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}
