---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_mask_rule"
sidebar_current: "docs-oci-resource-cloud_guard-data_mask_rule"
description: |-
  Provides the Data Mask Rule resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_data_mask_rule
This resource provides the Data Mask Rule resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a new Data Mask Rule Definition


## Example Usage

```hcl
resource "oci_cloud_guard_data_mask_rule" "test_data_mask_rule" {
	#Required
	compartment_id = var.compartment_id
	data_mask_categories = var.data_mask_rule_data_mask_categories
	display_name = var.data_mask_rule_display_name
	iam_group_id = oci_identity_group.test_group.id
	target_selected {
		#Required
		kind = var.data_mask_rule_target_selected_kind

		#Optional
		values = var.data_mask_rule_target_selected_values
	}

	#Optional
	data_mask_rule_status = var.data_mask_rule_data_mask_rule_status
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.data_mask_rule_description
	freeform_tags = {"bar-key"= "value"}
	state = var.data_mask_rule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier where the resource is created
* `data_mask_categories` - (Required) (Updatable) Data Mask Categories
* `data_mask_rule_status` - (Optional) (Updatable) The status of the dataMaskRule.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) The data mask rule description. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) Data mask rule name. 

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `iam_group_id` - (Required) (Updatable) IAM Group id associated with the data mask rule
* `state` - (Optional) The current state of the DataMaskRule.
* `target_selected` - (Required) (Updatable) Target Selection eg select ALL or select on basis of TargetResourceTypes or TargetIds.
	* `kind` - (Required) (Updatable) Target selection.
	* `values` - (Applicable when kind=TARGETIDS | TARGETTYPES) (Updatable) Types of Targets


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier where the resource is created.
* `data_mask_categories` - Data Mask Categories
* `data_mask_rule_status` - The status of the dataMaskRule.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The data mask rule description.
* `display_name` - Data Mask Rule Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `iam_group_id` - IAM Group id associated with the data mask rule
* `id` - Unique identifier that is immutable on creation
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the DataMaskRule.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_selected` - Target Selection eg select ALL or select on basis of TargetResourceTypes or TargetIds.
	* `kind` - Target selection.
	* `values` - Types of Targets
* `time_created` - The date and time the target was created. Format defined by RFC3339.
* `time_updated` - The date and time the target was updated. Format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Mask Rule
	* `update` - (Defaults to 20 minutes), when updating the Data Mask Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Data Mask Rule


## Import

DataMaskRules can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_data_mask_rule.test_data_mask_rule "id"
```

