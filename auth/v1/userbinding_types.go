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
	"context"
	"fmt"
	"strings"

	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/btcsuite/btcutil/base58"

	rc "gomod.alauda.cn/ait-apis/auth/constant"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ConditionType string

const (
	UserBindingBind ConditionType = "RoleBinding"
	UserBindingSync ConditionType = "UserbindingSync"
)

type UserBindingScope string

const (
	UserBindingScopePlatform  UserBindingScope = "platform"
	UserBindingScopeCluster   UserBindingScope = "cluster"
	UserBindingScopeProject   UserBindingScope = "project"
	UserBindingScopeNamespace UserBindingScope = "namespace"
)

type SubjectKind string

const (
	SubjectKindUser           = "User"
	SubjectKindGroup          = "Group"
	SubjectKindServiceAccount = "ServiceAccount"
)

type UserbindingsPhase string

const (
	UserbindingSuccess  UserbindingsPhase = "Success"
	UserbindingFailture UserbindingsPhase = "Failture"
	UserbindingUnknown  UserbindingsPhase = "Unknown"
)

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type Subject struct {
	Kind SubjectKind `json:"kind"`
	Name string      `json:"name"`
}

type Constraint struct {
	Project   string `json:"project,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Cluster   string `json:"cluster,omitempty"`
}

type UserBindingCondition struct {
	LastTransitionTime *metav1.Time    `json:"lastTransitionTime,omitempty"`
	Message            string          `json:"message,omitempty"`
	Reason             string          `json:"reason,omitempty"`
	Status             ConditionStatus `json:"status"`
	Type               ConditionType   `json:"type"`
}

// UserBindingSpec defines the desired state of UserBinding
type UserBindingSpec struct {
	Subjects   []Subject        `json:"subjects"`
	RoleRef    string           `json:"roleRef"`
	Scope      UserBindingScope `json:"scope"`
	Constraint []Constraint     `json:"constraint,omitempty"`
}

// UserBindingStatus defines the observed state of UserBinding
type UserBindingStatus struct {
	Conditions   []UserBindingCondition `json:"conditions,omitempty"`
	LastSpecHash string                 `json:"lastSpecHash,omitempty"`
	Phase        UserbindingsPhase      `json:"phase,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserBinding is the Schema for the userbindings API
// +k8s:openapi-gen=true
type UserBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserBindingSpec   `json:"spec,omitempty"`
	Status UserBindingStatus `json:"status,omitempty"`
}

func (u *UserBinding) RoleLevel() string {
	if level, ok := u.Labels[rc.LabelRoleLevel]; ok {
		return level
	}
	return ""
}

func (u *UserBinding) IsFederated() bool {
	enabled, ok := u.Labels[rc.LabelFederated]
	return ok && enabled == "true"
}

func (u *UserBinding) IsFederatedExsit() bool {
	_, ok := u.Labels[rc.LabelFederated]
	return ok
}

func (u *UserBinding) ShouldFederateSync() bool {
	return u.IsFederated() && u.RoleLevel() == rc.RoleLevelNamespace
}

func (u *UserBinding) RoleName() string {
	if name, ok := u.Labels[rc.LabelRoleName]; ok {
		return name
	}
	return ""
}

func (u *UserBinding) UserEmail() string {
	if u.IsFederatedExsit() {
		if email, ok := u.Labels[rc.LabelUserEmailBase58]; ok {
			return string(base58.Decode(email))
		}
	}
	if email, ok := u.Annotations[rc.AnnotationUserEmail]; ok {
		return email
	}
	return ""
}

func (u *UserBinding) NamespaceCluster() string {
	if cluster, ok := u.Labels[rc.LabelCluster]; ok {
		return cluster
	}
	return ""
}

func (u *UserBinding) ConstraintClusters() []string {
	clusters := []string{}
	for _, c := range u.Spec.Constraint {
		clusters = append(clusters, c.Cluster)
	}
	return clusters
}

func (u *UserBinding) SetAnnotationUserEmail(email string) {

	if len(u.GetAnnotations()) == 0 {
		u.Annotations = make(map[string]string, 0)
	}
	u.Annotations[rc.AnnotationUserEmail] = email
}

func (u *UserBinding) IsUserEmailExists() bool {
	if len(u.UserEmail()) == 0 {
		return false
	}
	return true
}

func (u *UserBinding) Email() string {
	if email, ok := u.Annotations[rc.AnnotationUserEmail]; ok {
		return email
	}
	return ""
}

func (u *UserBinding) UserEmailName() string {
	return u.Labels[rc.LabelUserEmail]
}

func (u *UserBinding) CurrentCluster() string {
	if cluster, ok := u.Annotations[rc.AnnotationCurrentCluster]; ok {
		return cluster
	}
	return ""
}

