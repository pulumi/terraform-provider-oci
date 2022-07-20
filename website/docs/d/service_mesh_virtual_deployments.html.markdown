---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_deployments"
sidebar_current: "docs-oci-datasource-service_mesh-virtual_deployments"
description: |-
  Provides the list of Virtual Deployments in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_virtual_deployments
This data source provides the list of Virtual Deployments in Oracle Cloud Infrastructure Service Mesh service.

Returns a list of VirtualDeployments.


## Example Usage

```hcl
data "oci_service_mesh_virtual_deployments" "test_virtual_deployments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.virtual_deployment_id
	name = var.virtual_deployment_name
	state = var.virtual_deployment_state
	virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `id` - (Optional) Unique VirtualDeployment identifier.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources that match the life cycle state given.
* `virtual_service_id` - (Optional) Unique VirtualService identifier.


## Attributes Reference

The following attributes are exported:

* `virtual_deployment_collection` - The list of virtual_deployment_collection.

### VirtualDeployment Reference

The following attributes are exported:

* `access_logging` - This configuration determines if logging is enabled and where the logs will be output.
	* `is_enabled` - Determines if the logging configuration is enabled.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `listeners` - The listeners for the virtual deployment
	* `port` - Port in which virtual deployment is running.
	* `protocol` - Type of protocol used in virtual deployment.
* `name` - A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `service_discovery` - Service Discovery configuration for virtual deployments.
	* `hostname` - The hostname of the virtual deployments.
	* `type` - Type of service discovery.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.
* `virtual_service_id` - The OCID of the virtual service in which this virtual deployment is created.

