// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Monitor The information about a monitor.
type Monitor struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the monitor.
	Id *string `mandatory:"true" json:"id"`

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the monitor.
	MonitorType MonitorTypesEnum `mandatory:"true" json:"monitorType"`

	// List of public and dedicated vantage points where the monitor is running.
	VantagePoints []VantagePointInfo `mandatory:"true" json:"vantagePoints"`

	// Number of vantage points where monitor is running.
	VantagePointCount *int `mandatory:"true" json:"vantagePointCount"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the script.
	// scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null.
	ScriptId *string `mandatory:"true" json:"scriptId"`

	// Name of the script.
	ScriptName *string `mandatory:"true" json:"scriptName"`

	// Enables or disables the monitor.
	Status MonitorStatusEnum `mandatory:"true" json:"status"`

	// Interval in seconds after the start time when the job should be repeated.
	// Minimum repeatIntervalInSeconds should be 300 seconds for Scripted REST, Scripted Browser and Browser monitors, and 60 seconds for REST monitor.
	RepeatIntervalInSeconds *int `mandatory:"true" json:"repeatIntervalInSeconds"`

	// If runOnce is enabled, then the monitor will run once.
	IsRunOnce *bool `mandatory:"true" json:"isRunOnce"`

	// Timeout in seconds. Timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors.
	// Also, timeoutInSeconds should be a multiple of 60 for Scripted REST, Scripted Browser and Browser monitors.
	// Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that.
	TimeoutInSeconds *int `mandatory:"true" json:"timeoutInSeconds"`

	// Specify the endpoint on which to run the monitor.
	// For BROWSER and REST monitor types, target is mandatory.
	// If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint.
	// If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is.
	Target *string `mandatory:"false" json:"target"`

	// List of script parameters. Example: `[{"monitorScriptParameter": {"paramName": "userid", "paramValue":"testuser"}, "isSecret": false, "isOverwritten": false}]`
	ScriptParameters []MonitorScriptParameterInfo `mandatory:"false" json:"scriptParameters"`

	Configuration MonitorConfiguration `mandatory:"false" json:"configuration"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Monitor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Monitor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonitorTypesEnum(string(m.MonitorType)); !ok && m.MonitorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MonitorType: %s. Supported values are: %s.", m.MonitorType, strings.Join(GetMonitorTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMonitorStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMonitorStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Monitor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Target                  *string                           `json:"target"`
		ScriptParameters        []MonitorScriptParameterInfo      `json:"scriptParameters"`
		Configuration           monitorconfiguration              `json:"configuration"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Id                      *string                           `json:"id"`
		DisplayName             *string                           `json:"displayName"`
		MonitorType             MonitorTypesEnum                  `json:"monitorType"`
		VantagePoints           []VantagePointInfo                `json:"vantagePoints"`
		VantagePointCount       *int                              `json:"vantagePointCount"`
		ScriptId                *string                           `json:"scriptId"`
		ScriptName              *string                           `json:"scriptName"`
		Status                  MonitorStatusEnum                 `json:"status"`
		RepeatIntervalInSeconds *int                              `json:"repeatIntervalInSeconds"`
		IsRunOnce               *bool                             `json:"isRunOnce"`
		TimeoutInSeconds        *int                              `json:"timeoutInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Target = model.Target

	m.ScriptParameters = make([]MonitorScriptParameterInfo, len(model.ScriptParameters))
	for i, n := range model.ScriptParameters {
		m.ScriptParameters[i] = n
	}

	nn, e = model.Configuration.UnmarshalPolymorphicJSON(model.Configuration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Configuration = nn.(MonitorConfiguration)
	} else {
		m.Configuration = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.MonitorType = model.MonitorType

	m.VantagePoints = make([]VantagePointInfo, len(model.VantagePoints))
	for i, n := range model.VantagePoints {
		m.VantagePoints[i] = n
	}

	m.VantagePointCount = model.VantagePointCount

	m.ScriptId = model.ScriptId

	m.ScriptName = model.ScriptName

	m.Status = model.Status

	m.RepeatIntervalInSeconds = model.RepeatIntervalInSeconds

	m.IsRunOnce = model.IsRunOnce

	m.TimeoutInSeconds = model.TimeoutInSeconds

	return
}
