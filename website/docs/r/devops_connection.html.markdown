---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_connection"
sidebar_current: "docs-oci-resource-devops-connection"
description: |-
  Provides the Connection resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_connection
This resource provides the Connection resource in Oracle Cloud Infrastructure Devops service.

Creates a new connection.


## Example Usage

```hcl
resource "oci_devops_connection" "test_connection" {
	#Required
	connection_type = var.connection_connection_type
	project_id = oci_devops_project.test_project.id

	#Optional
	access_token = var.connection_access_token
	app_password = var.connection_app_password
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.connection_description
	display_name = var.connection_display_name
	freeform_tags = {"bar-key"= "value"}
	username = var.connection_username
}
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Required when connection_type=GITHUB_ACCESS_TOKEN | GITLAB_ACCESS_TOKEN) (Updatable) The OCID of personal access token saved in secret store.
* `app_password` - (Required when connection_type=BITBUCKET_CLOUD_APP_PASSWORD) (Updatable) OCID of personal Bitbucket Cloud AppPassword saved in secret store
* `connection_type` - (Required) (Updatable) The type of connection.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) Optional description about the connection.
* `display_name` - (Optional) (Updatable) Optional connection display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `project_id` - (Required) The OCID of the DevOps project.
* `username` - (Required when connection_type=BITBUCKET_CLOUD_APP_PASSWORD) (Updatable) Public Bitbucket Cloud Username in plain text(not more than 30 characters)


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Connection
	* `update` - (Defaults to 20 minutes), when updating the Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Connection


## Import

Connections can be imported using the `id`, e.g.

```
$ terraform import oci_devops_connection.test_connection "id"
```

