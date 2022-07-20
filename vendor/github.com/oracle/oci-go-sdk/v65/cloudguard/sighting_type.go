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

// SightingType Specific behavior that can trigger a Sighting
type SightingType struct {

	// The unique identifier of sighting type
	Id *string `mandatory:"false" json:"id"`

	// Name of the sighting type
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the sighting type
	Description *string `mandatory:"false" json:"description"`

	// Link of the sighting type
	MitreLink *string `mandatory:"false" json:"mitreLink"`

	// Mitre Att&ck tactic
	Tactic *string `mandatory:"false" json:"tactic"`

	// List of Mitre Att&ck Techniques
	Techniques []string `mandatory:"false" json:"techniques"`
}

func (m SightingType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SightingType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
