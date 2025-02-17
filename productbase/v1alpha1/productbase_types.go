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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ChartVersion stands for chart version
type ChartVersion struct {
	// Name stands for the chart's name
	Name string `json:"name"`
	// Version stands for the chart's version
	Version string `json:"version"`
}

type SecretRef struct {
	// Name is the name of secret
	Name string `json:"name"`
}

// RegistryConfig ...
type RegistryConfig struct {
	// Address is Docker Registry's address
	Address string `json:"address"`
	// Username is Docker Registry's login user name
	// +optional
	Username string `json:"username"`
	// Password is Docker Registry's login password
	// +optional
	Password []byte `json:"password,omitempty"`
	// SecretRef is the reference for  pullImageSecret
	// +optional
	SecretRef *SecretRef `json:"secretRef,omitempty"`

	// External stands for a external registry
	// +optional
	External *bool `json:"external,omitempty"`

	// PreferPlatformURL stands for business cluster prefer platform URL as registry address
	PreferPlatformURL bool `json:"preferPlatformURL"`
}

// ChartMuseumConfig ...
type ChartMuseumConfig struct {
	// URL is Chart Museum's url
	URL string `json:"url"`
}

// ValuesOverride ...
type ValuesOverride map[string]runtime.RawExtension

// TLSConfig ...
type TLSConfig struct {
	// CA is the CA certificate content
	// +optional
	CA []byte `json:"ca,omitempty"`
	// Cert is the server certificate content
	Cert []byte `json:"cert"`
	// Key is the server's key
	Key []byte `json:"key"`
}

// HTTPConfig ...
type HTTPConfig struct {
	// HTTPPort is the port of http
	// +optional
	HTTPPort int `json:"httpPort"`
	// HTTPPort is the port of https
	// +optional
	HTTPSPort int `json:"httpsPort"`
	// ForceRedirectHTTPS stands for if force redirect to https
	ForceRedirectHTTPS bool `json:"forceRedirectHttps"`
	// TLS statnds for the server's TLS Certificate
	// +optional
	TLS *TLSConfig `json:"tls,omitempty"`
}

// UserConfig ...
type UserConfig struct {
	// Email stands for admin's email
	Email string `json:"email"`
	// Account stands for admin's account
	Account string `json:"account"`
	// Password stands for admin's password
	Password []byte `json:"password"`
}

// LogCenterConfig ...
type LogCenterConfig struct {
	// Enable stands for if log center enabled
	Enable bool `json:"enable"`
}

// LogAgentConfig ...
type LogAgentConfig struct {
	// Enable stands for if global cluster's log agent enabled
	Enable bool `json:"enable"`
}

// StorageType ...
type StorageType string

const (
	// LocalVolumeStorage ...
	LocalVolumeStorage StorageType = "LocalVolume"
	// StorageClassStorage ...
	StorageClassStorage StorageType = "StorageClass"
	// PVStorage ...
	PVStorage StorageType = "PV"
)

// StroageConfig ...
type StroageConfig struct {
	// Type stands for Prometheus storage type
	Type StorageType `json:"type"`
	// Path stands for Prometheus storage path
	// +optional
	Path string `json:"path,omitempty"`
	// Nodes stands for Prometheus deploy nodes
	// +optional
	Nodes []string `json:"nodes,omitempty"`
	// StorageClass stands for Prometheus's storage class
	// +optional
	StorageClass string `json:"storageClass,omitempty"`
	// Selector stands for Prometheus's PV selector
	// +optional
	Selector map[string]string `json:"selector,omitempty"`
	// Capacity stands for Prometheus storage capacity
	// +optional
	Capacity int `json:"capacity,omitempty"`
}

// PrometheusConfig ...
type PrometheusConfig struct {
	// Install stands for if install prometheus for global cluster
	Install bool `json:"install"`
	// Replicas stands for prometheus's replicas
	// +optional
	Replicas int `json:"replicas,omitempty"`
	// Storage stands for prometheus's storage config
	// +optional
	Storage *StroageConfig `json:"storage,omitempty"`
}

// ESConfig ...
type ESConfig struct {
	// Install stands for if install elasticsearch
	Install bool `json:"install"`
	// ExternalAddress statnds for external elasticsearch's address if not empty
	// +optional
	ExternalAddress *string `json:"extenalAddress,omitempty"`
	// Username stands for elasticsearch's username
	// +optional
	Username string `json:"username,omitempty"`
	// Username stands for elasticsearch's password
	// +optional
	Password []byte `json:"password,omitempty"`
	// Nodes stands for elasticsearch's deploy node
	// +optional
	Nodes []string `json:"nodes,omitempty"`
}

