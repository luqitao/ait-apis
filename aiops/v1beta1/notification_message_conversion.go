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
// +kubebuilder:docs-gen:collapse=Apache License

package v1beta1

/*
For imports, we'll need the controller-runtime
[`conversion`](https://godoc.org/sigs.k8s.io/controller-runtime/pkg/conversion)
package, plus the API version for our hub type (v1), and finally some of the
standard packages.
*/
import (
	av2 "gomod.alauda.cn/ait-apis/aiops/v1beta2"

	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:docs-gen:collapse=Imports

/*
Our "spoke" versions need to implement the
[`Convertible`](https://godoc.org/sigs.k8s.io/controller-runtime/pkg/conversion#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this NotificationMessage to the v1beta2 version.
func (in *NotificationMessage) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*av2.NotificationMessage)

	// ObjectMeta
	dst.ObjectMeta = in.ObjectMeta

	// Spec
	notifications := make([]av2.NotificationItem, len(in.Spec.Notifications))
	for i, v := range in.Spec.Notifications {
		notifications[i] = av2.NotificationItem{Name: v.Name}
	}
	dst.Spec.Notifications = notifications
	dst.Spec.Body = in.Spec.Body

	// Status
	conditions := make([]av2.NotificationMessageCondition, len(in.Status.Conditions))
	for i, v := range in.Status.Conditions {
		conditions[i] = av2.NotificationMessageCondition{
			Notification:       v.Notification,
			Message:            v.Message,
			ReceiverName:       v.ReceiverName,
			ReceiverNamespace:  v.ReceiverNamespace,
			Type:               av2.ConditionType(v.Type),
			Status:             av2.ConditionStatus(v.Status),
			RetryTimes:         v.RetryTimes,
			LastTransitionTime: v.LastTransitionTime,
			Reason:             v.Reason,
		}
	}
	phase := av2.NotificationMessagePhase(in.Status.Phase)
	if string(phase) == "" {
		phase = av2.NotificationMessagePending
	}
	dst.Status = av2.NotificationMessageStatus{Conditions: conditions, Phase: phase}
	// +kubebuilder:docs-gen:collapse=rote conversion
	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the v1beta2 to this version.
func (in *NotificationMessage) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*av2.NotificationMessage)

	// ObjectMeta
	in.ObjectMeta = src.ObjectMeta

	// Spec
	notifications := make([]NotificationName, len(src.Spec.Notifications))
	for i, v := range src.Spec.Notifications {
		notifications[i] = NotificationName{Name: v.Name}
	}
	in.Spec.Notifications = notifications
	in.Spec.Body = src.Spec.Body

	// Status
	conditions := make([]NotificationMessageCondition, len(in.Status.Conditions))
	for i, v := range src.Status.Conditions {
		conditions[i] = NotificationMessageCondition{
			Notification:       v.Notification,
			Message:            v.Message,
			ReceiverName:       v.ReceiverName,
			ReceiverNamespace:  v.ReceiverNamespace,
			Type:               ConditionType(v.Type),
			Status:             ConditionStatus(v.Status),
			RetryTimes:         v.RetryTimes,
			LastTransitionTime: v.LastTransitionTime,
			Reason:             v.Reason,
		}
	}
	phase := NotificationMessagePhase(src.Status.Phase)
	if string(phase) == "" {
		phase = NotificationMessagePending
	}
	in.Status = NotificationMessageStatus{Conditions: conditions, Phase: phase}

	// +kubebuilder:docs-gen:collapse=rote conversion
	return nil
}
