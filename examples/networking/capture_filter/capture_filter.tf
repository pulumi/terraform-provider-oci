// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_capture_filter" "example_capture_filter" {
  compartment_id = var.compartment_ocid
  display_name   = "exampleCaptureFilter"
  filter_type    = "VTAP"
  vtap_capture_filter_rules {
    traffic_direction = "INGRESS"
  }
}

