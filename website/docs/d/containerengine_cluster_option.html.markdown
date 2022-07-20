---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_option"
sidebar_current: "docs-oci-datasource-containerengine-cluster_option"
description: |-
  Provides details about a specific Cluster Option in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_option
This data source provides details about a specific Cluster Option resource in Oracle Cloud Infrastructure Container Engine service.

Get options available for clusters.

## Example Usage

```hcl
data "oci_containerengine_cluster_option" "test_cluster_option" {
	#Required
	cluster_option_id = oci_containerengine_cluster_option.test_cluster_option.id

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_option_id` - (Required) The id of the option set to retrieve. Use "all" get all options, or use a cluster ID to get options specific to the provided cluster.
* `compartment_id` - (Optional) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `cluster_pod_network_options` - Available CNIs and network options for existing and new node pools of the cluster
	* `cni_type` - The CNI used by the node pools of this cluster
* `kubernetes_versions` - Available Kubernetes versions.

