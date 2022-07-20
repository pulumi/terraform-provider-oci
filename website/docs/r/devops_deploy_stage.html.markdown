---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_stage"
sidebar_current: "docs-oci-resource-devops-deploy_stage"
description: |-
  Provides the Deploy Stage resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_deploy_stage
This resource provides the Deploy Stage resource in Oracle Cloud Infrastructure Devops service.

Creates a new deployment stage.

## Example Usage

```hcl
resource "oci_devops_deploy_stage" "test_deploy_stage" {
	#Required
	deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
	deploy_stage_predecessor_collection {
		#Required
		items {
			#Required
			id = var.deploy_stage_deploy_stage_predecessor_collection_items_id
		}
	}
	deploy_stage_type = var.deploy_stage_deploy_stage_type

	#Optional
	approval_policy {
		#Required
		approval_policy_type = var.deploy_stage_approval_policy_approval_policy_type
		number_of_approvals_required = var.deploy_stage_approval_policy_number_of_approvals_required
	}
	blue_backend_ips {

		#Optional
		items = var.deploy_stage_blue_backend_ips_items
	}
	blue_green_strategy {
		#Required
		ingress_name = var.deploy_stage_blue_green_strategy_ingress_name
		namespace_a = var.deploy_stage_blue_green_strategy_namespace_a
		namespace_b = var.deploy_stage_blue_green_strategy_namespace_b
		strategy_type = var.deploy_stage_blue_green_strategy_strategy_type
	}
	canary_strategy {
		#Required
		ingress_name = var.deploy_stage_canary_strategy_ingress_name
		namespace = var.deploy_stage_canary_strategy_namespace
		strategy_type = var.deploy_stage_canary_strategy_strategy_type
	}
	compute_instance_group_blue_green_deployment_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_canary_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_canary_traffic_shift_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	config = var.deploy_stage_config
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	deploy_artifact_ids = var.deploy_stage_deploy_artifact_ids
	deploy_environment_id_a = var.deploy_stage_deploy_environment_id_a
	deploy_environment_id_b = var.deploy_stage_deploy_environment_id_b
	deployment_spec_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	description = var.deploy_stage_description
	display_name = var.deploy_stage_display_name
	docker_image_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	failure_policy {
		#Required
		policy_type = var.deploy_stage_failure_policy_policy_type

		#Optional
		failure_count = var.deploy_stage_failure_policy_failure_count
		failure_percentage = var.deploy_stage_failure_policy_failure_percentage
	}
	freeform_tags = {"bar-key"= "value"}
	function_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	function_timeout_in_seconds = var.deploy_stage_function_timeout_in_seconds
	green_backend_ips {

		#Optional
		items = var.deploy_stage_green_backend_ips_items
	}
	helm_chart_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	is_async = var.deploy_stage_is_async
	is_validation_enabled = var.deploy_stage_is_validation_enabled
	kubernetes_manifest_deploy_artifact_ids = var.deploy_stage_kubernetes_manifest_deploy_artifact_ids
	load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	max_memory_in_mbs = var.deploy_stage_max_memory_in_mbs
	namespace = var.deploy_stage_namespace
	oke_blue_green_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_canary_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_canary_traffic_shift_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_cluster_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	production_load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_production_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	release_name = var.deploy_stage_release_name
	rollback_policy {

		#Optional
		policy_type = var.deploy_stage_rollback_policy_policy_type
	}
	rollout_policy {
		#Required
		policy_type = var.deploy_stage_rollout_policy_policy_type

		#Optional
		batch_count = var.deploy_stage_rollout_policy_batch_count
		batch_delay_in_seconds = var.deploy_stage_rollout_policy_batch_delay_in_seconds
		batch_percentage = var.deploy_stage_rollout_policy_batch_percentage
		ramp_limit_percent = var.deploy_stage_rollout_policy_ramp_limit_percent
	}
	test_load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_test_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	timeout_in_seconds = var.deploy_stage_timeout_in_seconds
	traffic_shift_target = var.deploy_stage_traffic_shift_target
	values_artifact_ids = var.deploy_stage_values_artifact_ids
	wait_criteria {
		#Required
		wait_duration = var.deploy_stage_wait_criteria_wait_duration
		wait_type = var.deploy_stage_wait_criteria_wait_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `approval_policy` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL | MANUAL_APPROVAL | OKE_CANARY_APPROVAL) (Updatable) Specifies the approval policy.
	* `approval_policy_type` - (Required) (Updatable) Approval policy type.
	* `number_of_approvals_required` - (Required) (Updatable) A minimum number of approvals required for stage to proceed.
* `blue_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `blue_green_strategy` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT) Specifies the required blue green release strategy for OKE deployment.
	* `ingress_name` - (Required) Name of the Ingress resource.
	* `namespace_a` - (Required) First Namespace for deployment.
	* `namespace_b` - (Required) Second Namespace for deployment.
	* `strategy_type` - (Required) Blue Green strategy type
