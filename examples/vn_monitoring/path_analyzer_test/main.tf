// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "path_analyzer_test_defined_tags_value" {
  default = "value"
}

variable "path_analyzer_test_destination_endpoint_address" {
  default = "172.16.1.51"
}

variable "path_analyzer_test_destination_endpoint_type" {
  default = "IP_ADDRESS"
}

variable "path_analyzer_test_display_name" {
  default = "Path Analyzer Test"
}

variable "path_analyzer_test_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "path_analyzer_test_protocol" {
  default = 1
}

variable "path_analyzer_test_protocol_parameters_destination_port" {
  default = 10
}

variable "path_analyzer_test_protocol_parameters_icmp_code" {
  default = 10
}

variable "path_analyzer_test_protocol_parameters_icmp_type" {
  default = 10
}

variable "path_analyzer_test_protocol_parameters_source_port" {
  default = 10
}

variable "path_analyzer_test_protocol_parameters_type" {
  default = "ICMP"
}

variable "path_analyzer_test_query_options_is_bi_directional_analysis" {
  default = false
}

variable "path_analyzer_test_source_endpoint_address" {
  default = "172.16.1.50"
}

variable "path_analyzer_test_source_endpoint_type" {
  default = "IP_ADDRESS"
}

variable "path_analyzer_test_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_vn_monitoring_path_analyzer_test" "test_path_analyzer_test" {
  #Required
  compartment_id = var.compartment_ocid
  destination_endpoint  {
    #Required
    type = var.path_analyzer_test_destination_endpoint_type

    #Optional
    address                  = var.path_analyzer_test_destination_endpoint_address

  }
  protocol = var.path_analyzer_test_protocol
  source_endpoint  {
    #Required
    type = var.path_analyzer_test_source_endpoint_type

    #Optional
    address                  = var.path_analyzer_test_source_endpoint_address
  }

  #Optional
  display_name  = var.path_analyzer_test_display_name
  freeform_tags = var.path_analyzer_test_freeform_tags
  protocol_parameters  {
    #Required
    type = var.path_analyzer_test_protocol_parameters_type

    #Optional
    destination_port = var.path_analyzer_test_protocol_parameters_destination_port
    icmp_code        = var.path_analyzer_test_protocol_parameters_icmp_code
    icmp_type        = var.path_analyzer_test_protocol_parameters_icmp_type
    source_port      = var.path_analyzer_test_protocol_parameters_source_port
  }
  query_options  {

    #Optional
    is_bi_directional_analysis = var.path_analyzer_test_query_options_is_bi_directional_analysis
  }
}

data "oci_vn_monitoring_path_analyzer_tests" "test_path_analyzer_tests" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.path_analyzer_test_display_name
  state        = var.path_analyzer_test_state
}
