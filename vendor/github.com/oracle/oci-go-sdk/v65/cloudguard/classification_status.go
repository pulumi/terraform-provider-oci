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
	"strings"
)

// ClassificationStatusEnum Enum with underlying type: string
type ClassificationStatusEnum string

// Set of constants representing the allowable values for ClassificationStatusEnum
const (
	ClassificationStatusFalsePositive ClassificationStatusEnum = "FALSE_POSITIVE"
	ClassificationStatusFalseNegative ClassificationStatusEnum = "FALSE_NEGATIVE"
	ClassificationStatusTruePositive  ClassificationStatusEnum = "TRUE_POSITIVE"
	ClassificationStatusTrueNegative  ClassificationStatusEnum = "TRUE_NEGATIVE"
	ClassificationStatusNotClassified ClassificationStatusEnum = "NOT_CLASSIFIED"
)

var mappingClassificationStatusEnum = map[string]ClassificationStatusEnum{
	"FALSE_POSITIVE": ClassificationStatusFalsePositive,
	"FALSE_NEGATIVE": ClassificationStatusFalseNegative,
	"TRUE_POSITIVE":  ClassificationStatusTruePositive,
	"TRUE_NEGATIVE":  ClassificationStatusTrueNegative,
	"NOT_CLASSIFIED": ClassificationStatusNotClassified,
}

var mappingClassificationStatusEnumLowerCase = map[string]ClassificationStatusEnum{
	"false_positive": ClassificationStatusFalsePositive,
	"false_negative": ClassificationStatusFalseNegative,
	"true_positive":  ClassificationStatusTruePositive,
	"true_negative":  ClassificationStatusTrueNegative,
	"not_classified": ClassificationStatusNotClassified,
}

// GetClassificationStatusEnumValues Enumerates the set of values for ClassificationStatusEnum
func GetClassificationStatusEnumValues() []ClassificationStatusEnum {
	values := make([]ClassificationStatusEnum, 0)
	for _, v := range mappingClassificationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetClassificationStatusEnumStringValues Enumerates the set of values in String for ClassificationStatusEnum
func GetClassificationStatusEnumStringValues() []string {
	return []string{
		"FALSE_POSITIVE",
		"FALSE_NEGATIVE",
		"TRUE_POSITIVE",
		"TRUE_NEGATIVE",
		"NOT_CLASSIFIED",
	}
}

// GetMappingClassificationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassificationStatusEnum(val string) (ClassificationStatusEnum, bool) {
	enum, ok := mappingClassificationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
