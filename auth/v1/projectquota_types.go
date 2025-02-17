/*
Copyright 2018 Alauda.io.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectQuota is the Schema for the projectquota API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ProjectQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   corev1.ResourceQuotaSpec   `json:"spec,omitempty"`
	Status corev1.ResourceQuotaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectQuotaList contains a list of ProjectQuota
type ProjectQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectQuota{}, &ProjectQuotaList{})
}

// InitStatus build status from spec
func (pq *ProjectQuota) InitStatus() {
	// construct the default Hard and Used in Status
	emptyQuota := pq.Spec.Hard.DeepCopy()
	for key, quantity := range pq.Spec.Hard {
		quantity.Sub(pq.Spec.Hard[key])
		emptyQuota[key] = quantity
	}
	pq.Status.Hard = emptyQuota
	pq.Status.Used = emptyQuota
}

// SetDefaults make sure that Labels are initialized
func (object *ProjectQuota) SetDefaults() {
	if len(object.Labels) == 0 {
		object.Labels = map[string]string{}
	}
}
