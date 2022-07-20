---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_esxi_hosts"
sidebar_current: "docs-oci-datasource-ocvp-esxi_hosts"
description: |-
  Provides the list of Esxi Hosts in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_esxi_hosts
This data source provides the list of Esxi Hosts in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists the ESXi hosts in the specified SDDC. The list can be filtered
by Compute instance OCID or ESXi display name.

Remember that in terms of implementation, an ESXi host is a Compute instance that
is configured with the chosen bundle of VMware software. Each `EsxiHost`
object has its own OCID (`id`), and a separate attribute for the OCID of
the Compute instance (`computeInstanceId`). When filtering the list of
ESXi hosts, you can specify the OCID of the Compute instance, not the
ESXi host OCID.


## Example Usage

```hcl
data "oci_ocvp_esxi_hosts" "test_esxi_hosts" {

	#Optional
	compute_instance_id = oci_core_instance.test_instance.id
	display_name = var.esxi_host_display_name
	sddc_id = oci_ocvp_sddc.test_sddc.id
	state = var.esxi_host_state
}
```

## Argument Reference

The following arguments are supported:

* `compute_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute instance.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `sddc_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `esxi_host_collection` - The list of esxi_host_collection.

### EsxiHost Reference

The following attributes are exported:

* `billing_contract_end_date` - Current billing cycle end date. If the value in `currentSku` and `nextSku` are different, the value specified in `nextSku` becomes the new `currentSKU` when the `contractEndDate` is reached. Example: `2016-08-25T21:10:29.600Z`
* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC.
* `compute_availability_domain` - The availability domain of the ESXi host.
* `compute_instance_id` - In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance.
* `current_sku` - The billing option currently used by the ESXi host. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A descriptive name for the ESXi host. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `failed_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is failed.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `grace_period_end_date` - The date and time when the new esxi host should start billing cycle. [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2021-07-25T21:10:29.600Z` 
* `host_ocpu_count` - The OCPU count of the ESXi host. 
* `host_shape_name` - The compute shape name of the ESXi host. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedHostShapes/ListSupportedHostShapes). 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host. 
* `next_sku` - The billing option to switch to after the current billing cycle ends. If `nextSku` is null or empty, `currentSku` continues to the next billing cycle. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `replacement_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is newly created to replace the failed node. 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to.
* `state` - The current state of the ESXi host.
* `time_created` - The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `time_updated` - The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
