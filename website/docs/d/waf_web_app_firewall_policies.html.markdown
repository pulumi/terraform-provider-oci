---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_web_app_firewall_policies"
sidebar_current: "docs-oci-datasource-waf-web_app_firewall_policies"
description: |-
  Provides the list of Web App Firewall Policies in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_web_app_firewall_policies
This data source provides the list of Web App Firewall Policies in Oracle Cloud Infrastructure Waf service.

Gets a list of all WebAppFirewallPolicies in a compartment.


## Example Usage

```hcl
data "oci_waf_web_app_firewall_policies" "test_web_app_firewall_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.web_app_firewall_policy_display_name
	id = var.web_app_firewall_policy_id
	state = var.web_app_firewall_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only the WebAppFirewallPolicy with the given [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `web_app_firewall_policy_collection` - The list of web_app_firewall_policy_collection.

### WebAppFirewallPolicy Reference

The following attributes are exported:

* `actions` - Predefined actions for use in multiple different rules. Not all actions are supported in every module. Some actions terminate further execution of modules and rules in a module and some do not. Actions names must be unique within this array. 
	* `body` - Type of returned HTTP response body.
		* `text` - Static response body text.
		* `type` - Type of HttpResponseBody.
	* `code` - Response code.

		The following response codes are valid values for this property:
		* 2xx

		200 OK 201 Created 202 Accepted 206 Partial Content
		* 3xx

		300 Multiple Choices 301 Moved Permanently 302 Found 303 See Other 307 Temporary Redirect
		* 4xx

		400 Bad Request 401 Unauthorized 403 Forbidden 404 Not Found 405 Method Not Allowed 408 Request Timeout 409 Conflict 411 Length Required 412 Precondition Failed 413 Payload Too Large 414 URI Too Long 415 Unsupported Media Type 416 Range Not Satisfiable 422 Unprocessable Entity 494 Request Header Too Large 495 Cert Error 496 No Cert 497 HTTP to HTTPS
		* 5xx

		500 Internal Server Error 501 Not Implemented 502 Bad Gateway 503 Service Unavailable 504 Gateway Timeout 507 Insufficient Storage

		Example: `200` 
	* `headers` - Adds headers defined in this array for HTTP response.

		Hop-by-hop headers are not allowed to be set:
		* Connection
		* Keep-Alive
		* Proxy-Authenticate
		* Proxy-Authorization
		* TE
		* Trailer
		* Transfer-Encoding
		* Upgrade 
		* `name` - The name of the header field.
		* `value` - The value of the header field.
	* `name` - Action name. Can be used to reference the action.
	* `type` - 
		* **CHECK** is a non-terminating action that does not stop the execution of rules in current module, just emits a log message documenting result of rule execution.
		* **ALLOW** is a non-terminating action which upon matching rule skips all remaining rules in the current module.
		* **RETURN_HTTP_RESPONSE** is a terminating action which is executed immediately, returns a defined HTTP response. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - WebAppFirewallPolicy display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebAppFirewallPolicy.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `request_access_control` - Module that allows inspection of HTTP request properties and to return a defined HTTP response. In this module, rules with the name 'Default Action' are not allowed, since this name is reserved for default action logs. 
	* `default_action_name` - References an default Action to take if no AccessControlRule was matched. Allowed action types:
		* **ALLOW** continues execution of other modules and their rules.
		* **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response. 
	* `rules` - Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `request_protection` - Module that allows to enable OCI-managed protection capabilities for incoming HTTP requests.
	* `body_inspection_size_limit_exceeded_action_name` - References action by name from actions defined in WebAppFirewallPolicy. Executed if HTTP message body size exceeds limit set in field `bodyInspectionSizeLimitInBytes`.

		If this field is `null` HTTP message body will inspected up to `bodyInspectionSizeLimitInBytes` and the rest will not be inspected by Protection Capabilities.

		Allowed action types:
		* **RETURN_HTTP_RESPONSE** terminates further execution of modules and rules and returns defined HTTP response. 
	* `body_inspection_size_limit_in_bytes` - Maximum size of inspected HTTP message body in bytes. Actions to take if this limit is exceeded are defined in `bodyInspectionSizeLimitExceededActionName`.

		Body inspection maximum size allowed is defined with per-tenancy limit: 8192 bytes.

		For steps to request a limit increase, see [Requesting a Service Limit Increase](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). 
	* `rules` - Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection Capabilities of REQUEST_PROTECTION_CAPABILITY type. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `is_body_inspection_enabled` - Enables/disables body inspection for this protection rule. Only Protection Rules in RequestProtection can have this option enabled. Response body inspection will be available at a later date. 
		* `name` - Rule name. Must be unique within the module.
		* `protection_capabilities` - An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are not necessarily executed in order of appearance. Their execution order is decided at runtime for improved performance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - Unique key of collaborative capability for which weight will be overridden.
				* `weight` - The value of weight to set.
			* `exclusions` - Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - Unique key of referenced protection capability.
			* `version` - Version of referenced protection capability.
		* `protection_capability_settings` - Settings for protection capabilities 
			* `allowed_http_methods` - List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - Type of WebAppFirewallPolicyRule.
* `request_rate_limiting` - Module that allows inspection of HTTP connection properties and to limit requests frequency for a given key.
	* `rules` - Ordered list of RequestRateLimitingRules. Rules are executed in order of appearance in this array. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `configurations` - Rate Limiting Configurations. Each configuration counts requests towards its own `requestsLimit`. 
			* `action_duration_in_seconds` - Duration of block action application in seconds when `requestsLimit` is reached. Optional and can be 0 (no block duration).
			* `period_in_seconds` - Evaluation period in seconds.
			* `requests_limit` - Requests allowed per evaluation period.
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `response_access_control` - Module that allows inspection of HTTP response properties and to return a defined HTTP response.
	* `rules` - Ordered list of AccessControlRules. Rules are executed in order of appearance in this array.
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `name` - Rule name. Must be unique within the module.
		* `type` - Type of WebAppFirewallPolicyRule.
* `response_protection` - Module that allows to enable OCI-managed protection capabilities for HTTP responses.
	* `rules` - Ordered list of ProtectionRules. Rules are executed in order of appearance in this array. ProtectionRules in this array can only use protection capabilities of RESPONSE_PROTECTION_CAPABILITY type. 
		* `action_name` - References action by name from actions defined in WebAppFirewallPolicy.
		* `condition` - An expression that determines whether or not the rule action should be executed. 
		* `condition_language` - The language used to parse condition from field `condition`. Available languages:
			* **JMESPATH** an extended JMESPath language syntax. 
		* `is_body_inspection_enabled` - Enables/disables body inspection for this protection rule. Only Protection Rules in RequestProtection can have this option enabled. Response body inspection will be available at a later date. 
		* `name` - Rule name. Must be unique within the module.
		* `protection_capabilities` - An ordered list that references OCI-managed protection capabilities. Referenced protection capabilities are not necessarily executed in order of appearance. Their execution order is decided at runtime for improved performance. The array cannot contain entries with the same pair of capability key and version more than once. 
			* `action_name` - Override action to take if capability was triggered, defined in Protection Rule for this capability. Only actions of type CHECK are allowed. 
			* `collaborative_action_threshold` - The minimum sum of weights of associated collaborative protection capabilities that have triggered which must be reached in order for _this_ capability to trigger. This field is ignored for non-collaborative capabilities. 
			* `collaborative_weights` - Explicit weight values to use for associated collaborative protection capabilities. 
				* `key` - Unique key of collaborative capability for which weight will be overridden.
				* `weight` - The value of weight to set.
			* `exclusions` - Identifies specific HTTP message parameters to exclude from inspection by a protection capability. 
				* `args` - List of URL query parameter values from form-urlencoded XML, JSON, AMP, or POST payloads to exclude from inspecting. Example: If we have query parameter 'argumentName=argumentValue' and args=['argumentName'], both 'argumentName' and 'argumentValue' will not be inspected. 
				* `request_cookies` - List of HTTP request cookie values (by cookie name) to exclude from inspecting. Example: If we have cookie 'cookieName=cookieValue' and requestCookies=['cookieName'], both 'cookieName' and 'cookieValue' will not be inspected. 
			* `key` - Unique key of referenced protection capability.
			* `version` - Version of referenced protection capability.
		* `protection_capability_settings` - Settings for protection capabilities 
			* `allowed_http_methods` - List of allowed HTTP methods. Each value as a RFC7230 formated token string. Used in protection capability 911100: Restrict HTTP Request Methods. 
			* `max_http_request_header_length` - Maximum allowed length of headers in an HTTP request. Used in protection capability: 9200024: Limit length of request header size. 
			* `max_http_request_headers` - Maximum number of headers allowed in an HTTP request. Used in protection capability 9200014: Limit Number of Request Headers. 
			* `max_number_of_arguments` - Maximum number of arguments allowed. Used in protection capability 920380: Number of Arguments Limits. 
			* `max_single_argument_length` - Maximum allowed length of a single argument. Used in protection capability 920370: Limit argument value length. 
			* `max_total_argument_length` - Maximum allowed total length of all arguments. Used in protection capability 920390: Limit arguments total length. 
		* `type` - Type of WebAppFirewallPolicyRule.
* `state` - The current state of the WebAppFirewallPolicy.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the WebAppFirewallPolicy was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the WebAppFirewallPolicy was updated. An RFC3339 formatted datetime string.

