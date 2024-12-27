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

// NotificationTemplateSpec defines the desired state of NotificationTemplate
type NotificationTemplateSpec struct {
	Email   NotificationTemplateEmail  `json:"email"`
	SMS     NotificationTemplateCommon `json:"sms"`
	Webhook NotificationTemplateCommon `json:"webhook"`
}

// NotificationTemplateEmail defines the desired state of email NotificationTemplate
type NotificationTemplateEmail struct {
	// Subject template to render notification message
	// +optional
	Subject string `json:"subject"`
	// Content template to render notification message
	Content string `json:"content"`
}

// NotificationTemplateCommon defines the desired state of common NotificationTemplate
type NotificationTemplateCommon struct {
	// Content template to render notification message
	Content string `json:"content"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotificationTemplate is the Schema for the templates API
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName=nft,scope=Namespaced
// +kubebuilder:storageversion
type NotificationTemplate struct {
	mv1.TypeMeta   `json:",inline"`
	mv1.ObjectMeta `json:"metadata,omitempty"`
	Spec           NotificationTemplateSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NotificationTemplateList contains a list of NotificationTemplate
type NotificationTemplateList struct {
	mv1.TypeMeta `json:",inline"`
	mv1.ListMeta `json:"metadata,omitempty"`
	Items        []NotificationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationTemplate{}, &NotificationTemplateList{})
}
