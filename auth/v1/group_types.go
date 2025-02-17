/*
Copyright 2021.

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

package v1

import (
	rc "gomod.alauda.cn/ait-apis/auth/constant"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type RefConnector struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type GroupMember struct {
	Name       string `json:"name,omitempty"`
	ImportTime string `json:"importTime,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	// Connector information related to user groups
	RefConnector *RefConnector `json:"refConnector,omitempty"`
	// List of group members. The list element is user resource name.
	Users []GroupMember `json:"users,omitempty"`
}

// GroupStatus defines the observed state of Group
type GroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:scope=Cluster,singular="group"
//+kubebuilder:subresource:status

// Group is the Schema for the groups API
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GroupSpec   `json:"spec,omitempty"`
	Status GroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GroupList contains a list of Group
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}

func (g *Group) DisplayName() string {
	annos := g.GetAnnotations()
	displayName, ok := annos[rc.AnnotationDisplayName]
	if !ok {
		return ""
	}
	return displayName
}
