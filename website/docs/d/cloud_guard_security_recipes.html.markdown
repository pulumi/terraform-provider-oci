---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_recipes"
sidebar_current: "docs-oci-datasource-cloud_guard-security_recipes"
description: |-
  Provides the list of Security Recipes in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_recipes
This data source provides the list of Security Recipes in Oracle Cloud Infrastructure Cloud Guard service.

Gets a list of all security zone recipes in a compartment.


## Example Usage

```hcl
data "oci_cloud_guard_security_recipes" "test_security_recipes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.security_recipe_display_name
	id = var.security_recipe_id
	state = var.security_recipe_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The unique identifier of the security zone recipe (`SecurityRecipe`)
* `state` - (Optional) The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `security_recipe_collection` - The list of security_recipe_collection.

### SecurityRecipe Reference

The following attributes are exported:

* `compartment_id` - The id of the compartment that contains the recipe
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The recipe's description
* `display_name` - The recipe's name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a recipe in the `Failed` state.
* `owner` - The owner of the recipe
* `security_policies` - The list of `SecurityPolicy` ids that are included in the recipe
* `state` - The current state of the recipe
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the recipe was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the recipe was last updated. An RFC3339 formatted datetime string.

