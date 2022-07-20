---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance"
sidebar_current: "docs-oci-resource-bds-bds_instance"
description: |-
  Provides the Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance
This resource provides the Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service.

Creates a new BDS instance.


## Example Usage

```hcl
resource "oci_bds_bds_instance" "test_bds_instance" {
	#Required
	cluster_admin_password = var.bds_instance_cluster_admin_password
	cluster_public_key = var.bds_instance_cluster_public_key
	cluster_version = var.bds_instance_cluster_version
	compartment_id = var.compartment_id
	display_name = var.bds_instance_display_name
	is_high_availability = var.bds_instance_is_high_availability
	is_secure = var.bds_instance_is_secure
	master_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	util_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	worker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	compute_only_worker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}

	#Optional
	bootstrap_script_url = var.bds_instance_bootstrap_script_url
	defined_tags = var.bds_instance_defined_tags
	freeform_tags = var.bds_instance_freeform_tags
	kerberos_realm_name = var.bds_instance_kerberos_realm_name
	network_config {

		#Optional
		cidr_block = var.bds_instance_network_config_cidr_block
		is_nat_gateway_required = var.bds_instance_network_config_is_nat_gateway_required
	}
}
```

## Argument Reference

The following arguments are supported:

* `bootstrap_script_url` - (Optional) (Updatable) Pre-authenticated URL of the script in Object Store that is downloaded and executed.
* `cluster_admin_password` - (Required) Base-64 encoded password for Cloudera Manager admin user
* `cluster_public_key` - (Required) The SSH public key used to authenticate the cluster connection.
* `cluster_version` - (Required) Version of the Hadoop distribution
* `compartment_id` - (Required) (Updatable) The OCID of the compartment
* `display_name` - (Required) (Updatable) Name of the BDS instance
* `is_cloud_sql_configured` -(Optional) (Updatable) Boolean flag specifying whether we configure Cloud SQL or not
* `cloud_sql_details` -(Optional) The information about added Cloud SQL capability
    * `block_volume_size_in_gbs` - (Required) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
    * `shape` - (Required) Shape of the node
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_high_availability` - (Required) Boolean flag specifying whether or not the cluster is HA
* `is_secure` - (Required) Boolean flag specifying whether or not the cluster should be setup as secure.
* `kerberos_realm_name` - (Optional) The user-defined kerberos realm name.
* `network_config` - (Optional) Additional configuration of customer's network.
    * `cidr_block` - (Required) The CIDR IP address block of the VCN.
    * `is_nat_gateway_required` - (Required) A boolean flag whether to configure a NAT gateway.
* `master_node` - (Required) The master node in the BDS instance
    * `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
    * `number_of_nodes` - (Required) The amount of master nodes should be created.
    * `shape` - (Required) Shape of the node
    * `subnet_id` - (Required) The OCID of the subnet in which the node should be created
    * `shape_config` - (Optional) The shape configuration requested for the node.
      * `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
      * `ocpus` - (Optional) The total number of OCPUs available to the node.
* `util_node` - (Required) The utility node in the BDS instance
    * `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
    * `number_of_nodes` - (Required) The amount of utility nodes should be created.
    * `shape` - (Required) Shape of the node
    * `subnet_id` - (Required) The OCID of the subnet in which the node should be created
    * `shape_config` - (Optional) The shape configuration requested for the node.
      * `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
      * `ocpus` - (Optional) The total number of OCPUs available to the node.
* `woker_node` - (Required) The worker node in the BDS instance
    * `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
    * `number_of_nodes` - (Required) The amount of worker nodes should be created, at least be 3.
    * `shape` - (Required) Shape of the node
    * `subnet_id` - (Required) The OCID of the subnet in which the node should be created
    * `shape_config` - (Optional) The shape configuration requested for the node.
      * `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
      * `ocpus` - (Optional) The total number of OCPUs available to the node.
* `compute_only_woker_node` - (Required) The worker node in the BDS instance
    * `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
    * `number_of_nodes` - (Required) The amount of worker nodes should be created
    * `shape` - (Required) Shape of the node
    * `subnet_id` - (Required) The OCID of the subnet in which the node should be created
    * `shape_config` - (Optional) The shape configuration requested for the node.
        * `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
        * `ocpus` - (Optional) The total number of OCPUs available to the node.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bootstrap_script_url` - pre-authenticated URL of the bootstrap script in Object Store that can be downloaded and executed.
