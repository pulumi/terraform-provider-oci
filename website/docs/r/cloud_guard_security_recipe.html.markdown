---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_recipe"
sidebar_current: "docs-oci-resource-cloud_guard-security_recipe"
description: |-
  Provides the Security Recipe resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_security_recipe
This resource provides the Security Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a security zone recipe. A security zone recipe is a collection of security zone policies.


## Example Usage

```hcl
resource "oci_cloud_guard_security_recipe" "test_security_recipe" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.security_recipe_display_name
	security_policies = var.security_recipe_security_policies

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.security_recipe_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment in which to create the recipe
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The recipe's description
* `display_name` - (Required) (Updatable) The recipe's name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `security_policies` - (Required) (Updatable) The list of `SecurityPolicy` ids to include in the recipe


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Recipe
	* `update` - (Defaults to 20 minutes), when updating the Security Recipe
	* `delete` - (Defaults to 20 minutes), when destroying the Security Recipe


## Import

SecurityRecipes can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_security_recipe.test_security_recipe "id"
```

