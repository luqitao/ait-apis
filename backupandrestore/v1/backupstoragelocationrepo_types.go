package v1

import (
	corev1api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupStorageLocationRepoSpec defines the desired state of BackupStorageLocationRepo
type BackupStorageLocationRepoSpec struct {
	// Provider is the provider of the backup storage.
	// +optional
	Provider string `json:"provider"`

	// Config is for provider-specific configuration fields.
	// +optional
	Config map[string]string `json:"config,omitempty"`

	// Credential contains the credential information intended to be used with this location
	// +optional
	Credential *corev1api.SecretKeySelector `json:"credential,omitempty"`

	StorageType `json:",inline"`

	// Default indicates this location is the default backup storage location.
	// +optional
	Default bool `json:"default,omitempty"`

	// AccessMode defines the permissions for the backup storage location.
	// +optional
	AccessMode BackupStorageLocationAccessMode `json:"accessMode,omitempty"`

	// BackupSyncPeriod defines how frequently to sync backup API objects from object storage. A value of 0 disables sync.
	// +optional
	// +nullable
	BackupSyncPeriod *metav1.Duration `json:"backupSyncPeriod,omitempty"`
}

// BackupStorageLocationRepoStatus defines the observed state of BackupStorageLocationRepo
type BackupStorageLocationRepoStatus struct {
	// SyncedClusters is the current synced clusters.
	// +optional
	SyncedClusters []string `json:"syncedClusters,omitempty"`

	// LastSyncedTime is the last time the contents of the location were synced into
	// the cluster.
	// +optional
	// +nullable
	LastSyncedTime *metav1.Time `json:"lastSyncedTime,omitempty"`
}

// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=bslr

// BackupStorageLocationRepo is the Schema for the BackupStorageLocationRepo API
type BackupStorageLocationRepo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupStorageLocationRepoSpec   `json:"spec,omitempty"`
	Status BackupStorageLocationRepoStatus `json:"status,omitempty"`
}

// the k8s:deepcopy marker will no longer be needed and should be removed.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:rbac:groups=ait.velero.io,resources=backupstoragelocationrepos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ait.velero.io,resources=backupstoragelocationrepos/status,verbs=get;update;patch

// BackupStorageLocationRepoList contains a list of BackupStorageLocationRepo
type BackupStorageLocationRepoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupStorageLocationRepo `json:"items"`
}

// StorageType represents the type of storage that a backup location uses.
// ObjectStorage must be non-nil, since it is currently the only supported StorageType.
type StorageType struct {
	ObjectStorage *ObjectStorageLocation `json:"objectStorage"`
}

// BackupStorageLocationAccessMode represents the permissions for a BackupStorageLocation.
// +kubebuilder:validation:Enum=ReadOnly;ReadWrite
type BackupStorageLocationAccessMode string

const (
	// BackupStorageLocationAccessModeReadOnly represents read-only access to a BackupStorageLocation.
	BackupStorageLocationAccessModeReadOnly BackupStorageLocationAccessMode = "ReadOnly"

	// BackupStorageLocationAccessModeReadWrite represents read and write access to a BackupStorageLocation.
	BackupStorageLocationAccessModeReadWrite BackupStorageLocationAccessMode = "ReadWrite"
)

// ObjectStorageLocation specifies the settings necessary to connect to a provider's object storage.
type ObjectStorageLocation struct {
	// Bucket is the bucket to use for object storage.
	Bucket string `json:"bucket"`

	// Prefix is the path inside a bucket to use for Velero storage. Optional.
	// +optional
	Prefix string `json:"prefix,omitempty"`

	// CACert defines a CA bundle to use when verifying TLS connections to the provider.
	// +optional
	CACert []byte `json:"caCert,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BackupStorageLocationRepo{}, &BackupStorageLocationRepoList{})
}
