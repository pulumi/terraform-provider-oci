---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_stage"
sidebar_current: "docs-oci-datasource-devops-deploy_stage"
description: |-
  Provides details about a specific Deploy Stage in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_stage
This data source provides details about a specific Deploy Stage resource in Oracle Cloud Infrastructure Devops service.

Retrieves a deployment stage by identifier.

## Example Usage

```hcl
data "oci_devops_deploy_stage" "test_deploy_stage" {
	#Required
	deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
}
```

## Argument Reference

The following arguments are supported:

* `deploy_stage_id` - (Required) Unique stage identifier.


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

