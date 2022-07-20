---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_connections"
sidebar_current: "docs-oci-datasource-devops-connections"
description: |-
  Provides the list of Connections in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_connections
This data source provides the list of Connections in Oracle Cloud Infrastructure Devops service.

Returns a list of connections.


## Example Usage

```hcl
data "oci_devops_connections" "test_connections" {

	#Optional
	compartment_id = var.compartment_id
	connection_type = var.connection_connection_type
	display_name = var.connection_display_name
	id = var.connection_id
	project_id = oci_devops_project.test_project.id
	state = var.connection_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `connection_type` - (Optional) A filter to return only resources that match the given connection type.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only connections that matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `connection_collection` - The list of connection_collection.

### Connection Reference

The following attributes are exported:

* `access_token` - The OCID of personal access token saved in secret store.
* `app_password` - OCID of personal Bitbucket Cloud AppPassword saved in secret store
* `compartment_id` - The OCID of the compartment containing the connection.
* `connection_type` - The type of connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Optional description about the connection.
* `display_name` - Connection display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `project_id` - The OCID of the DevOps project.
* `state` - The current state of the connection.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the connection was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the connection was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `username` - Public Bitbucket Cloud Username in plain text