// KafkaConfig ...
type KafkaConfig struct {
	// Install stands for if install kafka
	Install bool `json:"install"`
	// Auth stands for if enable authentication
	// +optinal
	Auth *bool `json:"auth,omitempty"`
	// ExternalAddress stands for external kafka's address if not empty
	// +optional
	ExternalAddress *string `json:"extenalAddress,omitempty"`
	// Username stands for kafka's username
	// +optional
	Username string `json:"username,omitempty"`
	// Password stands for password's username
	// +optional
	Password []byte `json:"password,omitempty"`
	// Nodes stands for kafka's deploy node
	// +optional
	Nodes []string `json:"nodes,omitempty"`
}

// ContainerPlatformConfig ...
type ContainerPlatformConfig struct {
	// Install stands for if install ContainerPlatform
	Install bool `json:"install"`
}

// DevopsConfig ...
type DevopsConfig struct {
	// Install stands for if install DevopsConfig
	Install bool `json:"install"`
}

// ServiceMeshConfig ...
type ServiceMeshConfig struct {
	// Install stands for if install ServiceMesh
	Install bool `json:"install"`
}

// DataServicesConfig ...
type DataServicesConfig struct {
	// Install stands for if install DataServices
	Install bool `json:"install"`
}

// DeployType is platform deploy type
type DeployType string

const (
	// NormalDeploy stands for the normal deployment.
	NormalDeploy DeployType = "Normal"
	// BaseDeploy stands for the base deployment.
	BaseDeploy DeployType = "Base"
	// ImportedGlobalDeploy stands for deploy on an imported global cluster.
	ImportedGlobalDeploy DeployType = "ImportedGlobal"
)

type IngressController struct {
	// Install stands for if install ingress controller
	Install bool `json:"install"`
	// Nodes stands for the ingress controller's running node
	// +optional
	Nodes []string `json:"nodes,omitempty"`
}

