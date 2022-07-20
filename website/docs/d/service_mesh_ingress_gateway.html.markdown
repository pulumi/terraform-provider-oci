---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_ingress_gateway"
sidebar_current: "docs-oci-datasource-service_mesh-ingress_gateway"
description: |-
  Provides details about a specific Ingress Gateway in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_ingress_gateway
This data source provides details about a specific Ingress Gateway resource in Oracle Cloud Infrastructure Service Mesh service.

Gets an IngressGateway by identifier.

## Example Usage

```hcl
data "oci_service_mesh_ingress_gateway" "test_ingress_gateway" {
	#Required
	ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
}
```

## Argument Reference

The following arguments are supported:

* `ingress_gateway_id` - (Required) Unique IngressGateway identifier.


## Attributes Reference

The following attributes are exported:

* `access_logging` - This configuration determines if logging is enabled and where the logs will be output.
	* `is_enabled` - Determines if the logging configuration is enabled.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hosts` - Array of hostnames and their listener configuration that this gateway will bind to.
	* `hostnames` - Hostnames of the host. Applicable only for HTTP and TLS_PASSTHROUGH listeners. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com". 
	* `listeners` - The listeners for the ingress gateway.
		* `port` - Port on which ingress gateway is listening.
		* `protocol` - Type of protocol used.
		* `tls` - TLS enforcement config for the ingress listener.
			* `client_validation` - Resource representing the TLS configuration used for validating client certificates. 
				* `subject_alternate_names` - A list of alternate names to verify the subject identity in the certificate presented by the client. 
				* `trusted_ca_bundle` - Resource representing the CA bundle.
					* `ca_bundle_id` - The OCID of the CA Bundle resource.
					* `secret_name` - Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt 
					* `type` - Type of certificate.
			* `mode` - DISABLED: Connection can only be plaintext. PERMISSIVE: Connection can be either plaintext or TLS/mTLS. If the clientValidation.trustedCaBundle property is configured for the listener, mTLS is performed and the client's certificates are validated by the gateway. TLS: Connection can only be TLS.  MUTUAL_TLS: Connection can only be MTLS. 
			* `server_certificate` - Resource representing the location of the TLS certificate.
				* `certificate_id` - The OCID of the leaf certificate resource.
				* `secret_name` - Name of the secret. For Kubernetes this is the name of the Kubernetes secret of type tls. For other platforms the secrets must be mounted at: /etc/oci/secrets/${secretName}/tls.{key,crt} 
				* `type` - Type of certificate.
	* `name` - A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: `MyExampleHost` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `mesh_id` - The OCID of the service mesh in which this ingress gateway is created.
* `mtls` - Mutual TLS settings used when sending requests to virtual services within the mesh. 
	* `certificate_id` - The OCID of the certificate resource that will be used for mTLS authentication with other virtual services in the mesh. 
	* `maximum_validity` - The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days. 
* `name` - A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