func (u *UserBinding) IsCurrentClusterExists() bool {
	if len(u.CurrentCluster()) == 0 {
		return false
	}
	return true
}

func (u *UserBinding) SetRoleName(roleName string) {
	if len(u.GetLabels()) == 0 {
		u.Labels = make(map[string]string, 0)
	}
	u.Labels[rc.LabelRoleName] = roleName
}

func (u *UserBinding) SetCurrentCluster(cluster string) {

	if len(u.GetAnnotations()) == 0 {
		u.Annotations = make(map[string]string, 0)
	}
	u.Annotations[rc.AnnotationCurrentCluster] = cluster
}

func (u *UserBinding) ProjectName() string {
	if project, ok := u.Labels[rc.LabelProject]; ok {
		return project
	}
	return ""
}

func (u *UserBinding) NamespaceName() string {
	if ns, ok := u.Labels[rc.LabelNamespace]; ok {
		return ns
	}
	return ""
}

func (u *UserBinding) Validate() error {
	if level, ok := u.Labels[rc.LabelRoleLevel]; level == "" || !ok {
		return fmt.Errorf("role level label is invalid")
	}

	if name, ok := u.Labels[rc.LabelRoleName]; name == "" || !ok {
		return fmt.Errorf("role name label is invalid")
	}

	email, _ := u.Labels[rc.LabelUserEmail]
	group, _ := u.Labels[rc.LabelGroupName]
	if email == "" && group == "" {
		return fmt.Errorf("subject user or group label is invalid")
	}

	switch u.RoleLevel() {
	case rc.RoleLevelPlatform:
		if u.ProjectName() != "" || u.NamespaceName() != "" || u.NamespaceCluster() != "" {
			return fmt.Errorf("platform level parameters are invalid")
		}
	case rc.RoleLevelCluster:
		if len(u.Spec.Constraint) == 0 {
			return fmt.Errorf("cluster level constraint are invalid")
		}
	case rc.RoleLevelProject:
		if u.ProjectName() == "" || u.NamespaceName() != "" || u.NamespaceCluster() != "" {
			return fmt.Errorf("project level parameters are invalid")
		}
	case rc.RoleLevelNamespace:
		if u.ProjectName() == "" || u.NamespaceName() == "" || u.NamespaceCluster() == "" {
			return fmt.Errorf("namespace level parameters are invalid")
		}
	default:
		return fmt.Errorf("role level is invalid")
	}
	return nil
}

func (u *UserBinding) Update(client client.Client) error {
	return client.Update(context.TODO(), u)
}

func (u *UserBinding) ObjForCreate() *UserBinding {
	return &UserBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:        u.Name,
			Labels:      u.Labels,
			Annotations: u.Annotations,
		},
		Spec: u.Spec,
	}
}

func (u *UserBinding) ParseUserFromEmail() string {
	arr := strings.Split(u.UserEmailName(), ".")
	return arr[0]
}

func (u *UserBinding) InstanceForBusinessCluster() UserBinding {
	newU := u.DeepCopy()
	annotations := newU.GetAnnotations()
	if len(annotations) == 0 {
		annotations = make(map[string]string, 0)
	}
	annotations[rc.AnnotationCurrentCluster] = rc.ClusterBusiness
	return UserBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:        newU.Name,
			Labels:      newU.GetLabels(),
			Annotations: annotations,
		},
		Spec: newU.Spec,
	}
}

func (u *UserBinding) IsNeedSync() bool {
	if currentCluster, ok := u.Annotations[rc.AnnotationCurrentCluster]; ok && currentCluster == rc.ClusterGlobal {
		return true
	}
	return false
}

func (u *UserBinding) Subjects() []rbacv1.Subject {
	subjects := []rbacv1.Subject{}
	for _, s := range u.Spec.Subjects {
		subjects = append(subjects, rbacv1.Subject{
			APIGroup: rbacv1.GroupName,
			Kind:     string(s.Kind),
			Name:     s.Name,
		})
	}
	if len(subjects) == 0 {
		subjects = append(subjects, rbacv1.Subject{
			APIGroup: rbacv1.GroupName,
			Kind:     rbacv1.UserKind,
			Name:     u.Email(),
		})
	}
	return subjects
}

func (u *UserBinding) GroupName() string {
	if group, ok := u.Labels[rc.LabelGroupName]; ok {
		return group
	}
	return ""
}

func (u *UserBinding) SubjectKind() string {
	if u.Spec.Subjects != nil && len(u.Spec.Subjects) > 0 {
		return string(u.Spec.Subjects[0].Kind)
	}
	if _, ok := u.Labels[rc.LabelUserEmail]; ok {
		return SubjectKindUser
	}
	if _, ok := u.Labels[rc.LabelGroupName]; ok {
		return SubjectKindGroup
	}
	return ""
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserBindingList contains a list of UserBinding
type UserBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserBinding{}, &UserBindingList{})
}
