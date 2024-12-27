package v1beta1

import (
	"fmt"
)

const (
	// Available notification methods
	MethodWebhook  = "Webhook"
	MethodEmail    = "Email"
	MethodSMS      = "Sms"
	MethodDingtalk = "DingTalk"
	MethodWechat   = "WeChat"

	// Types for secret
	NotificationSender = "NotificationSender"
	NotificationServer = "NotificationServer"

	// Default notification settings
	DefaultSMSTemplateID = "462079"
	Done                 = "Done"
	MigrateCMName        = "aiops-nf-migrate-config"
	EmailServerName      = "platform-email-server"

	// Label values for notification
	LabelValueNotification         = "Notification"
	LabelValueNotificationTemplate = "NotificationTemplate"
	LabelValueTemplateMessageAlert = "Alert"

	// Message types for wechat and dingtalk
	MessageTypeText     = "text"
	MessageTypeMarkdown = "markdown"

	// Rule label keys
	RuleLabelKeyAlertResource        = "alert_resource"
	RuleLabelKeyAlertMessage         = "alert_message"
	RuleLabelKeyAlertName            = "alert_name"
	RuleLabelKeyAlertObjectName      = "alert_involved_object_name"
	RuleLabelKeyAlertObjectKind      = "alert_involved_object_kind"
	RuleLabelKeyAlertObjectNamespace = "alert_involved_object_namespace"
	RuleLabelKeyAlertObjectOptions   = "alert_involved_object_options"
	RuleLabelKeyAlertCreator         = "alert_creator"
	RuleLabelKeyAlertProject         = "alert_project"
	RuleLabelKeyAlertCluster         = "alert_cluster"
	RuleLabelKeyAlertIndicator       = "alert_indicator"
	RuleLabelKeyAlertIndicatorQuery  = "alert_indicator_query"
	RuleLabelKeyAlertCompare         = "alert_compare"
	RuleLabelKeyAlertThreshold       = "alert_threshold"
	RuleLabelKeyAlertUnit            = "alert_unit"
	RuleLabelKeyAlertKind            = "alert_kind"
	RuleLabelKeyAlertSource          = "alert_source"
	RuleLabelKeySeverity             = "severity"
	RuleLabelKeyApplication          = "application"
)

var (
	// Annotations keys
	AnnotationKeyDescription = "%s/description"
	AnnotationKeyDisplayName = "%s/display-name"

	// Label keys for notification
	LabelKeyBackup                   = "aiops.%s/backup"
	LabelKeyGroup                    = "%s/notification.group."
	LabelKeyKind                     = "aiops.%s/kind"
	LabelKeyMigrate                  = "aiops.%s/migrate"
	LabelKeyNotification             = "alert.%s/notifications"
	LabelKeyProject                  = "%s/project"
	LabelKeyServer                   = "%s/server"
	LabelKeyServerType               = "%s/notification.server.type"
	LabelKeyServerSkipVerify         = "%s/notification.server.skip.verify"
	LabelKeySource                   = "%s/source"
	LabelKeyTemplate                 = "%s/notification.template."
	LabelKeyTemplateMessageType      = "%s/template.message.type"
	LabelKeyTemplateEmailBodyType    = "%s/template.email.body.type"
	LabelKeyTemplateEmailBodyTypeOld = "%s/message-type"
	LabelKeyTemplateWebhookBodyType  = "%s/template.webhook.body.type"
	LabelKeyType                     = "%s/type"
)

func InitWithDomain(domain string) {
	// Initialize common annotations keys
	AnnotationKeyDescription = fmt.Sprintf(AnnotationKeyDescription, domain)
	AnnotationKeyDisplayName = fmt.Sprintf(AnnotationKeyDisplayName, domain)

	// Initialize label keys for notification
	LabelKeyBackup = fmt.Sprintf(LabelKeyBackup, domain)
	LabelKeyGroup = fmt.Sprintf(LabelKeyGroup, domain)
	LabelKeyKind = fmt.Sprintf(LabelKeyKind, domain)
	LabelKeyMigrate = fmt.Sprintf(LabelKeyMigrate, domain)
	LabelKeyNotification = fmt.Sprintf(LabelKeyNotification, domain)
	LabelKeyProject = fmt.Sprintf(LabelKeyProject, domain)
	LabelKeyServer = fmt.Sprintf(LabelKeyServer, domain)
	LabelKeyServerType = fmt.Sprintf(LabelKeyServerType, domain)
	LabelKeyServerSkipVerify = fmt.Sprintf(LabelKeyServerSkipVerify, domain)
	LabelKeySource = fmt.Sprintf(LabelKeySource, domain)
	LabelKeyTemplate = fmt.Sprintf(LabelKeyTemplate, domain)
	LabelKeyTemplateMessageType = fmt.Sprintf(LabelKeyTemplateMessageType, domain)
	LabelKeyTemplateEmailBodyType = fmt.Sprintf(LabelKeyTemplateEmailBodyType, domain)
	LabelKeyTemplateEmailBodyTypeOld = fmt.Sprintf(LabelKeyTemplateEmailBodyTypeOld, domain)
	LabelKeyTemplateWebhookBodyType = fmt.Sprintf(LabelKeyTemplateWebhookBodyType, domain)
	LabelKeyType = fmt.Sprintf(LabelKeyType, domain)
}