type IngressConfig struct {
	// Host stands for the host config in ingress
	Host string `json:"host"`
	// IngressClassName stands for the ingressClassName of the ingress
	IngressClassName string `json:"ingressClassName"`
	// Controller stands for the ingress controller configs
	Controller  IngressController `json:"controller"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ETCDConfig struct {
	// HostPath stands for the host path on node to storage etcd dat
	// +optional
	HostPath string `json:"hostPath"`
	// PVCNames stands for the etcd's pvc names
	// +optional
	PVCNames []string `json:"pvcNames,omitempty"`
	// Nodes stands for the etcd's running node
	// +optional
	Nodes []string `json:"nodes,omitempty"`
}

// ConnectType ...
// +kubebuilder:validation:Enum=SSH-Tunnel;Public
type ConnectType string

const (
	// SSHTunnel stands for the connect type is SSH-Tunnel
	SSHTunnel ConnectType = "SSH-Tunnel"
	// Public stands for the connect type is Public
	Public ConnectType = "Public"
)

// CloudConfig ...
type CloudConfig struct {
	// Connect stands for if connect to cloud
	Connect bool `json:"connect"`
	// ConnectType stands for the connect type
	ConnectType ConnectType `json:"connectType"`
}

// ConnectPhase ...
// +kubebuilder:validation:Enum=ONLine;ConnectFailed;Connecting;OFFLine
type ConnectPhase string

const (
	ONLine        ConnectPhase = "ONLine"
	ConnectFailed ConnectPhase = "ConnectFailed"
	Connecting    ConnectPhase = "Connecting"
	OFFLine       ConnectPhase = "OFFLine"
)

// CloudStatus ...
type CloudStatus struct {
	// Phase stands for the connect phase
	Phase ConnectPhase `json:"phase"`
	// Message stands for the connect message
	Message string `json:"message"`
	// ConnectURL stands for the connect url
	ConnectURL string `json:"connectURL"`
}

// ProductBaseSpec defines the desired state of ProductBase
type ProductBaseSpec struct {

	// Version stands for this product's version to deploy
	Version string `json:"version"`
	// ChartVersions stands for the charts when deploying, they are optional and will be added to the Prouct CR
	// +optional
	ChartVersions []ChartVersion `json:"chartVersions,omitempty"`
	// ValuesOverride is the special values
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	ValuesOverride ValuesOverride `json:"valuesOverride,omitempty"`
	// Registry is Docker Registry's Config
	Registry RegistryConfig `json:"registry"`
	// ChartMuseum is chart meseum's config
	ChartMuseum ChartMuseumConfig `json:"chartMuseum"`
	// PlatformURL is the platform's access url
	PlatformURL string `json:"platformURL"`
	// AlternativeURLs is the platform's alternative urls
	AlternativeURLs []string `json:"alternativeURLs,omitempty"`
	// ACEEnable if enable ACE
	// +optional
	ACEEnable bool `json:"aceEnable,omitempty"`
	// Replicas stands for the deploy replicas
	// +optional
	Replicas int `json:"replicas"`
	// DefaultRedirectPath stands for the default redirect path
	// +optional
	DefaultRedirectPath string `json:"defaultRedirectPath,omitempty"`
	// HTTP stands for the http config
	HTTP HTTPConfig `json:"http"`
	// DefaultAdmin stands for the admin's config
	DefaultAdmin UserConfig `json:"defaultAdmin"`
	// LogCenter stands for the log center's config
	LogCenter LogCenterConfig `json:"logCenter"`
	// LogAgent stands for the log agent's config
	LogAgent LogAgentConfig `json:"logAgent"`
	// Prometheus stands for the prometheus's config
	Prometheus PrometheusConfig `json:"prometheus"`
	// Elasticsearch stands for the elasticsearch's config
	Elasticsearch ESConfig `json:"elasticsearch"`
	// Kafka statnds for the kafka's config
	Kafka KafkaConfig `json:"kafka"`
	// ContainerPlatform stands for ContainerPlatform's config
	ContainerPlatform ContainerPlatformConfig `json:"containerPlatform"`
	// Devops stands for Devops's config
	Devops DevopsConfig `json:"devops"`
	// ServiceMesh stands for ServiceMesh's config
	ServiceMesh ServiceMeshConfig `json:"serviceMesh"`
	// DataServices stands for DataServices's config
	DataServices *DataServicesConfig `json:"dataServices,omitempty"`

	// DeployType stands for deploy type
	// +optional
	DeployType DeployType `json:"deployType,omitempty"`
	// UseCustomerDefinedPorts stands for if use customer defined ports
	// +optional
	UseCustomerDefinedPorts bool `json:"useCustomerDefinedPorts,omitempty"`

	// Ingress stands for the ingress configs
	// +optional
	Ingress *IngressConfig `json:"ingress"`
	// ETCD stands for the standalone etcd configs
	// +optional
	ETCD *ETCDConfig `json:"etcd,omitempty"`

	// Cloud stands for the cloud connecting configs
	Cloud *CloudConfig `json:"cloud,omitempty"`

	// PlatformPublicURL stands the platform's public url
	PlatformPublicURL string `json:"platformPublicURL,omitempty"`

	// ProductBrand values: Standard、Runtime
	PackageEdition string `json:"packageEdition,omitempty"`

	// ProductBrand values: TKE、ACP、Common
	// +optional
	ProductBrand string `json:"productBrand,omitempty"`
}

// ProductBasePhase ...
// +kubebuilder:validation:Enum=Pending;Provisioning;Deleting;Failed;Running
type ProductBasePhase string

const (
	// BasePending ...
	BasePending ProductBasePhase = "Pending"
	// BaseProvisioning ...
	BaseProvisioning ProductBasePhase = "Provisioning"
	// BaseDeleting ...
	BaseDeleting ProductBasePhase = "Deleting"
	// BaseFailed ...
	BaseFailed ProductBasePhase = "Failed"
	// BaseRunning ...
	BaseRunning ProductBasePhase = "Running"
)

// ProductBaseConditionType ...
type ProductBaseConditionType string

const (
	// ReadyCondition ...
	ReadyCondition ProductBaseConditionType = "Ready"
	// InitializedCondition ...
	InitializedCondition ProductBaseConditionType = "Initialized"
)

// ProductBaseCondition ...
type ProductBaseCondition struct {
	// Type stands for Condition's type
	Type ProductBaseConditionType `json:"type"`
	// Status stands for Condition's status
	Status corev1.ConditionStatus `json:"status"`
	// Reason stands for Condition's reason
	// +optional
	Reason string `json:"reason"`
	// Message stands for Condition's message
	// +optional
	Message string `json:"message"`
	// LastTransitionTime stands for the timestamp of Condition's status changed
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// LastHeartbeatTime stands for the timestamp to check Condition
	// +optional
	LastHeartbeatTime *metav1.Time `json:"lastHeartbeatTime,omitempty"`
}

// DeployStatus ...
// +kubebuilder:validation:Enum=NotDeployed;Pending;Deploying;DeployFailed;Upgrading;UpgradeFailed;Completed
type DeployStatus string

const (
	// NotDeployed ...
	NotDeployed DeployStatus = "NotDeployed"
	// DeployPending ...
	DeployPending DeployStatus = "Pending"
	// Deploying ...
	Deploying DeployStatus = "Deploying"
	// DeployFailed ...
	DeployFailed DeployStatus = "DeployFailed"
	// Upgrading ...
	Upgrading DeployStatus = "Upgrading"
	// UpgradeFailed ...
	UpgradeFailed DeployStatus = "UpgradeFailed"
	// DeployCompleted ...
	DeployCompleted DeployStatus = "Completed"
)

// RunningStatus ...
// +kubebuilder:validation:Enum=NotRunning;Healthy;Abnormal
type RunningStatus string

const (
	// NotRunning ...
	NotRunning RunningStatus = "NotRunning"
	// RunningHealthy ...
	RunningHealthy RunningStatus = "Healthy"
	// RunningAbnormal ...
	RunningAbnormal RunningStatus = "Abnormal"
)

// MultiLaguageText stands for text with different lanaguaes
type MultiLaguageText struct {
	// ZH stands for text with Simple Chinese
	// +optional
	ZH string `json:"zh,omitempty"`
	// EN stands for text with English
	// +optional
	EN string `json:"en,omitempty"`
}

// ComponentStatus ...
type ComponentStatus struct {
	// DeployStatus stands for the deploy status
	DeployStatus DeployStatus `json:"deployStatus"`
	// RunningStatus stands for the running status
	RunningStatus RunningStatus `json:"runningStatus"`
	// Reason contains the reason why in current status
	// +optional
	Reason string `json:"reason"`
	// Message contains the detail infomation why in current status
	// +optional
	Message string `json:"message"`
	// LastTransitionTime stands for the timestamp of the status changed
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// LastProbeTime stands for the timestamp to probe status
	// +optional
	LastProbeTime *metav1.Time `json:"lastProbeTime,omitempty"`
	// DeployTime stands for the timestamp to deploy this component
	// +optional
	DeployTime *metav1.Time `json:"deployTime,omitempty"`
	// Entrypoint is the entrypoint of this component
	// +optional
	Entrypoint string `json:"entrypoint,omitempty"`
	// ProdutctType stands for the product type
	// +optional
	ProdutctType string `json:"productType,omitempty"`
	// LogPath stands for the path of log
	// +optional
	LogPath string `json:"logPath,omitempty"`
	// Logo stands for the product logo with base64 encoded data uri
	// +optional
	Logo string `json:"logo,omitempty"`
	// DisplayName stands for the product's display name
	// +optional
	DisplayName MultiLaguageText `json:"displayName,omitempty"`
	// Description stands for the product's description
	// +optional
	Description MultiLaguageText `json:"description,omitempty"`
	// Brief stands for the product's brief
	// +optional
	Brief MultiLaguageText `json:"brief,omitempty"`
}

// AppReleaseStatus ...
type AppReleaseStatus struct {
	// Synced stands for if AppRelease synced
	Synced bool `json:"synced"`
	// Synced stands for if AppRelease ready
	Ready bool `json:"ready"`
	// Synced stands for if AppRelease failed
	Failed bool `json:"failed"`
	// Reason stands for AppRelease's reason
	// +optional
	Reason string `json:"reason"`
	// Reason stands for AppRelease's message
	// +optional
	Message string `json:"message"`
	// LastProbeTime stands for the timestamp to probe AppRelease's status
	// +optional
	LastProbeTime *metav1.Time `json:"lastProbeTime,omitempty"`
}

// ProductBaseStatus defines the observed state of ProductBase
type ProductBaseStatus struct {
	// Phase stands for this product's phase
	Phase ProductBasePhase `json:"phase"`
	// Conditions contains this product's conditions
	Conditions []ProductBaseCondition `json:"conditions,omitempty"`
	// AppReleases records this product's AppReleases' status
	// +optional
	AppReleases map[string]AppReleaseStatus `json:"appReleases,omitempty"`
	// Version statnds for this product's current version
	Version string `json:"version"`
	// Deletable tell if this product can be delete
	Deletable bool `json:"deletable"`
	// ProductEntrypoint stands for this product's entrypoint
	// +optional
	ProductEntrypoint *string `json:"productEntrypoint,omitempty"`
	// Base stands for the base component's status
	// +optional
	Base *ComponentStatus `json:"base,omitempty"`
	// ContainerPlatform stands for the ContainerPlatform component's status
	// +optional
	ContainerPlatform *ComponentStatus `json:"containerPlatform,omitempty"`
	// Devops stands for the Devops component's status
	// +optional
	Devops *ComponentStatus `json:"devops,omitempty"`
	// ServiceMesh stands for the ServiceMesh component's status
	// +optional
	ServiceMesh *ComponentStatus `json:"serviceMesh,omitempty"`
	// DataServices stands for the DataServices component's status
	// +optional
	DataServices *ComponentStatus `json:"dataServices,omitempty"`

	// Cloud stands for the cloud connecting status
	Cloud *CloudStatus `json:"cloud,omitempty"`

	// GlobalID stands for the global Id
	GlobalID string `json:"globalID,omitempty"`

	// Progress stands for the progress of this product
	// +optional
	Progress int `json:"progress,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName="prdb",singular="productbase"
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.status.version`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=='Ready')].status`

// ProductBase is the Schema for the productbases API
type ProductBase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProductBaseSpec   `json:"spec,omitempty"`
	Status ProductBaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductBaseList contains a list of ProductBase
type ProductBaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductBase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProductBase{}, &ProductBaseList{})
}
