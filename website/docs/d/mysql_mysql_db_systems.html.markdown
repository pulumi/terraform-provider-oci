---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_db_systems"
sidebar_current: "docs-oci-datasource-mysql-mysql_db_systems"
description: |-
  Provides the list of Mysql Db Systems in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_mysql_db_systems
This data source provides the list of Mysql Db Systems in Oracle Cloud Infrastructure MySQL Database service.

Get a list of DB Systems in the specified compartment.
The default sort order is by timeUpdated, descending.


## Example Usage

```hcl
data "oci_mysql_mysql_db_systems" "test_mysql_db_systems" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	configuration_id = var.mysql_configuration_id
	db_system_id = oci_mysql_mysql_db_system.test_db_system.id
	display_name = var.mysql_db_system_display_name
	is_analytics_cluster_attached = var.mysql_db_system_is_analytics_cluster_attached
	is_heat_wave_cluster_attached = var.mysql_db_system_is_heat_wave_cluster_attached
	is_up_to_date = var.mysql_db_system_is_up_to_date
	state = var.mysql_db_system_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `configuration_id` - (Optional) The requested Configuration instance.
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `is_analytics_cluster_attached` - (Optional) DEPRECATED -- please use HeatWave API instead. If true, return only DB Systems with an Analytics Cluster attached, if false return only DB Systems with no Analytics Cluster attached. If not present, return all DB Systems. 
* `is_heat_wave_cluster_attached` - (Optional) If true, return only DB Systems with a HeatWave cluster attached, if false return only DB Systems with no HeatWave cluster attached. If not present, return all DB Systems. 
* `is_up_to_date` - (Optional) Filter instances if they are using the latest revision of the Configuration they are associated with. 
* `state` - (Optional) DbSystem Lifecycle State


## Attributes Reference

The following attributes are exported:

* `db_systems` - The list of db_systems.

### MysqlDbSystem Reference

The following attributes are exported:

* `analytics_cluster` - DEPRECATED -- please use HeatWave API instead. A summary of an Analytics Cluster. 
	* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the Analytics Cluster. 
	* `shape_name` - The shape determines resources to allocate to the Analytics Cluster nodes - CPU cores, memory. 
	* `state` - The current state of the MySQL Analytics Cluster.
	* `time_created` - The date and time the Analytics Cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
	* `time_updated` - The time the Analytics Cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `availability_domain` - The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the availability domain in which the DB System is placed. 
* `backup_policy` - The Backup policy for the DB System.
	* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.

		Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.

		Example: `{"bar-key": "value"}` 
	* `is_enabled` - If automated backups are enabled or disabled.
	* `pitr_policy` - The PITR policy for the DB System.
		* `is_enabled` - Specifies if PITR is enabled or disabled.
	* `retention_in_days` - The number of days automated backups are retained. 
	* `window_start_time` - The start of a 30-minute window of time in which daily, automated backups occur.

		This should be in the format of the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.

		At some point in the window, the system may incur a brief service disruption as the backup is performed.

		If not defined, a window is selected from the following Region-based time-spans:
		* eu-frankfurt-1: 20:00 - 04:00 UTC
		* us-ashburn-1: 03:00 - 11:00 UTC
		* uk-london-1: 06:00 - 14:00 UTC
		* ap-tokyo-1: 13:00 - 21:00
		* us-phoenix-1: 06:00 - 14:00 
* `channels` - A list with a summary of all the Channels attached to the DB System.
	* `compartment_id` - The OCID of the compartment.
	* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - The user-friendly name for the Channel. It does not have to be unique.
	* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the Channel.
	* `is_enabled` - Whether the Channel has been enabled by the user.
	* `lifecycle_details` - A message describing the state of the Channel.
	* `source` - Parameters detailing how to provision the source for the given Channel.
		* `hostname` - The network address of the MySQL instance.
		* `port` - The port the source MySQL instance listens on.
		* `source_type` - The specific source identifier.
		* `ssl_ca_certificate` - The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
			* `certificate_type` - The type of CA certificate.
			* `contents` - The string containing the CA certificate in PEM format.
		* `ssl_mode` - The SSL mode of the Channel.
		* `username` - The name of the replication user on the source MySQL instance. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
	* `state` - The state of the Channel.
	* `target` - Details about the Channel target.
		* `applier_username` - The username for the replication applier of the target MySQL DB System.
		* `channel_name` - The case-insensitive name that identifies the replication channel. Channel names must follow the rules defined for [MySQL identifiers](https://dev.mysql.com/doc/refman/8.0/en/identifiers.html). The names of non-Deleted Channels must be unique for each DB System. 
		* `db_system_id` - The OCID of the source DB System.
		* `target_type` - The specific target identifier.
	* `time_created` - The date and time the Channel was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The time the Channel was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `compartment_id` - The OCID of the compartment the DB System belongs in.
* `configuration_id` - The OCID of the Configuration to be used for Instances in this DB System.
* `crash_recovery` - Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled, and whether to enable or disable syncing of the Binary Logs. 
* `current_placement` - The availability domain and fault domain a DB System is placed in.
	* `availability_domain` - The availability domain in which the DB System is placed.
	* `fault_domain` - The fault domain in which the DB System is placed.
* `data_storage_size_in_gb` - Initial size of the data volume in GiBs that will be created and attached. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deletion_policy` - The Deletion policy for the DB System.
	* `automatic_backup_retention` - Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted. 
	* `final_backup` - Specifies whether or not a backup is taken when the DB System is deleted. REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted. SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted. 
	* `is_delete_protected` - Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow. 
