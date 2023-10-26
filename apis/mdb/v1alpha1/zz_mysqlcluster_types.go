/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionLimitsInitParameters struct {

	// Max connections per hour.
	MaxConnectionsPerHour *float64 `json:"maxConnectionsPerHour,omitempty" tf:"max_connections_per_hour,omitempty"`

	// Max questions per hour.
	MaxQuestionsPerHour *float64 `json:"maxQuestionsPerHour,omitempty" tf:"max_questions_per_hour,omitempty"`

	// Max updates per hour.
	MaxUpdatesPerHour *float64 `json:"maxUpdatesPerHour,omitempty" tf:"max_updates_per_hour,omitempty"`

	// Max user connections.
	MaxUserConnections *float64 `json:"maxUserConnections,omitempty" tf:"max_user_connections,omitempty"`
}

type ConnectionLimitsObservation struct {

	// Max connections per hour.
	MaxConnectionsPerHour *float64 `json:"maxConnectionsPerHour,omitempty" tf:"max_connections_per_hour,omitempty"`

	// Max questions per hour.
	MaxQuestionsPerHour *float64 `json:"maxQuestionsPerHour,omitempty" tf:"max_questions_per_hour,omitempty"`

	// Max updates per hour.
	MaxUpdatesPerHour *float64 `json:"maxUpdatesPerHour,omitempty" tf:"max_updates_per_hour,omitempty"`

	// Max user connections.
	MaxUserConnections *float64 `json:"maxUserConnections,omitempty" tf:"max_user_connections,omitempty"`
}

type ConnectionLimitsParameters struct {

	// Max connections per hour.
	// +kubebuilder:validation:Optional
	MaxConnectionsPerHour *float64 `json:"maxConnectionsPerHour,omitempty" tf:"max_connections_per_hour,omitempty"`

	// Max questions per hour.
	// +kubebuilder:validation:Optional
	MaxQuestionsPerHour *float64 `json:"maxQuestionsPerHour,omitempty" tf:"max_questions_per_hour,omitempty"`

	// Max updates per hour.
	// +kubebuilder:validation:Optional
	MaxUpdatesPerHour *float64 `json:"maxUpdatesPerHour,omitempty" tf:"max_updates_per_hour,omitempty"`

	// Max user connections.
	// +kubebuilder:validation:Optional
	MaxUserConnections *float64 `json:"maxUserConnections,omitempty" tf:"max_user_connections,omitempty"`
}

