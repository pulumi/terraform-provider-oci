// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations in Cloud Guard from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImpactedResourceSummary Impacted Resource summary Definition.
type ImpactedResourceSummary struct {

	// Unique identifier for finding event
	Id *string `mandatory:"true" json:"id"`

	// Unique id of the Impacted Resource
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Problem Id to which the Impacted Resource is associated
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Compartment Id where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the Impacted Resource
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// Type of the Impacted Resource
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Region where the resource is created
	Region *string `mandatory:"true" json:"region"`

	// Time when the problem was identified
	TimeIdentified *common.SDKTime `mandatory:"true" json:"timeIdentified"`

	// Identifier for the sighting type
	SightingType *string `mandatory:"false" json:"sightingType"`

	// Name of the sighting type
	SightingTypeDisplayName *string `mandatory:"false" json:"sightingTypeDisplayName"`
}

func (m ImpactedResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImpactedResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
