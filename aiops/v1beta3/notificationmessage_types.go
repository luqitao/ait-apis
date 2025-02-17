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

package v1beta3

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
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
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationMessageSpec   `json:"spec,omitempty"`
	Status NotificationMessageStatus `json:"status,omitempty"`
}

// NotificationMessageList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NotificationMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []NotificationMessage `json:"items"`
}

// NotificationMessageSpec defines the desired state of NotificationMessage
// NotificationMessageSpec defines the desired state of NotificationMessage
type NotificationMessageSpec struct {
	// Notifications associated by this message
	Notifications []NotificationItem `json:"notifications"`
	// Message body, json in any form
	Body runtime.RawExtension `json:"body"`
}

type NotificationItem struct {
	Name      string                     `json:"name"`
	Method    string                     `json:"method,omitempty"`
	Sender    string                     `json:"sender,omitempty"`
	Receivers []NotificationReceiverItem `json:"receivers,omitempty"`
	Template  NotificationTemplateItem   `json:"template,omitempty"`
}

type NotificationReceiverItem struct {
	// Name of notification receiver
	Name string `json:"name"`
	// Namespace for notification receiver
	Namespace string `json:"namespace"`
	// Destination for notification receiver
	// +optional
	Destination string `json:"destination,omitempty"`
}

type NotificationTemplateItem struct {
	// Name for notification template
	Name string `json:"name"`
	// Content for notification template
	// +optional
	Content string `json:"content,omitempty"`
	// Subject for notification template
	// +optional
	Subject string `json:"subject,omitempty"`
}

var _ resource.Object = &NotificationMessage{}
var _ resourcestrategy.Validater = &NotificationMessage{}

func (in *NotificationMessage) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *NotificationMessage) NamespaceScoped() bool {
	return false
}

func (in *NotificationMessage) New() runtime.Object {
	return &NotificationMessage{}
}

func (in *NotificationMessage) NewList() runtime.Object {
	return &NotificationMessageList{}
}

func (in *NotificationMessage) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "aiops.alauda.io",
		Version:  "v1beta3",
		Resource: "notificationmessages",
	}
}

func (in *NotificationMessage) IsStorageVersion() bool {
	return true
}

func (in *NotificationMessage) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &NotificationMessageList{}

func (in *NotificationMessageList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// NotificationMessageStatus defines the observed state of NotificationMessage
type NotificationMessageStatus struct {
	Conditions []NotificationMessageCondition `json:"conditions,omitempty"`
	// +optional
	Phase Phase `json:"phase,omitempty"`
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

func (in *NotificationMessageStatus) SubResourceName() string {
	return ""
}

// NotificationMessage implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &NotificationMessage{}

func (in *NotificationMessage) GetStatus() resource.StatusSubResource {
	return &in.Status
}

// NotificationMessageStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &NotificationMessageStatus{}

func (in NotificationMessageStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*NotificationMessage).Status = in
}
