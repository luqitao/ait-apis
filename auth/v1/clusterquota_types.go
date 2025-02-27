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

// ClusterQuotaSpec defines the desired state of ClusterQuota
type ClusterQuotaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// ClusterQuotaStatus defines the observed state of ClusterQuota
type ClusterQuotaStatus struct {
	// Capacity represents the total resources of a cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// +optional
	Capacity corev1.ResourceList `json:"capacity,omitempty"`
	// Allocatable represents the resources of a cluster that are available for scheduling.
	// Defaults to Capacity.
	// +optional
	Allocatable corev1.ResourceList `json:"allocatable,omitempty"`
	// Used represents the resources of a cluster that are consumed when creating project.
	// Defaults to Zere.
	// +optional
	Used corev1.ResourceList `json:"used,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterQuota is the Schema for the clusterquota API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ClusterQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterQuotaSpec   `json:"spec,omitempty"`
	Status ClusterQuotaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterQuotaList contains a list of ClusterQuota
type ClusterQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterQuota{}, &ClusterQuotaList{})
}
