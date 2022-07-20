// Copyright (c) 2017, 2022, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscribed_service_id" {}
variable "subscription_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_onesubscription_subscribed_service" "test_subscribed_service" {
  #Required
  subscribed_service_id = var.subscribed_service_id
}

data "oci_onesubscription_subscribed_services" "test_subscribed_services" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id
}