* `canary_strategy` - (Required when deploy_stage_type=OKE_CANARY_DEPLOYMENT) Specifies the required canary release strategy for OKE deployment.
	* `ingress_name` - (Required) Name of the Ingress resource.
	* `namespace` - (Required) Canary namespace to be used for Kubernetes canary deployment.
	* `strategy_type` - (Required) Canary strategy type
* `compute_instance_group_blue_green_deployment_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT) The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
* `compute_instance_group_canary_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT) A compute instance group canary stage OCID for load balancer.
* `compute_instance_group_canary_traffic_shift_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL) (Updatable) A compute instance group canary traffic shift stage OCID for load balancer.
* `compute_instance_group_deploy_environment_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) A compute instance group environment OCID for rolling deployment.
* `config` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - (Applicable when deploy_stage_type=INVOKE_FUNCTION) (Updatable) Optional binary artifact OCID user may provide to this stage.
* `deploy_artifact_ids` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) The list of file artifact OCIDs to deploy.
* `deploy_environment_id_a` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT) First compute instance group environment OCID for deployment.
* `deploy_environment_id_b` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT) Second compute instance group environment OCID for deployment.
* `deploy_pipeline_id` - (Required) The OCID of a pipeline.
* `deploy_stage_predecessor_collection` - (Required) (Updatable) Collection containing the predecessors of a stage.
	* `items` - (Required) (Updatable) A list of stage predecessors for a stage.
		* `id` - (Required) (Updatable) The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
* `deploy_stage_type` - (Required) (Updatable) Deployment stage type.
* `deployment_spec_deploy_artifact_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) The OCID of the artifact that contains the deployment specification.
* `description` - (Optional) (Updatable) Optional description about the deployment stage.
* `display_name` - (Optional) (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `docker_image_deploy_artifact_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) A Docker image artifact OCID.
* `failure_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	* `failure_count` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT) (Updatable) The threshold count of failed instances in the group, which when reached or exceeded sets the stage as FAILED.
	* `failure_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE) (Updatable) The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	* `policy_type` - (Required) (Updatable) Specifies if the failure instance size is given by absolute number or by percentage.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_deploy_environment_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION | INVOKE_FUNCTION) (Updatable) Function environment OCID.
* `function_timeout_in_seconds` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Timeout for execution of the Function. Value in seconds.
* `green_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `helm_chart_deploy_artifact_id` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Helm chart artifact OCID.
* `is_async` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether this stage executes asynchronously.
* `is_validation_enabled` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether the invoked function should be validated.
* `kubernetes_manifest_deploy_artifact_ids` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT | OKE_CANARY_DEPLOYMENT | OKE_DEPLOYMENT) (Updatable) List of Kubernetes manifest artifact OCIDs.
* `load_balancer_config` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The OCID of the load balancer.
* `max_memory_in_mbs` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Maximum usable memory for the Function (in MB).
* `namespace` - (Applicable when deploy_stage_type=OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_blue_green_deploy_stage_id` - (Required when deploy_stage_type=OKE_BLUE_GREEN_TRAFFIC_SHIFT) The OCID of the upstream OKE blue-green deployment stage in this pipeline.
* `oke_canary_deploy_stage_id` - (Required when deploy_stage_type=OKE_CANARY_TRAFFIC_SHIFT) The OCID of an upstream OKE canary deployment stage in this pipeline.
* `oke_canary_traffic_shift_deploy_stage_id` - (Required when deploy_stage_type=OKE_CANARY_APPROVAL) The OCID of an upstream OKE canary deployment traffic shift stage in this pipeline.
* `oke_cluster_deploy_environment_id` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT | OKE_CANARY_DEPLOYMENT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Kubernetes cluster environment OCID for deployment.
* `production_load_balancer_config` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Specifies configuration for load balancer traffic shift stages. The load balancer specified here should be an Application load balancer type. Network load balancers are not supported.
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) The OCID of the load balancer.
* `release_name` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Default name of the chart instance. Must be unique within a Kubernetes namespace.
* `rollback_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies type of the deployment stage rollback policy.
* `rollout_policy` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE) (Updatable) The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - (Required) (Updatable) The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) Indicates the criteria to stop.
* `test_load_balancer_config` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) The OCID of the load balancer.
* `timeout_in_seconds` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Time to wait for execution of a helm stage. Defaults to 300 seconds.
* `traffic_shift_target` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies the target or destination backend set.
* `values_artifact_ids` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) List of values.yaml file artifact OCIDs.
* `wait_criteria` - (Required when deploy_stage_type=WAIT) (Updatable) Specifies wait criteria for the Wait stage.
	* `wait_duration` - (Required) (Updatable) The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - (Required) (Updatable) Wait criteria type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approval_policy` - Specifies the approval policy.
	* `approval_policy_type` - Approval policy type.
	* `number_of_approvals_required` - A minimum number of approvals required for stage to proceed.
