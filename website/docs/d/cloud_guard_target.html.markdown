---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_target"
sidebar_current: "docs-oci-datasource-cloud_guard-target"
description: |-
  Provides details about a specific Target in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_target
This data source provides details about a specific Target resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a Target identified by targetId

## Example Usage

```hcl
data "oci_cloud_guard_target" "test_target" {
	#Required
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) OCID of target


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier where the resource is created
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The target description.
* `display_name` - Target display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that is immutable on creation.
* `inherited_by_compartments` - List of inherited compartments
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `recipe_count` - Total number of recipes attached to target
* `state` - The current state of the Target.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_details` - Details specific to the target type.
	* `security_zone_display_name` - The name of the security zone to associate this compartment with.
	* `security_zone_id` - The OCID of the security zone to associate this compartment with.
	* `target_resource_type` - Possible type of targets.
	* `target_security_zone_recipes` - The list of security zone recipes to associate this compartment with.
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
* `target_detector_recipes` - List of detector recipes associated with target
	* `compartment_id` - compartmentId of detector recipe
	* `description` - Detector recipe description.
	* `detector` - Type of detector
	* `detector_recipe_id` - Unique identifier for Detector Recipe of which this is an extension
	* `detector_rules` - List of detector rules for the detector type for recipe - user input
		* `description` - Description for TargetDetectorRecipeDetectorRule. information.
		* `details` - Overriden settings of a Detector Rule applied on target
			* `condition_groups` - Condition group corresponding to each compartment
				* `compartment_id` - compartment associated with condition
				* `condition` - 
			* `configurations` - Configuration details
				* `config_key` - Unique name of the configuration
				* `data_type` - configuration data type
				* `name` - configuration name
				* `value` - configuration value
				* `values` - List of configuration values
					* `list_type` - configuration list item type, either CUSTOM or MANAGED
					* `managed_list_type` - type of the managed list
					* `value` - configuration value
			* `is_configuration_allowed` - configuration allowed or not
			* `is_enabled` - Enables the control
			* `labels` - user defined labels for a detector rule
			* `risk_level` - The Risk Level
		* `detector` - detector for the rule
		* `detector_rule_id` - The unique identifier of the detector rule.
		* `display_name` - Display name for TargetDetectorRecipeDetectorRule. information.
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
		* `managed_list_types` - List of cloudguard managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule
		* `resource_type` - resource type of the configuration to which the rule is applied
		* `service_type` - service type of the configuration to which the rule is applied
		* `state` - The current state of the DetectorRule.
		* `time_created` - The date and time the target detector recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target detector recipe rule was updated. Format defined by RFC3339.
	* `display_name` - Display name of detector recipe.
	* `effective_detector_rules` - List of effective detector rules for the detector type for recipe after applying defaults
		* `description` - Description for TargetDetectorRecipeDetectorRule. information.
		* `details` - Overriden settings of a Detector Rule applied on target
			* `condition_groups` - Condition group corresponding to each compartment
				* `compartment_id` - compartment associated with condition
				* `condition` - 
			* `configurations` - Configuration details
				* `config_key` - Unique name of the configuration
				* `data_type` - configuration data type
				* `name` - configuration name
				* `value` - configuration value
				* `values` - List of configuration values
					* `list_type` - configuration list item type, either CUSTOM or MANAGED
					* `managed_list_type` - type of the managed list
					* `value` - configuration value
			* `is_configuration_allowed` - configuration allowed or not
			* `is_enabled` - Enables the control
			* `labels` - user defined labels for a detector rule
			* `risk_level` - The Risk Level
		* `detector` - detector for the rule
		* `detector_rule_id` - The unique identifier of the detector rule.
		* `display_name` - Display name for TargetDetectorRecipeDetectorRule. information.
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
		* `managed_list_types` - List of cloudguard managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule
		* `resource_type` - resource type of the configuration to which the rule is applied
		* `service_type` - service type of the configuration to which the rule is applied
		* `state` - The current state of the DetectorRule.
		* `time_created` - The date and time the target detector recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target detector recipe rule was updated. Format defined by RFC3339.
	* `id` - Ocid for detector recipe
	* `owner` - Owner of detector recipe
	* `state` - The current state of the resource.
	* `time_created` - The date and time the target detector recipe was created. Format defined by RFC3339.
	* `time_updated` - The date and time the target detector recipe was updated. Format defined by RFC3339.
* `target_resource_id` - Resource ID which the target uses to monitor
* `target_resource_type` - possible type of targets
* `target_responder_recipes` - List of responder recipes associated with target
	* `compartment_id` - Compartment Identifier
	* `description` - ResponderRecipe description.
	* `display_name` - ResponderRecipe display name.
	* `effective_responder_rules` - List of responder rules associated with the recipe after applying all defaults
		* `compartment_id` - Compartment Identifier
		* `description` - ResponderRule description.
		* `details` - Details of ResponderRule.
			* `condition` - 
			* `configurations` - ResponderRule configurations
				* `config_key` - Unique name of the configuration
				* `name` - configuration name
				* `value` - configuration value
			* `is_enabled` - Identifies state for ResponderRule
			* `mode` - Execution Mode for ResponderRule
		* `display_name` - ResponderRule display name.
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
		* `policies` - List of Policy
		* `responder_rule_id` - Unique ResponderRule identifier.
		* `state` - The current state of the ResponderRule.
		* `supported_modes` - Supported Execution Modes
		* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target responder recipe rule was updated. Format defined by RFC3339.
		* `type` - Type of Responder
	* `id` - Unique identifier of TargetResponderRecipe that can't be changed after creation.
	* `owner` - Owner of ResponderRecipe
	* `responder_recipe_id` - Unique identifier for Responder Recipe of which this is an extension.
	* `responder_rules` - List of responder rules associated with the recipe - user input
		* `compartment_id` - Compartment Identifier
		* `description` - ResponderRule description.
		* `details` - Details of ResponderRule.
			* `condition` - 
			* `configurations` - ResponderRule configurations
				* `config_key` - Unique name of the configuration
				* `name` - configuration name
				* `value` - configuration value
			* `is_enabled` - Identifies state for ResponderRule
			* `mode` - Execution Mode for ResponderRule
		* `display_name` - ResponderRule display name.
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
		* `policies` - List of Policy
		* `responder_rule_id` - Unique ResponderRule identifier.
		* `state` - The current state of the ResponderRule.
		* `supported_modes` - Supported Execution Modes
		* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target responder recipe rule was updated. Format defined by RFC3339.
		* `type` - Type of Responder
	* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the target responder recipe rule was updated. Format defined by RFC3339.
* `time_created` - The date and time the target was created. Format defined by RFC3339.
* `time_updated` - The date and time the target was updated. Format defined by RFC3339.

