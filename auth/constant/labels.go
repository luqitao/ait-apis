package constant

var (
	LabelUser              = "%s/user"
	LabelUserName          = "auth.%s/user.name"
	LabelUserUsername      = "auth.%s/user.username"
	LabelUserEmail         = "auth.%s/user.email"
	LabelUserEmailBase58   = "auth.%s/user.email.base58"
	LabelUserValid         = "auth.%s/user.valid"
	LabelUserState         = "auth.%s/user.state"
	LabelUserConnectorType = "auth.%s/user.connector_type"
	LabelUserConnectorID   = "auth.%s/user.connector_id"
	LabelConnectorType     = "auth.%s/connector.type"
	LabelConnectorID       = "auth.%s/connector.id"

	LabelUserBindingName = "auth.%s/userbinding.name"

	LabelGroupDisplayName = "auth.%s/group.display-name"
	LabelGroupName        = "auth.%s/group.name"

	LabelRoleName      = "auth.%s/role.name"
	LabelRoleLevel     = "auth.%s/role.level"
	LabelRoleVisible   = "auth.%s/role.visible"
	LabelRoleOfficial  = "auth.%s/role.official"
	LabelRoleRelative  = "auth.%s/role.relative"
	LabelRoleBindScope = "auth.%s/role.bindscope"
	//LabelRoleBindNamespaceType  = "auth.%s/role.bindnamespacetype"
	//LabelRoleBindNamespaceValue = "auth.%s/role.bindnamespacevalue"
	LabelRolePart        = "auth.%s/role.part"
	LabelRoleBindCluster = "auth.%s/role.bindcluster"
	LabelFederated       = "auth.%s/federated"

	LabelRoleTemplateOfficial = "auth.%s/roletemplate.official"
	LabelRoleTemplateLevel    = "auth.%s/roletemplate.level"
	LabelRoleTemplateName     = "auth.%s/roletemplate.name" // for labelSelector query

	LabelCreatorEmail = "auth.%s/creator.email"

	LabelCluster     = "%s/cluster"
	LabelClusterName = "%s/cluster.name"
	LabelClusterType = "%s/cluster.type"

	LabelProject       = "%s/project"
	LabelProjectParent = "%s/project.parent"
	LabelProjectLevel  = "%s/project.level"
	LabelProjectId     = "%s/project.id"
	LabelNamespace     = "%s/namespace"

	LabelSystemRoleBinding        = "%s/system-rolebinding"
	LabelSystemClusterRoleBinding = "%s/system-clusterrolebinding"

	LabelFunctionResourceModule   = "auth.%s/functionresource.module"
	LabelFunctionResourceFunction = "auth.%s/functionresource.function"

	FederationName     = "federation.%s/name"
	FederationHostName = "federation.%s/hostname"

	// Project log policy related labels
	LabelIndicesKeepDays = "%s/project.esIndicesKeepDays"
	LabelEsPolicyEnabled = "%s/project.esPolicyEnabled"

	LabelDeprecated = "deprecated"

	LabelResourceType = "%s/resource.type"

	LabelRoleTemplateAggregate = "auth.%s/aggregate"
	LabelRolePartNamespace     = "auth.%s/role.part.namespace"
	LabelRolePartProject       = "auth.%s/role.part.project"

	// Labels about certificate manager
	LabelCertificateKind = "cert.%s/kind"

	// Labels about features
	LabelFeatureType         = "%s/type"
	LabelFeatureInstanceType = "%s/instance.type"
	LabelFeatureVMCluster    = "%s/victoriametrics.cluster"

	// Labels about vmrule
	LabelRuleCluster        = "rule.%s/cluster"
	LabelRuleNamespace      = "rule.%s/namespace"
	LabelRuleName           = "rule.%s/name"
	LabelRuleAlertCluster   = "alert.%s/cluster"
	LabelRuleAlertKind      = "alert.%s/kind"
	LabelRuleAlertName      = "alert.%s/name"
	LabelRuleAlertNamespace = "alert.%s/namespace"
	LabelRuleAlertOptions   = "alert.%s/options"
	LabelRuleVMOwnerRefer   = "%s/rule.owner.refer"

	// Labels about dex connector
	LabelLdapAutoSync = "%s/ldap.autoSync"
	LabelLdapName     = "%s/ldap.name"

	LabelProduct = "%s/product"
)
