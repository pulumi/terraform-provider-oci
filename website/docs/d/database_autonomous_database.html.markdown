---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database"
sidebar_current: "docs-oci-datasource-database-autonomous_database"
description: |-
  Provides details about a specific Autonomous Database in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database
This data source provides details about a specific Autonomous Database resource in Oracle Cloud Infrastructure Database service.

Gets the details of the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `actual_used_data_storage_size_in_tbs` - The current amount of storage in use for user and system data, in terabytes (TB). 
* `allocated_storage_size_in_tbs` - The amount of storage currently allocated for the database tables and billed for, rounded up. When auto-scaling is not enabled, this value is equal to the `dataStorageSizeInTBs` value. You can compare this value to the `actualUsedDataStorageSizeInTBs` value to determine if a manual shrink operation is appropriate for your allocated storage.

	**Note:** Auto-scaling does not automatically decrease allocated storage when data is deleted from the database. 
* `apex_details` - Information about Oracle APEX Application Development.
	* `apex_version` - The Oracle APEX Application Development version.
	* `ords_version` - The Oracle REST Data Services (ORDS) version.
* `are_primary_whitelisted_ips_used` - This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby. It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.	
* `autonomous_container_database_id` - The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_maintenance_schedule_type` - The maintenance schedule type of the Autonomous Database on shared Exadata infrastructure. The EARLY maintenance schedule of this Autonomous Database follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle. 
* `available_upgrade_versions` - List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
* `backup_config` - Autonomous Database configuration details for storing [manual backups](https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/backup-restore.html#GUID-9035DFB8-4702-4CEB-8281-C2A303820809) in the [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) service. 
	* `manual_backup_bucket_name` - Name of [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) bucket to use for storing manual backups.
	* `manual_backup_type` - The manual backup destination type.
* `character_set` - The character set for the autonomous database.  The default is AL32UTF8. Allowed values are:

	AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Database. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `dedicated` - The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
	* `profiles` - A list of connection string profiles to allow clients to group, filter and select connection string values based on structured metadata. 
		* `consumer_group` - Consumer group used by the connection.
		* `display_name` - A user-friendly name for the connection.
		* `host_format` - Host format used in connection string.
		* `protocol` - Protocol used by the connection.
		* `session_mode` - Specifies whether the listener performs a direct hand-off of the session, or redirects the session. In RAC deployments where SCAN is used, sessions are redirected to a Node VIP. Use `DIRECT` for direct hand-offs. Use `REDIRECT` to redirect the session.
		* `syntax_format` - Specifies whether the connection string is using the long (`LONG`), Easy Connect (`EZCONNECT`), or Easy Connect Plus (`EZCONNECTPLUS`) format. Autonomous Databases on shared Exadata infrastructure always use the long format. 
		* `tls_authentication` - Specifies whether the TLS handshake is using one-way (`SERVER`) or mutual (`MUTUAL`) authentication.
		* `value` - Connection string value.
* `connection_urls` - The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).  Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}` 
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `graph_studio_url` - The URL of the Graph Studio for the Autonomous Database.
	* `machine_learning_user_management_url` - Oracle Machine Learning user management URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details. 

	**Note:** This parameter cannot be used with the `ocpuCount` parameter. 
* `customer_contacts` - Customer Contacts.
	* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_gb` - The quantity of data in the database, in gigabytes.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `database_edition` - The Oracle Database Edition that applies to the Autonomous databases. 
* `database_management_status` - Status of Database Management for this Autonomous Database.
* `dataguard_region_type` - The Autonomous Data Guard region type of the Autonomous Database. For Autonomous Databases on shared Exadata infrastructure, Data Guard associations have designated primary and standby regions, and these region types do not change when the database changes roles. The standby regions in Data Guard associations can be the same region designated as the primary region, or they can be remote regions. Certain database administrative operations may be available only in the primary region of the Data Guard association, and cannot be performed when the database using the "primary" role is operating in a remote Data Guard standby region.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `failed_data_recovery_in_seconds` - Indicates the number of seconds of data loss for a Data Guard failover.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_access_control_enabled` - Indicates if the database-level access control is enabled. If disabled, database access is defined by the network security rules. If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional, if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console. When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.

	This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. 
