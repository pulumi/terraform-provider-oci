---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_budgets"
sidebar_current: "docs-oci-datasource-budget-budgets"
description: |-
  Provides the list of Budgets in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_budgets
This data source provides the list of Budgets in Oracle Cloud Infrastructure Budget service.

Gets a list of budgets in a compartment.

By default, ListBudgets returns budgets of the 'COMPARTMENT' target type, and the budget records with only one target compartment OCID.

To list all budgets, set the targetType query parameter to ALL (for example: 'targetType=ALL').

Additional targetTypes would be available in future releases. Clients should ignore new targetTypes, 
or upgrade to the latest version of the client SDK to handle new targetTypes.


## Example Usage

```hcl
data "oci_budget_budgets" "test_budgets" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	display_name = var.budget_display_name
	state = var.budget_state
	target_type = var.budget_target_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. This does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) The current state of the resource to filter by.
* `target_type` - (Optional) The type of target to filter by:
	* ALL - List all budgets
	* COMPARTMENT - List all budgets with targetType == "COMPARTMENT"
	* TAG - List all budgets with targetType == "TAG" 


## Attributes Reference

The following attributes are exported:

* `budgets` - The list of budgets.

### Budget Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle.
* `alert_rule_count` - The total number of alert rules in the budget.
* `amount` - The amount of the budget, expressed in the currency of the customer's rate card. 
* `budget_processing_period_start_offset` - The number of days offset from the first day of the month, at which the budget processing period starts. In months that have fewer days than this value, processing will begin on the last day of that month. For example, for a value of 12, processing starts every month on the 12th at midnight.
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the budget.
* `display_name` - The display name of the budget. Avoid entering confidential information.
* `forecasted_spend` - The forecasted spend in currency by the end of the current budget cycle.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the budget.
* `processing_period_type` - The type of the budget processing period. Valid values are INVOICE and MONTH. 
* `reset_period` - The reset period for the budget. 
* `state` - The current state of the budget.
* `target_compartment_id` - This is DEPRECATED. For backwards compatability, the property is populated when the targetType is "COMPARTMENT", and targets contain the specific target compartment OCID. For all other scenarios, this property will be left empty. 
* `target_type` - The type of target on which the budget is applied. 
* `targets` - The list of targets on which the budget is applied. If the targetType is "COMPARTMENT", the targets contain the list of compartment OCIDs. If the targetType is "TAG", the targets contain the list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}". 
* `time_created` - The time that the budget was created.
* `time_spend_computed` - The time that the budget spend was last computed.
* `time_updated` - The time that the budget was updated.
* `version` - The version of the budget. Starts from 1 and increments by 1.

