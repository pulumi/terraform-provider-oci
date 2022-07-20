// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EmdwControlPlane API
//
// A description of the EmdwControlPlane API
//

package emwarehouse

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmInstancesDetails Results of a emWarehouse search. Contains boh EmWarehouseSummary items and other information, such as metadata.
type EmInstancesDetails struct {

	// operations Insights Warehouse Identifier
	EmId *string `mandatory:"true" json:"emId"`

	// EmInstance Target count
	TargetsCount *int `mandatory:"false" json:"targetsCount"`

	// emHost name
	EmHost *string `mandatory:"false" json:"emHost"`
}

func (m EmInstancesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmInstancesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
