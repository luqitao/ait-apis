package constant

var (
	AnnotationForceDelete                           = "%s/force-delete"
	AnnotationDisplayName                           = "%s/display-name"
	AnnotationDisplayNameEn                         = "%s/displayNameEn"
	AnnotationFunctionResourceModuleDisplayName     = "%s/functionresource.module.display-name"
	AnnotationFunctionResourceModuleDisplayNameEn   = "%s/functionresource.module.display-name.en"
	AnnotationFunctionResourceFunctionDisplayName   = "%s/functionresource.function.display-name"
	AnnotationFunctionResourceFunctionDisplayNameEn = "%s/functionresource.function.display-name.en"
	AnnotationProduct                               = "%s/product"
	AnnotationProductVersion                        = "%s/productVersion"
	AnnotationRoleVersion                           = "%s/roleVersion"
	AnnotationProject                               = "%s/project"
	AnnotationProjectLevel                          = "%s/project.level"
	AnnotationProjectParent                         = "%s/project.parent"
	AnnotationCluster                               = "%s/cluster"
	AnnotationClusterName                           = "%s/cluster.name"
	AnnotationClusterType                           = "%s/cluster.type"
	AnnotationClusterAttr                           = "legacy.cluster.%s/attr"
	AnnotationPipelineLastNumber                    = "%s/pipeline.last.number"
	AnnotationPipelineNumber                        = "%s/pipeline.number"
	AnnotationPipelineConfig                        = "%s/pipelineConfig.name"
	AnnotationJenkinsBuildURI                       = "%s/jenkins-build-uri"
	AnnotationSecretType                            = "%s/secretType"
	AnnotationCreateAppUrl                          = "%s/createAppUrl"
	AnnotationUpdatedAt                             = "%s/updated-at"
	AnnotationDescription                           = "%s/description"
	AnnotationCreator                               = "%s/creator"
	AnnotationCurrentCluster                        = "%s/current-cluster"
	AnnotationUserEmail                             = "auth.%s/user.email"
	AnnotationRoleDisplayName                       = "auth.%s/role.display-name"
	AnnotationRetryTimes                            = "%s/retryTimes"
	AnnotationSilenceConfig                         = "alert.%s/silence.config"
	AnnotationRBACStatus                            = "%s/rbac-status"
	AnnotationGroupName                             = "auth.%s/group.name"

	// Project log policy related annotations
	AnnotationLastEnabledTime  = "%s/project.esPolicyLastEnabledTimestamp"
	AnnotationLastDisabledTime = "%s/project.esPolicyLastDisabledTimestamp"

	// Annotations about certificate manager
	AnnotationCertificateNode      = "cert.%s/node"
	AnnotationCertificateNamespace = "cert.%s/namespace"
	AnnotationCertificateName      = "cert.%s/name"
	AnnotationCertificateNotBefore = "cert.%s/notBefore"
	AnnotationCertificateNotAfter  = "cert.%s/notAfter"

	AnnotationUniteQuota = "%s/unite-quota-clusters"

	// Annotations about dex connector
	AnnotationLdapAutoSyncRule  = "%s/ldap.autoSyncRule"
	AnnotationLdapType          = "%s/ldap.type"
	AnnotationLastSyncTimestamp = "%s/ldap.lastSyncTimestamp"
	AnnotationLastSyncStatus    = "%s/ldap.lastSyncStatus"
	AnnotationLastSyncMessage   = "%s/ldap.lastSyncMessage"
	AnnotationIdpValidation     = "%s/idp.validation"

	AnnotationConnectorData = "auth.%s/connector.data"
)
