package constant

import "time"

const (
	// role bind namespace type
	//RoleBindNamespaceTypeSystem   = "system"
	//RoleBindNamespaceTypeProject  = "project"
	//RoleBindNamespaceTypeCustomer = "customer"
	//RoleBindNamespaceTypeSpecific = "specific"
	RoleLabelPartSystem    = "system"
	RoleLabelPartCommon    = "common"
	RoleLabelPartProjectNs = "project_ns"

	//
	GroupName       = "auth.alauda.io"
	ClusterGlobal   = "global"
	ClusterBusiness = "business"

	// role bind cluster
	RoleBindClusterGlobal   = ClusterGlobal
	RoleBindClusterBusiness = ClusterBusiness
	RoleBindClusterUnlimit  = "unlimit"

	// role level
	RoleLevelPlatform  = "platform"
	RoleLevelCluster   = "cluster"
	RoleLevelProject   = "project"
	RoleLevelNamespace = "namespace"

	// role bind scope
	RoleBindScopeCluster   = "cluster"   // ClusterRoleBinding
	RoleBindScopeNamespace = "namespace" // RoleBinding

	ConstraintScopeClutser = "cluster"
	ConstraintScopeProject = "project"

	// auth.alauda.io
	RoleTemplateLabelOfficial      = "auth.alauda.io/roletemplate.official"
	RoleTemplateLabelLevel         = "auth.alauda.io/roletemplate.level"
	RoleTemplateLabelName          = "auth.alauda.io/roletemplate.name" // for labelSelector query
	RoleTemplateLabelEmail         = "auth.alauda.io/creator.email"
	UserEmailAnnotationKey         = "auth.alauda.io/user.email"
	ClusterRoleLabelRelative       = "auth.alauda.io/role.relative"
	RoleLevelLabelKey              = "auth.alauda.io/role.level"
	RoleOfficialKey                = "auth.alauda.io/role.official"
	RoleNameLabelKey               = "auth.alauda.io/role.name"
	RoleBindNamespaceTypeLabelKey  = "auth.alauda.io/role.bindnamespacetype"
	RoleBindNamespaceValueLabelKey = "auth.alauda.io/role.bindnamespacevalue"
	RoleBindScopeLabelKey          = "auth.alauda.io/role.bindscope"
	RoleBindClusterLabelKey        = "auth.alauda.io/role.bindcluster"
	UserEmailLabelKey              = UserEmailAnnotationKey

	// alauda.io
	ClusterRoleAnnotationsCurrentCluster = "alauda.io/current-cluster"
	ClusterLabelKey                      = "alauda.io/cluster"
	ProjectLabelKey                      = "alauda.io/project"
	NamespaceLabelKey                    = "alauda.io/namespace"

	// ClusterRole命名规则
	ClusterRoleNameSchemaSpecificBusiness = "%s-%s-b-%s" // 业务集群资源角色
	ClusterRoleNameSchemaCustomer         = "%s-%s"
	ClusterRoleNameSchemaProject          = "%s-pns-%s"        // 项目同名资源角色
	ClusterRoleNameSchemaSystem           = "%s-sys-%s"        // global集群系统资源角色
	ClusterRoleNameSchemaSystemBusiness   = "%s-sys-b-%s"      // 业务集群系统资源角色
	ClusterRoleNameSchemaCluster          = "ext-%s-%s"        // 项目/命名空间的集群资源扩展角色
	ClusterRoleNameSchemaConstraint       = "%s-%s-constraint" // 资源约束实例角色(模板),
	ClusterRoleNameSchemaRegion           = "%s-%s-region"     // 业务集群资源角色

	// query selector
	ClusterRoleLabelRelativeSelector = "auth.alauda.io/role.relative,auth.alauda.io/role.relative=%s"

	// ProjectId is a non-negative random number in [1, 0.2billion]
	ProjectIDRange      = 200000000
	ProjectIDRetryLimit = 5

	// common Context Timeout
	DefaultRequestTimeout = 120 * time.Second

	// Time format
	TimeFormat = "2006-01-02T15:04:05Z"
)