* `description` - User-provided data about the DB System.
* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
* `endpoints` - The network endpoints available for this DB System. 
	* `hostname` - The network address of the DB System.
	* `ip_address` - The IP address the DB System is configured to listen on.
	* `modes` - The access modes from the client that this endpoint supports.
	* `port` - The port the MySQL instance listens on.
	* `port_x` - The network port where to connect to use this endpoint using the X protocol.
	* `status` - The state of the endpoints, as far as it can seen from the DB System. There may be some inconsistency with the actual state of the MySQL service. 
	* `status_details` - Additional information about the current endpoint status.
* `fault_domain` - The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.

	In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains and the MySQL instance in that domain is promoted to the primary instance. This redirection does not affect the IP address of the DB System in any way.

	For a standalone DB System, this defines the fault domain in which the DB System is placed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `heat_wave_cluster` - A summary of a HeatWave cluster. 
	* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster. 
	* `shape_name` - The shape determines resources to allocate to the HeatWave nodes - CPU cores, memory. 
	* `state` - The current state of the MySQL HeatWave cluster.
	* `time_created` - The date and time the HeatWave cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The time the HeatWave cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `hostname_label` - The hostname for the primary endpoint of the DB System. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com"). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. 
* `id` - The OCID of the DB System.
* `ip_address` - The IP address the DB System is configured to listen on. A private IP address of the primary endpoint of the DB System. Must be an available IP address within the subnet's CIDR. This will be a "dotted-quad" style IPv4 address. 
* `is_analytics_cluster_attached` - DEPRECATED -- please use `isHeatWaveClusterAttached` instead. If the DB System has an Analytics Cluster attached. 
* `is_heat_wave_cluster_attached` - If the DB System has a HeatWave Cluster attached. 
* `is_highly_available` - Specifies if the DB System is highly available. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance` - The Maintenance Policy for the DB System. 
	* `window_start_time` - The start time of the maintenance window.

		This string is of the format: "{day-of-week} {time-of-day}".

		"{day-of-week}" is a case-insensitive string like "mon", "tue", &c.

		"{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero. 
* `mysql_version` - Name of the MySQL Version in use for the DB System.
* `point_in_time_recovery_details` - Point-in-time Recovery details like earliest and latest recovery time point for the DB System. 
	* `time_earliest_recovery_point` - Earliest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_latest_recovery_point` - Latest recovery time point for the DB System, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `port` - The port for primary endpoint of the DB System to listen on.
* `port_x` - The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. 
* `shape_name` - The shape of the primary instances of the DB System. The shape determines resources allocated to a DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use (the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20181021/ShapeSummary/ListShapes) operation. 
* `source` - Parameters detailing how to provision the initial data of the DB System. 
	* `backup_id` - The OCID of the backup to be used as the source for the new DB System. 
	* `db_system_id` - The OCID of the DB System from which a backup shall be selected to be restored when creating the new DB System. Use this together with recovery point to perform a point in time recovery operation. 
	* `recovery_point` - The date and time, as per RFC 3339, of the change up to which the new DB System shall be restored to, using a backup and logs from the original DB System. In case no point in time is specified, then this new DB System shall be restored up to the latest change recorded for the original DB System. 
	* `source_type` - The specific source identifier. 
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with. 
* `time_created` - The date and time the DB System was created.
* `time_updated` - The time the DB System was last updated.

