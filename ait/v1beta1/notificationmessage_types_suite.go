package v1beta1

import (
	"context"

	mv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

var _ resource.Object = &NotificationMessage{}

func (in *NotificationMessage) GetObjectMeta() *mv1.ObjectMeta {
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
		Group:    "ait.alauda.io",
		Version:  "v1beta1",
		Resource: "notificationmessages",
	}
}

func (in *NotificationMessage) IsStorageVersion() bool {
	return true
}

var _ resourcestrategy.Validater = &NotificationMessage{}

func (in *NotificationMessage) Validate(_ context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

// NotificationMessage implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &NotificationMessage{}

func (in *NotificationMessage) GetStatus() resource.StatusSubResource {
	return &in.Status
}

// NotificationMessageList implements ObjectList interface.
var _ resource.ObjectList = &NotificationMessageList{}

func (in *NotificationMessageList) GetListMeta() *mv1.ListMeta {
	return &in.ListMeta
}

// NotificationMessageStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &NotificationMessageStatus{}

func (in NotificationMessageStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*NotificationMessage).Status = in
}

func (in *NotificationMessageStatus) SubResourceName() string {
	return ""
}
