---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_autonomous_vm_cluster"
sidebar_current: "docs-oci-datasource-database-cloud_autonomous_vm_cluster"
description: |-
  Provides details about a specific Cloud Autonomous Vm Cluster in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_autonomous_vm_cluster
This data source provides details about a specific Cloud Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Custustomer systems, see [GetAutonomousVmCluster ](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/AutonomousVmCluster/GetAutonomousVmCluster).


## Example Usage

```hcl
data "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
	#Required
	cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_autonomous_vm_cluster_id` - (Required) The Cloud VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `availability_domain` - The name of the availability domain that the cloud Autonomous VM cluster is located in.
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_container_databases` - The number of Autonomous Container Databases that can be created with the currently available local storage.
* `available_cpus` - CPU cores available for allocation to Autonomous Databases.
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_time_zone` - The time zone of the Cloud Autonomous VM Cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the cloud Autonomous VM cluster.
* `data_storage_size_in_gb` - The total data storage allocated, in gigabytes (GB).
* `data_storage_size_in_tbs` - The total data storage allocated, in terabytes (TB).
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - User defined description of the cloud Autonomous VM cluster.
* `display_name` - The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
* `domain` - The domain name for the cloud Autonomous VM cluster.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for the cloud Autonomous VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per each OCPU core.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `node_count` - The number of database servers in the cloud VM cluster. 
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `ocpu_count` - The number of CPU cores enabled on the cloud Autonomous VM cluster. Only 1 decimal place is allowed for the fractional part.
* `reclaimable_cpus` - CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
* `shape` - The model name of the Exadata hardware running the cloud Autonomous VM cluster. 
* `state` - The current state of the cloud Autonomous VM cluster.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.

	**Subnet Restrictions:**
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time that the cloud Autonomous VM cluster was created.
* `time_updated` - The last date and time that the cloud Autonomous VM cluster was updated.
* `total_container_databases` - The total number of Autonomous Container Databases that can be created with the allocated local storage.

