---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_users"
sidebar_current: "docs-oci-datasource-database_management-managed_database_users"
description: |-
  Provides the list of Managed Database Users in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_users
This data source provides the list of Managed Database Users in Oracle Cloud Infrastructure Database Management service.

Gets the list of users for the specified managedDatabaseId.

## Example Usage

```hcl
data "oci_database_management_managed_database_users" "test_managed_database_users" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name = var.managed_database_user_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.


## Attributes Reference

The following attributes are exported:

* `user_collection` - The list of user_collection.

### ManagedDatabaseUser Reference

The following attributes are exported:

* `all_shared` - In a sharded database, indicates whether the user is created with shard DDL enabled (YES) or not (NO).
* `authentication` - The authentication mechanism for the user.
* `common` - Indicates whether a given user is common(Y) or local(N).
* `consumer_group` - The initial resource consumer group for the User.
* `default_collation` - The default collation for the user schema.
* `default_tablespace` - The default tablespace for data.
* `editions_enabled` - Indicates whether editions have been enabled for the corresponding user (Y) or not (N).
* `external_name` - The external name of the user.
* `external_shared` - In a federated sharded database, indicates whether the user is an external shard user (YES) or not (NO).
* `implicit` - Indicates whether the user is a common user created by an implicit application (YES) or not (NO).
* `inherited` - Indicates whether the user definition is inherited from another container (YES) or not (NO).
* `local_temp_tablespace` - The default local temporary tablespace for the user.
* `name` - The name of the User.
* `oracle_maintained` - Indicates whether the user was created and is maintained by Oracle-supplied scripts (such as catalog.sql or catproc.sql).
* `password_versions` - The list of existing versions of the password hashes (also known as "verifiers") for the account.
* `profile` - The profile name of the user.
* `proxy_connect` - Indicates whether a user can connect directly (N) or whether the account can only be proxied (Y) by users who have proxy privileges for this account (that is, by users who have been granted the "connect through" privilege for this account). 
* `status` - The status of the user account.
* `temp_tablespace` - The name of the default tablespace for temporary tables or the name of a tablespace group.
* `time_created` - The date and time the user was created.
* `time_expiring` - The date and time of the expiration of the user account.
* `time_last_login` - The date and time of the last user login. This column is not populated when a user connects to the database with administrative privileges, that is, AS { SYSASM | SYSBACKUP | SYSDBA | SYSDG | SYSOPER | SYSRAC | SYSKM }. 
* `time_locked` - The date the account was locked, if the status of the account is LOCKED.
* `time_password_changed` - The date and time when the user password was last set. This column is populated only when the value of the AUTHENTICATION_TYPE column is PASSWORD. Otherwise, this column is null. 

