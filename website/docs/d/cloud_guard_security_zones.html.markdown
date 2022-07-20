---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_zones"
sidebar_current: "docs-oci-datasource-cloud_guard-security_zones"
description: |-
  Provides the list of Security Zones in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_zones
This data source provides the list of Security Zones in Oracle Cloud Infrastructure Cloud Guard service.

Gets a list of all security zones in a compartment.


## Example Usage

```hcl
data "oci_cloud_guard_security_zones" "test_security_zones" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.security_zone_display_name
	id = var.security_zone_id
	is_required_security_zones_in_subtree = var.security_zone_is_required_security_zones_in_subtree
	security_recipe_id = oci_cloud_guard_security_recipe.test_security_recipe.id
	state = var.security_zone_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The unique identifier of the security zone (`SecurityZone`)
* `is_required_security_zones_in_subtree` - (Optional) security zones in the subtree
* `security_recipe_id` - (Optional) The unique identifier of the security zone recipe (`SecurityRecipe`)
* `state` - (Optional) The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `security_zone_collection` - The list of security_zone_collection.

### SecurityZone Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment for the security zone
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The security zone's description
* `display_name` - The security zone's name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that is immutable on creation
* `inherited_by_compartments` - List of inherited compartments
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a zone in the `Failed` state.
* `security_zone_recipe_id` - The OCID of the recipe (`SecurityRecipe`) for the security zone
* `security_zone_target_id` - The OCID of the target associated with the security zone
* `state` - The current state of the security zone
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the security zone was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the security zone was last updated. An RFC3339 formatted datetime string.

