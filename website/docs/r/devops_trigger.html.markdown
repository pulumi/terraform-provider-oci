---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_trigger"
sidebar_current: "docs-oci-resource-devops-trigger"
description: |-
  Provides the Trigger resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_trigger
This resource provides the Trigger resource in Oracle Cloud Infrastructure Devops service.

Creates a new trigger.


## Example Usage

```hcl
resource "oci_devops_trigger" "test_trigger" {
	#Required
	actions {
		#Required
		build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
		type = var.trigger_actions_type

		#Optional
		filter {
			#Required
			trigger_source = var.trigger_actions_filter_trigger_source

			#Optional
			events = var.trigger_actions_filter_events
			include {

				#Optional
				base_ref = var.trigger_actions_filter_include_base_ref
				head_ref = var.trigger_actions_filter_include_head_ref
			}
		}
	}
	project_id = oci_devops_project.test_project.id
	trigger_source = var.trigger_trigger_source

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.trigger_description
	display_name = var.trigger_display_name
	freeform_tags = {"bar-key"= "value"}
	repository_id = oci_artifacts_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `actions` - (Required) (Updatable) The list of actions that are to be performed for this trigger.
	* `build_pipeline_id` - (Required) (Updatable) The OCID of the build pipeline to be triggered.
	* `filter` - (Optional) (Updatable) The filters for the trigger.
		* `events` - (Optional) (Updatable) The events, for example, PUSH, PULL_REQUEST_MERGE.
		* `include` - (Optional) (Updatable) Attributes to filter DevOps code repository events.
			* `base_ref` - (Applicable when trigger_source=BITBUCKET_CLOUD | GITHUB | GITLAB) (Updatable) The target branch for pull requests; not applicable for push requests.
			* `head_ref` - (Optional) (Updatable) Branch for push event; source branch for pull requests.
		* `trigger_source` - (Required) (Updatable) Source of the trigger. Allowed values are, GITHUB, GITLAB and BITBUCKET_CLOUD.
	* `type` - (Required) (Updatable) The type of action that will be taken. Allowed value is TRIGGER_BUILD_PIPELINE.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) Optional description about the trigger.
* `display_name` - (Optional) (Updatable) Trigger display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `project_id` - (Required) The OCID of the DevOps project to which the trigger belongs to.
* `repository_id` - (Applicable when trigger_source=DEVOPS_CODE_REPOSITORY) (Updatable) The OCID of the DevOps code repository.
* `trigger_source` - (Required) (Updatable) Source of the trigger. Allowed values are, GITHUB,GITLAB and BITBUCKET_CLOUD.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actions` - The list of actions that are to be performed for this trigger.
	* `build_pipeline_id` - The OCID of the build pipeline to be triggered.
	* `filter` - The filters for the trigger.
		* `events` - The events, for example, PUSH, PULL_REQUEST_MERGE.
		* `include` - Attributes to filter DevOps code repository events.
			* `base_ref` - The target branch for pull requests; not applicable for push requests.
			* `head_ref` - Branch for push event; source branch for pull requests.
		* `trigger_source` - Source of the trigger. Allowed values are, GITHUB, GITLAB and BITBUCKET_CLOUD.
	* `type` - The type of action that will be taken. Allowed value is TRIGGER_BUILD_PIPELINE.
* `compartment_id` - The OCID of the compartment that contains the trigger.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Description about the trigger.
* `display_name` - Trigger display name. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of the DevOps project to which the trigger belongs to.
* `repository_id` - The OCID of the DevOps code repository.
* `state` - The current state of the trigger.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the trigger was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the trigger was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `trigger_source` - Source of the trigger. Allowed values are, GITHUB, GITLAB, BITBUCKET_CLOUD and DEVOPS_CODE_REPOSITORY.
* `trigger_url` - The endpoint that listens to trigger events.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Trigger
	* `update` - (Defaults to 20 minutes), when updating the Trigger
	* `delete` - (Defaults to 20 minutes), when destroying the Trigger


## Import

Triggers can be imported using the `id`, e.g.

```
$ terraform import oci_devops_trigger.test_trigger "id"
```