* `cloud_sql_details` - The information about added Cloud SQL capability
	* `block_volume_size_in_gbs` - The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
	* `ip_address` - IP address of the Cloud SQL node
	* `is_kerberos_mapped_to_database_users` - Boolean flag specifying whether or not are Kerberos principals mapped to database users. 
	* `kerberos_details` - Details about Kerberos principals
		* `keytab_file` - Location of the keytab file
		* `principal_name` - Name of the Kerberos principal
	* `shape` - Shape of the node
* `cluster_details` - Specific info about a Hadoop cluster
    * `bd_cell_version` - Cloud SQL cell version
    * `bda_version` - BDA version installed in the cluster
    * `bdm_version` - Big Data Manager version installed in the cluster
    * `bds_version` - Big Data Service version installed in the cluster
    * `big_data_manager_url` - The URL of a Big Data Manager
    * `cloudera_manager_url` - The URL of a Cloudera Manager
    * `csql_cell_version` - Big Data SQL version
    * `db_version` - Query Server Database version
    * `hue_server_url` - The URL of a Hue Server
    * `jupyter_hub_url` - The URL of the Jupyterhub.
    * `os_version` - Oracle Linux version installed in the cluster
    * `time_created` - The time the cluster was created. An RFC3339 formatted datetime string
    * `time_refreshed` - The time the BDS instance was automatically, or manually refreshed. An RFC3339 formatted datetime string 
* `cluster_version` - Version of the Hadoop distribution
* `compartment_id` - The OCID of the compartment
* `created_by` - The user who created the BDS instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the BDS instance
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the BDS resource
* `is_cloud_sql_configured` - Boolean flag specifying whether we configure Cloud SQL or not
* `is_high_availability` - Boolean flag specifying whether or not the cluster is HA
* `is_secure` - Boolean flag specifying whether or not the cluster should be setup as secure.
* `network_config` - Additional configuration of customer's network.
	* `cidr_block` - The CIDR IP address block of the VCN.
	* `is_nat_gateway_required` - A boolean flag whether to configure a NAT gateway.
* `nodes` - The list of nodes in the BDS instance
    * `attached_block_volumes` - The list of block volumes attached to a given node.
        * `volume_attachment_id` - The OCID of the volume attachment.
        * `volume_size_in_gbs` - The size of the volume in GBs.
    * `availability_domain` - The name of the availability domain the node is running in
    * `display_name` - The name of the node
    * `fault_domain` - The name of the fault domain the node is running in
    * `hostname` - The fully-qualified hostname (FQDN) of the node
    * `image_id` - The OCID of the image from which the node was created
    * `instance_id` - The OCID of the underlying compute instance
    * `ip_address` - IP address of the node
    * `memory_in_gbs` - The total amount of memory available to the node, in gigabytes.
    * `node_type` - BDS instance node type
    * `ocpus` - The total number of OCPUs available to the node.
    * `shape` - Shape of the node
    * `ssh_fingerprint` - The fingerprint of the SSH key used for node access
    * `state` - The state of the node
    * `subnet_id` - The OCID of the subnet in which the node should be created
    * `time_created` - The time the node was created. An RFC3339 formatted datetime string
    * `time_updated` - The time the BDS instance was updated. An RFC3339 formatted datetime string
* `number_of_nodes` - Number of nodes that forming the cluster
* `state` - The state of the BDS instance
* `time_created` - The time the BDS instance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the BDS instance was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance


## Import

BdsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance.test_bds_instance "id"
```

