// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectTemplatizedQuery Information required in a structured template to form and execute query on a data object.
type DataObjectTemplatizedQuery struct {

	// List of items to be added into the SELECT clause of the query; items will be added with comma separation.
	SelectList []string `mandatory:"false" json:"selectList"`

	// List of items to be added into the WHERE clause of the query; items will be added with AND separation.
	// Item can contain a single condition or multiple conditions.
	// Single condition e.g:  "optimizer_mode='mode1'"
	// Multiple conditions e.g: (module='module1' OR module='module2')
	WhereConditionsList []string `mandatory:"false" json:"whereConditionsList"`

	// List of items to be added into the GROUP BY clause of the query; items will be added with comma separation.
	GroupByList []string `mandatory:"false" json:"groupByList"`

	// List of items to be added into the HAVING clause of the query; items will be added with AND separation.
	HavingConditionsList []string `mandatory:"false" json:"havingConditionsList"`

	// List of items to be added into the ORDER BY clause of the query; items will be added with comma separation.
	OrderByList []string `mandatory:"false" json:"orderByList"`

	TimeFilters *DataObjectQueryTimeFilters `mandatory:"false" json:"timeFilters"`
}

func (m DataObjectTemplatizedQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectTemplatizedQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectTemplatizedQuery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectTemplatizedQuery DataObjectTemplatizedQuery
	s := struct {
		DiscriminatorParam string `json:"queryType"`
		MarshalTypeDataObjectTemplatizedQuery
	}{
		"TEMPLATIZED_QUERY",
		(MarshalTypeDataObjectTemplatizedQuery)(m),
	}

	return json.Marshal(&s)
}
