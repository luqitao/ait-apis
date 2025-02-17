package v1beta1

import (
	"os"
	"strings"
)

var ControllerNamespace string

const (
	// Available notification methods
	MethodWebhook  = "webhook"
	MethodEmail    = "email"
	MethodSMS      = "sms"
	MethodDingtalk = "dingtalk"
	MethodWechat   = "wechat"

	// Type for notification sender secret
	SecretTypeNotificationSender = "NotificationSender"

	// Default notification settings
	DefaultEmailSender   = "default-email-sender"
	DefaultSMSSender     = "default-sms-sender"
	DefaultTemplate      = "default-template"
	DefaultSMSTemplateID = "462079"

	// Message types
	MessageTypeText     = "text"
	MessageTypeMarkdown = "markdown"
)

var (
	// Common annotations keys
	AnnotationKeyDisplayName           = "display-name"
	AnnotationKeySynchronizationResult = "synchronization-result"
	AnnotationKeyCreator               = "creator"

	// Common label keys
	LabelKeyOwner             = "owner"
	LabelKeyServer            = "server"
	LabelKeyType              = "type"
	LabelKeyVersion           = "version"
	LabelKeyMessageType       = "message-type"
	LabelKeyUniqueName        = "unique-name"
	LabelKeyDestinationMD5    = "destination-md5"
	LabelKeyDisplayNameMD5    = "display-name-md5"
	LabelNotificationTemplate = "notificationtemplate/"
)

func InitWithDomain(domain string) {
	AnnotationKeyDisplayName = strings.Join([]string{domain, AnnotationKeyDisplayName}, "/")
	AnnotationKeySynchronizationResult = strings.Join([]string{domain, AnnotationKeySynchronizationResult}, "/")
	AnnotationKeyCreator = strings.Join([]string{domain, AnnotationKeyCreator}, "/")
	LabelKeyOwner = strings.Join([]string{domain, LabelKeyOwner}, "/")
	LabelKeyServer = strings.Join([]string{domain, LabelKeyServer}, "/")
	LabelKeyType = strings.Join([]string{domain, LabelKeyType}, "/")
	LabelKeyVersion = strings.Join([]string{domain, LabelKeyVersion}, "/")
	LabelKeyMessageType = strings.Join([]string{domain, LabelKeyMessageType}, "/")
	LabelKeyUniqueName = strings.Join([]string{domain, LabelKeyUniqueName}, "/")
	LabelKeyDestinationMD5 = strings.Join([]string{domain, LabelKeyDestinationMD5}, "/")
	LabelKeyDisplayNameMD5 = strings.Join([]string{domain, LabelKeyDisplayNameMD5}, "/")
	LabelNotificationTemplate = strings.Replace(LabelNotificationTemplate, "/", "."+domain+"/", -1)

	ControllerNamespace = os.Getenv("NAMESPACE")
	if ControllerNamespace == "" {
		ControllerNamespace = "cpaas-system"
	}
}
