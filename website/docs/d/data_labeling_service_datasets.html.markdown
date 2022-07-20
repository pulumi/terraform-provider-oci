---
subcategory: "Data Labeling Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_labeling_service_datasets"
sidebar_current: "docs-oci-datasource-data_labeling_service-datasets"
description: |-
  Provides the list of Datasets in Oracle Cloud Infrastructure Data Labeling Service service
---

# Data Source: oci_data_labeling_service_datasets
This data source provides the list of Datasets in Oracle Cloud Infrastructure Data Labeling Service service.

Returns a list of Datasets.


## Example Usage

```hcl
data "oci_data_labeling_service_datasets" "test_datasets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	annotation_format = var.dataset_annotation_format
	display_name = var.dataset_display_name
	id = var.dataset_id
	state = var.dataset_state
}
```

## Argument Reference

The following arguments are supported:

* `annotation_format` - (Optional) A filter to return only resources that match the entire annotation format given.
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique Dataset OCID
* `state` - (Optional) A filter to return only resources whose lifecycleState matches this query param.


## Attributes Reference

The following attributes are exported:

* `dataset_collection` - The list of dataset_collection.

### Dataset Reference

The following attributes are exported:

* `annotation_format` - The annotation format name required for labeling records.
* `compartment_id` - The OCID of the compartment of the resource.
* `dataset_format_details` - It specifies how to process the data. Supported formats include DOCUMENT, IMAGE, and TEXT.
	* `format_type` - The format type. DOCUMENT format is for record contents that are PDFs or TIFFs. IMAGE format is for record contents that are JPEGs or PNGs. TEXT format is for record contents that are TXT files.
	* `text_file_type_metadata` - Metadata for files with text content.
		* `column_delimiter` - A column delimiter
		* `column_index` - The index of a selected column. This is a zero-based index.
		* `column_name` - The name of a selected column.
		* `escape_character` - An escape character.
		* `format_type` - It defines the format type of text files.
		* `line_delimiter` - A line delimiter.
* `dataset_source_details` - This allows the customer to specify the source of the dataset.
	* `bucket` - The object storage bucket that contains the dataset data source.
	* `namespace` - The namespace of the bucket that contains the dataset data source.
	* `prefix` - A common path prefix shared by the objects that make up the dataset. Except for the CSV file type, records are not generated for the objects whose names exactly match with the prefix.
	* `source_type` - The source type. OBJECT_STORAGE allows the user to describe where in object storage the dataset is.
* `defined_tags` - The defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - A user provided description of the dataset
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - The OCID of the Dataset.
* `initial_record_generation_configuration` - The initial generate records configuration. It generates records from the dataset's source.
* `label_set` - An ordered collection of labels that are unique by name. 
	* `items` - An ordered collection of labels that are unique by name.
		* `name` - An unique name for a label within its dataset.
* `labeling_instructions` - The labeling instructions for human labelers in rich text format
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in FAILED or NEEDS_ATTENTION state.
* `state` - The state of a dataset. CREATING - The dataset is being created.  It will transition to ACTIVE when it is ready for labeling. ACTIVE   - The dataset is ready for labeling. UPDATING - The dataset is being updated.  It and its related resources may be unavailable for other updates until it returns to ACTIVE. NEEDS_ATTENTION - A dataset updation operation has failed due to validation or other errors and needs attention. DELETING - The dataset and its related resources are being deleted. DELETED  - The dataset has been deleted and is no longer available. FAILED   - The dataset has failed due to validation or other errors. 
* `time_created` - The date and time the resource was created, in the timestamp format defined by RFC3339.
* `time_updated` - The date and time the resource was last updated, in the timestamp format defined by RFC3339.