* `blue_backend_ips` - Collection of backend environment IP addresses.
	* `items` - The IP address of the backend server. A server could be a compute instance or a load balancer.
* `blue_green_strategy` - Specifies the required blue green release strategy for OKE deployment.
	* `ingress_name` - Name of the Ingress resource.
	* `namespace_a` - First Namespace for deployment.
	* `namespace_b` - Second Namespace for deployment.
	* `strategy_type` - Blue Green strategy type
* `canary_strategy` - Specifies the required canary release strategy for OKE deployment.
	* `ingress_name` - Name of the Ingress resource.
	* `namespace` - Canary namespace to be used for Kubernetes canary deployment.
	* `strategy_type` - Canary strategy type
* `compartment_id` - The OCID of a compartment.
* `compute_instance_group_blue_green_deployment_deploy_stage_id` - The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
* `compute_instance_group_canary_deploy_stage_id` - The OCID of an upstream compute instance group canary deployment stage ID in this pipeline.
* `compute_instance_group_canary_traffic_shift_deploy_stage_id` - A compute instance group canary traffic shift stage OCID for load balancer.
* `compute_instance_group_deploy_environment_id` - A compute instance group environment OCID for rolling deployment.
* `config` - User provided key and value pair configuration, which is assigned through constants or parameter.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - Optional binary artifact OCID user may provide to this stage.
* `deploy_artifact_ids` - The list of file artifact OCIDs to deploy.
* `deploy_environment_id_a` - First compute instance group environment OCID for deployment.
* `deploy_environment_id_b` - Second compute instance group environment OCID for deployment.
* `deploy_pipeline_id` - The OCID of a pipeline.
* `deploy_stage_predecessor_collection` - Collection containing the predecessors of a stage.
	* `items` - A list of stage predecessors for a stage.
		* `id` - The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
* `deploy_stage_type` - Deployment stage type.
* `deployment_spec_deploy_artifact_id` - The OCID of the artifact that contains the deployment specification.
* `description` - Optional description about the deployment stage.
* `display_name` - Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `docker_image_deploy_artifact_id` - A Docker image artifact OCID.
* `failure_policy` - Specifies a failure policy for a compute instance group rolling deployment stage.
	* `failure_count` - The threshold count of failed instances in the group, which when reached or exceeded sets the stage as FAILED.
	* `failure_percentage` - The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	* `policy_type` - Specifies if the failure instance size is given by absolute number or by percentage.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_deploy_environment_id` - Function environment OCID.
* `function_timeout_in_seconds` - Timeout for execution of the Function. Value in seconds.
* `green_backend_ips` - Collection of backend environment IP addresses.
	* `items` - The IP address of the backend server. A server could be a compute instance or a load balancer.
* `helm_chart_deploy_artifact_id` - Helm chart artifact OCID. 
* `id` - Unique identifier that is immutable on creation.
* `is_async` - A boolean flag specifies whether this stage executes asynchronously.
* `is_validation_enabled` - A boolean flag specifies whether the invoked function must be validated.
* `kubernetes_manifest_deploy_artifact_ids` - List of Kubernetes manifest artifact OCIDs.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `max_memory_in_mbs` - Maximum usable memory for the Function (in MB).
* `namespace` - Default Namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_blue_green_deploy_stage_id` - The OCID of the upstream OKE blue-green deployment stage in this pipeline.
* `oke_canary_deploy_stage_id` - The OCID of an upstream OKE canary deployment stage in this pipeline.
* `oke_canary_traffic_shift_deploy_stage_id` - The OCID of an upstream OKE canary deployment traffic shift stage in this pipeline.
* `oke_cluster_deploy_environment_id` - Kubernetes cluster environment OCID for deployment.
* `production_load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `project_id` - The OCID of a project.
* `release_name` - Release name of the Helm chart.
* `rollback_policy` - Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - Specifies type of the deployment stage rollback policy.
* `rollout_policy` - Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - Indicates the criteria to stop.
* `state` - The current state of the deployment stage.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `test_load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `time_created` - Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `timeout_in_seconds` - Time to wait for execution of a helm stage. Defaults to 300 seconds.
* `traffic_shift_target` - Specifies the target or destination backend set.
* `values_artifact_ids` - List of values.yaml file artifact OCIDs.
* `wait_criteria` - Specifies wait criteria for the Wait stage.
	* `wait_duration` - The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - Wait criteria type.

## Import

DeployStages can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deploy_stage.test_deploy_stage "id"
```