type MySQLClusterAccessInitParameters struct {

	// Allow access for Yandex DataLens.
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// Allows access for SQL queries in the management console.
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type MySQLClusterAccessObservation struct {

	// Allow access for Yandex DataLens.
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// Allows access for SQL queries in the management console.
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type MySQLClusterAccessParameters struct {

	// Allow access for Yandex DataLens.
	// +kubebuilder:validation:Optional
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// Allows access for SQL queries in the management console.
	// +kubebuilder:validation:Optional
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type MySQLClusterBackupWindowStartInitParameters struct {

	// The hour at which backup will be started.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type MySQLClusterBackupWindowStartObservation struct {

	// The hour at which backup will be started.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type MySQLClusterBackupWindowStartParameters struct {

	// The hour at which backup will be started.
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type MySQLClusterDatabaseInitParameters struct {

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MySQLClusterDatabaseObservation struct {

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MySQLClusterDatabaseParameters struct {

	// The name of the database.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type MySQLClusterHostInitParameters struct {

	// Sets whether the host should get a public IP address. It can be changed on the fly only when name is set.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Host backup priority. Value is between 0 and 100, default is 0.
	BackupPriority *float64 `json:"backupPriority,omitempty" tf:"backup_priority,omitempty"`

	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to replication_source_name parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Host master promotion priority. Value is between 0 and 100, default is 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Host replication source name points to host's name from which this host should replicate. When not set then host in HA group. It works only when name is set.
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	// The availability zone where the MySQL host will be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MySQLClusterHostObservation struct {

	// Sets whether the host should get a public IP address. It can be changed on the fly only when name is set.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Host backup priority. Value is between 0 and 100, default is 0.
	BackupPriority *float64 `json:"backupPriority,omitempty" tf:"backup_priority,omitempty"`

	// (Computed) The fully qualified domain name of the host.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to replication_source_name parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Host master promotion priority. Value is between 0 and 100, default is 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (Computed) Host replication source (fqdn), when replication_source is empty then host is in HA group.
	ReplicationSource *string `json:"replicationSource,omitempty" tf:"replication_source,omitempty"`

	// Host replication source name points to host's name from which this host should replicate. When not set then host in HA group. It works only when name is set.
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The availability zone where the MySQL host will be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MySQLClusterHostParameters struct {

	// Sets whether the host should get a public IP address. It can be changed on the fly only when name is set.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Host backup priority. Value is between 0 and 100, default is 0.
	// +kubebuilder:validation:Optional
	BackupPriority *float64 `json:"backupPriority,omitempty" tf:"backup_priority,omitempty"`

	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to replication_source_name parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Host master promotion priority. Value is between 0 and 100, default is 0.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Host replication source name points to host's name from which this host should replicate. When not set then host in HA group. It works only when name is set.
	// +kubebuilder:validation:Optional
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The availability zone where the MySQL host will be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type MySQLClusterInitParameters struct {

	// Access policy to the MySQL cluster. The structure is documented below.
	Access []MySQLClusterAccessInitParameters `json:"access,omitempty" tf:"access,omitempty"`

	// A host of the MySQL cluster. The structure is documented below.
	AllowRegenerationHost *bool `json:"allowRegenerationHost,omitempty" tf:"allow_regeneration_host,omitempty"`

	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart []MySQLClusterBackupWindowStartInitParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// (Deprecated) To manage databases, please switch to using a separate resource type yandex_mdb_mysql_databases.
	Database []MySQLClusterDatabaseInitParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the MySQL cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the MySQL cluster.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// A host of the MySQL cluster. The structure is documented below.
	Host []MySQLClusterHostInitParameters `json:"host,omitempty" tf:"host,omitempty"`

	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow []MySQLClusterMaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MySQLConfig map[string]*string `json:"mysqlConfig,omitempty" tf:"mysql_config,omitempty"`

	// Name of the MySQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Cluster performance diagnostics settings. The structure is documented below. YC Documentation
	PerformanceDiagnostics []PerformanceDiagnosticsInitParameters `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources []MySQLClusterResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// The cluster will be created from the specified backup. The structure is documented below.
	Restore []RestoreInitParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// (Deprecated) To manage users, please switch to using a separate resource type yandex_mdb_mysql_user.
	User []MySQLClusterUserInitParameters `json:"user,omitempty" tf:"user,omitempty"`

	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MySQLClusterMaintenanceWindowInitParameters struct {

	// Day of the week (in DDD format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of the day in UTC (in HH format). Allowed value is between 0 and 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MySQLClusterMaintenanceWindowObservation struct {

	// Day of the week (in DDD format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of the day in UTC (in HH format). Allowed value is between 0 and 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MySQLClusterMaintenanceWindowParameters struct {

	// Day of the week (in DDD format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of the day in UTC (in HH format). Allowed value is between 0 and 23.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MySQLClusterObservation struct {

	// Access policy to the MySQL cluster. The structure is documented below.
	Access []MySQLClusterAccessObservation `json:"access,omitempty" tf:"access,omitempty"`

	// A host of the MySQL cluster. The structure is documented below.
	AllowRegenerationHost *bool `json:"allowRegenerationHost,omitempty" tf:"allow_regeneration_host,omitempty"`

	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart []MySQLClusterBackupWindowStartObservation `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// Creation timestamp of the cluster.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Deprecated) To manage databases, please switch to using a separate resource type yandex_mdb_mysql_databases.
	Database []MySQLClusterDatabaseObservation `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the MySQL cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the MySQL cluster.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Aggregated health of the cluster.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// A host of the MySQL cluster. The structure is documented below.
	Host []MySQLClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the MySQL cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow []MySQLClusterMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MySQLConfig map[string]*string `json:"mysqlConfig,omitempty" tf:"mysql_config,omitempty"`

	// Name of the MySQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the MySQL cluster uses.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Cluster performance diagnostics settings. The structure is documented below. YC Documentation
	PerformanceDiagnostics []PerformanceDiagnosticsObservation `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources []MySQLClusterResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// The cluster will be created from the specified backup. The structure is documented below.
	Restore []RestoreObservation `json:"restore,omitempty" tf:"restore,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Status of the cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Deprecated) To manage users, please switch to using a separate resource type yandex_mdb_mysql_user.
	User []MySQLClusterUserObservation `json:"user,omitempty" tf:"user,omitempty"`

	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MySQLClusterParameters struct {

	// Access policy to the MySQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Access []MySQLClusterAccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// A host of the MySQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	AllowRegenerationHost *bool `json:"allowRegenerationHost,omitempty" tf:"allow_regeneration_host,omitempty"`

	// Time to start the daily backup, in the UTC. The structure is documented below.
	// +kubebuilder:validation:Optional
	BackupWindowStart []MySQLClusterBackupWindowStartParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// (Deprecated) To manage databases, please switch to using a separate resource type yandex_mdb_mysql_databases.
	// +kubebuilder:validation:Optional
	Database []MySQLClusterDatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the MySQL cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the MySQL cluster.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A host of the MySQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Host []MySQLClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// A set of key/value label pairs to assign to the MySQL cluster.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy of the MySQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MySQLClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// MySQL cluster config. Detail info in "MySQL config" section (documented below).
	// +kubebuilder:validation:Optional
	MySQLConfig map[string]*string `json:"mysqlConfig,omitempty" tf:"mysql_config,omitempty"`

	// Name of the MySQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the MySQL cluster uses.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Cluster performance diagnostics settings. The structure is documented below. YC Documentation
	// +kubebuilder:validation:Optional
	PerformanceDiagnostics []PerformanceDiagnosticsParameters `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Resources []MySQLClusterResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// The cluster will be created from the specified backup. The structure is documented below.
	// +kubebuilder:validation:Optional
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// (Deprecated) To manage users, please switch to using a separate resource type yandex_mdb_mysql_user.
	// +kubebuilder:validation:Optional
	User []MySQLClusterUserParameters `json:"user,omitempty" tf:"user,omitempty"`

	// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MySQLClusterResourcesInitParameters struct {

	// Volume of the storage available to a MySQL host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of MySQL hosts.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type MySQLClusterResourcesObservation struct {

	// Volume of the storage available to a MySQL host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of MySQL hosts.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type MySQLClusterResourcesParameters struct {

	// Volume of the storage available to a MySQL host, in gigabytes.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// Type of the storage of MySQL hosts.
	// +kubebuilder:validation:Optional
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type MySQLClusterUserInitParameters struct {

	// Authentication plugin. Allowed values: MYSQL_NATIVE_PASSWORD, CACHING_SHA2_PASSWORD, SHA256_PASSWORD (for version 5.7 MYSQL_NATIVE_PASSWORD, SHA256_PASSWORD)
	AuthenticationPlugin *string `json:"authenticationPlugin,omitempty" tf:"authentication_plugin,omitempty"`

	// User's connection limits. The structure is documented below.
	// If the attribute is not specified there will be no changes.
	ConnectionLimits []ConnectionLimitsInitParameters `json:"connectionLimits,omitempty" tf:"connection_limits,omitempty"`

	// List user's global permissions
	// Allowed permissions:  REPLICATION_CLIENT, REPLICATION_SLAVE, PROCESS for clear list use empty list.
	// If the attribute is not specified there will be no changes.
	GlobalPermissions []*string `json:"globalPermissions,omitempty" tf:"global_permissions,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []UserPermissionInitParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

type MySQLClusterUserObservation struct {

	// Authentication plugin. Allowed values: MYSQL_NATIVE_PASSWORD, CACHING_SHA2_PASSWORD, SHA256_PASSWORD (for version 5.7 MYSQL_NATIVE_PASSWORD, SHA256_PASSWORD)
	AuthenticationPlugin *string `json:"authenticationPlugin,omitempty" tf:"authentication_plugin,omitempty"`

	// User's connection limits. The structure is documented below.
	// If the attribute is not specified there will be no changes.
	ConnectionLimits []ConnectionLimitsObservation `json:"connectionLimits,omitempty" tf:"connection_limits,omitempty"`

	// List user's global permissions
	// Allowed permissions:  REPLICATION_CLIENT, REPLICATION_SLAVE, PROCESS for clear list use empty list.
	// If the attribute is not specified there will be no changes.
	GlobalPermissions []*string `json:"globalPermissions,omitempty" tf:"global_permissions,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []UserPermissionObservation `json:"permission,omitempty" tf:"permission,omitempty"`
}

type MySQLClusterUserParameters struct {

	// Authentication plugin. Allowed values: MYSQL_NATIVE_PASSWORD, CACHING_SHA2_PASSWORD, SHA256_PASSWORD (for version 5.7 MYSQL_NATIVE_PASSWORD, SHA256_PASSWORD)
	// +kubebuilder:validation:Optional
	AuthenticationPlugin *string `json:"authenticationPlugin,omitempty" tf:"authentication_plugin,omitempty"`

	// User's connection limits. The structure is documented below.
	// If the attribute is not specified there will be no changes.
	// +kubebuilder:validation:Optional
	ConnectionLimits []ConnectionLimitsParameters `json:"connectionLimits,omitempty" tf:"connection_limits,omitempty"`

	// List user's global permissions
	// Allowed permissions:  REPLICATION_CLIENT, REPLICATION_SLAVE, PROCESS for clear list use empty list.
	// If the attribute is not specified there will be no changes.
	// +kubebuilder:validation:Optional
	GlobalPermissions []*string `json:"globalPermissions,omitempty" tf:"global_permissions,omitempty"`

	// The name of the user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The password of the user.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of permissions granted to the user. The structure is documented below.
	// +kubebuilder:validation:Optional
	Permission []UserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

type PerformanceDiagnosticsInitParameters struct {

	// Enable performance diagnostics
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval (in seconds) for my_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval,omitempty" tf:"sessions_sampling_interval,omitempty"`

	// Interval (in seconds) for my_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval,omitempty" tf:"statements_sampling_interval,omitempty"`
}

type PerformanceDiagnosticsObservation struct {

	// Enable performance diagnostics
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval (in seconds) for my_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval,omitempty" tf:"sessions_sampling_interval,omitempty"`

	// Interval (in seconds) for my_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval,omitempty" tf:"statements_sampling_interval,omitempty"`
}

type PerformanceDiagnosticsParameters struct {

	// Enable performance diagnostics
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval (in seconds) for my_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	// +kubebuilder:validation:Optional
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval" tf:"sessions_sampling_interval,omitempty"`

	// Interval (in seconds) for my_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	// +kubebuilder:validation:Optional
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval" tf:"statements_sampling_interval,omitempty"`
}

type RestoreInitParameters struct {

	// Backup ID. The cluster will be created from the specified backup. How to get a list of MySQL backups.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Timestamp of the moment to which the MySQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type RestoreObservation struct {

	// Backup ID. The cluster will be created from the specified backup. How to get a list of MySQL backups.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Timestamp of the moment to which the MySQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type RestoreParameters struct {

	// Backup ID. The cluster will be created from the specified backup. How to get a list of MySQL backups.
	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId" tf:"backup_id,omitempty"`

	// Timestamp of the moment to which the MySQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	// +kubebuilder:validation:Optional
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type UserPermissionInitParameters struct {

	// The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// List user's roles in the database.
	// Allowed roles: ALL,ALTER,ALTER_ROUTINE,CREATE,CREATE_ROUTINE,CREATE_TEMPORARY_TABLES,
	// CREATE_VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK_TABLES,SELECT,SHOW_VIEW,TRIGGER,UPDATE.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type UserPermissionObservation struct {

	// The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// List user's roles in the database.
	// Allowed roles: ALL,ALTER,ALTER_ROUTINE,CREATE,CREATE_ROUTINE,CREATE_TEMPORARY_TABLES,
	// CREATE_VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK_TABLES,SELECT,SHOW_VIEW,TRIGGER,UPDATE.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type UserPermissionParameters struct {

	// The name of the database that the permission grants access to.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// List user's roles in the database.
	// Allowed roles: ALL,ALTER,ALTER_ROUTINE,CREATE,CREATE_ROUTINE,CREATE_TEMPORARY_TABLES,
	// CREATE_VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK_TABLES,SELECT,SHOW_VIEW,TRIGGER,UPDATE.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

// MySQLClusterSpec defines the desired state of MySQLCluster
type MySQLClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLClusterParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MySQLClusterInitParameters `json:"initProvider,omitempty"`
}

// MySQLClusterStatus defines the observed state of MySQLCluster.
type MySQLClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLCluster is the Schema for the MySQLClusters API. Manages a MySQL cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MySQLCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || has(self.initProvider.environment)",message="environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.host) || has(self.initProvider.host)",message="host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resources) || has(self.initProvider.resources)",message="resources is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || has(self.initProvider.version)",message="version is a required parameter"
	Spec   MySQLClusterSpec   `json:"spec"`
	Status MySQLClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLClusterList contains a list of MySQLClusters
type MySQLClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLCluster `json:"items"`
}

// Repository type metadata.
var (
	MySQLCluster_Kind             = "MySQLCluster"
	MySQLCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLCluster_Kind}.String()
	MySQLCluster_KindAPIVersion   = MySQLCluster_Kind + "." + CRDGroupVersion.String()
	MySQLCluster_GroupVersionKind = CRDGroupVersion.WithKind(MySQLCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLCluster{}, &MySQLClusterList{})
}
