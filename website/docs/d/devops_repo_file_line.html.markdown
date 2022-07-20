---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repo_file_line"
sidebar_current: "docs-oci-datasource-devops-repo_file_line"
description: |-
  Provides details about a specific Repo File Line in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repo_file_line
This data source provides details about a specific Repo File Line resource in Oracle Cloud Infrastructure Devops service.

Retrieve lines of a specified file. Supports starting line number and limit.


## Example Usage

```hcl
data "oci_devops_repo_file_line" "test_repo_file_line" {
	#Required
	repository_id = oci_devops_repository.test_repository.id
	revision = var.repo_file_line_revision
	file_path = var.repo_file_line_file_path

	#Optional
	start_line_number = var.repo_file_line_start_line_number
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) (Required) A filter to return file contents of the specified paths.
* `repository_id` - (Required) Unique repository identifier.
* `revision` - (Required) Retrieve file lines from specific revision.
* `start_line_number` - (Optional) Line number from where to start returning file lines.


## Attributes Reference

The following attributes are exported:

* `lines` - The list of lines in the file.
	* `line_content` - The content of the line.
	* `line_number` - The line number.