* `is_auto_scaling_for_storage_enabled` - Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`. 
* `is_data_guard_enabled` - **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_local_data_guard_enabled` - Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `is_mtls_connection_required` - Indicates whether the Autonomous Database requires mTLS connections.
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `is_reconnect_clone_enabled` - Indicates if the refreshable clone can be reconnected to its source database.
* `is_refreshable_clone` - Indicates whether the Autonomous Database is a refreshable clone.
* `is_remote_data_guard_enabled` - Indicates whether the Autonomous Database has Cross Region Data Guard enabled. Not applicable to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `key_history_entry` - Key History Entry.
	* `id` - The id of the Autonomous Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
	* `time_activated` - The date and time the kms key activated.
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_lifecycle_details` - KMS key lifecycle details.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `lifecycle_details` - Information about the current lifecycle state.
* `local_standby_db` - Autonomous Data Guard standby database details. 
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
	* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per each OCPU core in Autonomous VM Cluster.
* `max_cpu_core_count` - The number of Max OCPU cores to be made available to the autonomous database with auto scaling of cpu enabled. 
* `ncharacter_set` - The national character set for the autonomous database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8.
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `ocpu_count` - The number of OCPU cores to be made available to the database. 

	The following points apply:
	* For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	* To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.

	For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

	**Note:** This parameter cannot be used with the `cpuCoreCount` parameter. 
* `open_mode` - The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
* `operations_insights_status` - Status of Operations Insights for this Autonomous Database.
* `peer_db_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of standby databases located in Autonomous Data Guard remote regions that are associated with the source database. Note that for shared Exadata infrastructure, standby databases located in the same region as the source primary database do not have OCIDs.
* `permission_level` - The Autonomous Database permission level. Restricted mode allows access only to admin users.
* `private_endpoint` - The private endpoint for the resource.
* `private_endpoint_ip` - The private endpoint Ip address for the resource.
* `private_endpoint_label` - The private endpoint label for the resource.
* `provisionable_cpus` - An array of CPU values that an Autonomous Database can be scaled to.
* `refreshable_mode` - The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `refreshable_status` - The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `scheduled_operations` - list of scheduled operations
	* `day_of_week` - Day of the week.
		* `name` - Name of the day of the week.
	* `scheduled_start_time` - auto start time. value must be of ISO-8601 format "HH:mm"
	* `scheduled_stop_time` - auto stop time. value must be of ISO-8601 format "HH:mm"
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
* `standby_db` - **Deprecated** Autonomous Data Guard standby database details. 
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
	* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
* `standby_whitelisted_ips` - The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. 
* `state` - The current state of the Autonomous Database.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `supported_regions_to_clone_to` - The list of regions that support the creation of an Autonomous Database clone or an Autonomous Data Guard standby database. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Autonomous Database was created.
* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the Autonomous Database. For databases that have standbys in both the primary Data Guard region and a remote Data Guard standby region, this is the latest timestamp of either the database using the "primary" role in the primary Data Guard region, or database located in the remote Data Guard standby region.
* `time_deletion_of_free_autonomous_database` - The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted. 
* `time_local_data_guard_enabled` - The date and time that Autonomous Data Guard was enabled for an Autonomous Database where the standby was provisioned in the same region as the primary database.
* `time_maintenance_begin` - The date and time when maintenance will begin.
* `time_maintenance_end` - The date and time when maintenance will end.
* `time_of_last_failover` - The timestamp of the last failover operation.
* `time_of_last_refresh` - The date and time when last refresh happened.
* `time_of_last_refresh_point` - The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.
* `time_of_last_switchover` - The timestamp of the last switchover operation for the Autonomous Database.
* `time_of_next_refresh` - The date and time of next refresh.
* `time_reclamation_of_free_autonomous_database` - The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state. 
* `time_until_reconnect_clone_enabled` - The time and date as an RFC3339 formatted string, e.g., 2022-01-01T12:00:00.000Z, to set the limit for a refreshable clone to be reconnected to its source database.
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. 

