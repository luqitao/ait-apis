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
	mv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotificationGroupSpec defines the desired state of NotificationGroup
type NotificationGroupSpec struct {
	// Users for notification group
	// +optional
	Users []string `json:"users,omitempty"`
	// Mail for notification group
	// +optional
	Mail string `json:"mail,omitempty"`
	// WebhookUrl for notification group
	// +optional
	WebhookUrl string `json:"webhookUrl,omitempty"`
	// WebhookType for notification group
	// +optional
	WebhookType string `json:"webhookType,omitempty"`
}

// NoticationGroup is the Schema for the templates API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=nfg
// +kubebuilder:storageversion
type NotificationGroup struct {
	mv1.TypeMeta   `json:",inline"`
	mv1.ObjectMeta `json:"metadata,omitempty"`
	Spec           NotificationGroupSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotificationGroupList contains a list of NotificationGroup
type NotificationGroupList struct {
	mv1.TypeMeta `json:",inline"`
	mv1.ListMeta `json:"metadata,omitempty"`
	Items        []NotificationGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationGroup{}, &NotificationGroupList{})
}
