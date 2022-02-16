// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Rule The configuration details of an Events rule. For more information, see
// Managing Rules for Events (https://docs.cloud.oracle.com/iaas/Content/Events/Task/managingrules.htm).
type Rule struct {

	// A string that describes the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	// Example: `"This rule sends a notification upon completion of DbaaS backup."`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the rule.
	LifecycleState RuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A filter that specifies the event that will trigger actions associated with this rule. A few
	// important things to remember about filters:
	// * Fields not mentioned in the condition are ignored. You can create a valid filter that matches
	// all events with two curly brackets: `{}`
	//   For more examples, see
	// Matching Events with Filters (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm).
	// * For a condition with fields to match an event, the event must contain all the field names
	// listed in the condition. Field names must appear in the condition with the same nesting
	// structure used in the event.
	//   For a list of reference events, see
	// Services that Produce Events (https://docs.cloud.oracle.com/iaas/Content/Events/Reference/eventsproducers.htm).
	// * Rules apply to events in the compartment in which you create them and any child compartments.
	// This means that a condition specified by a rule only matches events emitted from resources in
	// the compartment or any of its child compartments.
	// * Wildcard matching is supported with the asterisk (*) character.
	//   For examples of wildcard matching, see
	// Matching Events with Filters (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/filterevents.htm)
	// Example: `\"eventType\": \"com.oraclecloud.databaseservice.autonomous.database.backup.end\"`
	Condition *string `mandatory:"true" json:"condition"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not this rule is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	Actions *ActionList `mandatory:"true" json:"actions"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
	Id *string `mandatory:"true" json:"id"`

	// The time this rule was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A string that describes the details of the rule. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A message generated by the Events service about the current state of this rule.
	LifecycleMessage *string `mandatory:"false" json:"lifecycleMessage"`
}

func (m Rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingRuleLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRuleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleLifecycleStateEnum Enum with underlying type: string
type RuleLifecycleStateEnum string

// Set of constants representing the allowable values for RuleLifecycleStateEnum
const (
	RuleLifecycleStateCreating RuleLifecycleStateEnum = "CREATING"
	RuleLifecycleStateActive   RuleLifecycleStateEnum = "ACTIVE"
	RuleLifecycleStateInactive RuleLifecycleStateEnum = "INACTIVE"
	RuleLifecycleStateUpdating RuleLifecycleStateEnum = "UPDATING"
	RuleLifecycleStateDeleting RuleLifecycleStateEnum = "DELETING"
	RuleLifecycleStateDeleted  RuleLifecycleStateEnum = "DELETED"
	RuleLifecycleStateFailed   RuleLifecycleStateEnum = "FAILED"
)

var mappingRuleLifecycleStateEnum = map[string]RuleLifecycleStateEnum{
	"CREATING": RuleLifecycleStateCreating,
	"ACTIVE":   RuleLifecycleStateActive,
	"INACTIVE": RuleLifecycleStateInactive,
	"UPDATING": RuleLifecycleStateUpdating,
	"DELETING": RuleLifecycleStateDeleting,
	"DELETED":  RuleLifecycleStateDeleted,
	"FAILED":   RuleLifecycleStateFailed,
}

// GetRuleLifecycleStateEnumValues Enumerates the set of values for RuleLifecycleStateEnum
func GetRuleLifecycleStateEnumValues() []RuleLifecycleStateEnum {
	values := make([]RuleLifecycleStateEnum, 0)
	for _, v := range mappingRuleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleLifecycleStateEnumStringValues Enumerates the set of values in String for RuleLifecycleStateEnum
func GetRuleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}
