// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityShapeFromTable The table data entity details.
type EntityShapeFromTable struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The shape ID.
	ShapeId *string `mandatory:"false" json:"shapeId"`

	// Specifies other type label.
	OtherTypeLabel *string `mandatory:"false" json:"otherTypeLabel"`

	// An array of unique keys.
	UniqueKeys []UniqueKey `mandatory:"false" json:"uniqueKeys"`

	// An array of foreign keys.
	ForeignKeys []ForeignKey `mandatory:"false" json:"foreignKeys"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The entity type.
	EntityType EntityShapeFromTableEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

//GetMetadata returns Metadata
func (m EntityShapeFromTable) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m EntityShapeFromTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityShapeFromTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEntityShapeFromTableEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetEntityShapeFromTableEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EntityShapeFromTable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEntityShapeFromTable EntityShapeFromTable
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeEntityShapeFromTable
	}{
		"TABLE_ENTITY",
		(MarshalTypeEntityShapeFromTable)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *EntityShapeFromTable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Metadata       *ObjectMetadata                    `json:"metadata"`
		Key            *string                            `json:"key"`
		ModelVersion   *string                            `json:"modelVersion"`
		ParentRef      *ParentReference                   `json:"parentRef"`
		Name           *string                            `json:"name"`
		Description    *string                            `json:"description"`
		ObjectVersion  *int                               `json:"objectVersion"`
		ExternalKey    *string                            `json:"externalKey"`
		Shape          *Shape                             `json:"shape"`
		ShapeId        *string                            `json:"shapeId"`
		EntityType     EntityShapeFromTableEntityTypeEnum `json:"entityType"`
		OtherTypeLabel *string                            `json:"otherTypeLabel"`
		UniqueKeys     []uniquekey                        `json:"uniqueKeys"`
		ForeignKeys    []ForeignKey                       `json:"foreignKeys"`
		ResourceName   *string                            `json:"resourceName"`
		DataFormat     *DataFormat                        `json:"dataFormat"`
		ObjectStatus   *int                               `json:"objectStatus"`
		Identifier     *string                            `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Metadata = model.Metadata

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.Shape = model.Shape

	m.ShapeId = model.ShapeId

	m.EntityType = model.EntityType

	m.OtherTypeLabel = model.OtherTypeLabel

	m.UniqueKeys = make([]UniqueKey, len(model.UniqueKeys))
	for i, n := range model.UniqueKeys {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.UniqueKeys[i] = nn.(UniqueKey)
		} else {
			m.UniqueKeys[i] = nil
		}
	}

	m.ForeignKeys = make([]ForeignKey, len(model.ForeignKeys))
	for i, n := range model.ForeignKeys {
		m.ForeignKeys[i] = n
	}

	m.ResourceName = model.ResourceName

	m.DataFormat = model.DataFormat

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	return
}

// EntityShapeFromTableEntityTypeEnum Enum with underlying type: string
type EntityShapeFromTableEntityTypeEnum string

// Set of constants representing the allowable values for EntityShapeFromTableEntityTypeEnum
const (
	EntityShapeFromTableEntityTypeTable     EntityShapeFromTableEntityTypeEnum = "TABLE"
	EntityShapeFromTableEntityTypeView      EntityShapeFromTableEntityTypeEnum = "VIEW"
	EntityShapeFromTableEntityTypeFile      EntityShapeFromTableEntityTypeEnum = "FILE"
	EntityShapeFromTableEntityTypeSql       EntityShapeFromTableEntityTypeEnum = "SQL"
	EntityShapeFromTableEntityTypeDataStore EntityShapeFromTableEntityTypeEnum = "DATA_STORE"
)

var mappingEntityShapeFromTableEntityTypeEnum = map[string]EntityShapeFromTableEntityTypeEnum{
	"TABLE":      EntityShapeFromTableEntityTypeTable,
	"VIEW":       EntityShapeFromTableEntityTypeView,
	"FILE":       EntityShapeFromTableEntityTypeFile,
	"SQL":        EntityShapeFromTableEntityTypeSql,
	"DATA_STORE": EntityShapeFromTableEntityTypeDataStore,
}

var mappingEntityShapeFromTableEntityTypeEnumLowerCase = map[string]EntityShapeFromTableEntityTypeEnum{
	"table":      EntityShapeFromTableEntityTypeTable,
	"view":       EntityShapeFromTableEntityTypeView,
	"file":       EntityShapeFromTableEntityTypeFile,
	"sql":        EntityShapeFromTableEntityTypeSql,
	"data_store": EntityShapeFromTableEntityTypeDataStore,
}

// GetEntityShapeFromTableEntityTypeEnumValues Enumerates the set of values for EntityShapeFromTableEntityTypeEnum
func GetEntityShapeFromTableEntityTypeEnumValues() []EntityShapeFromTableEntityTypeEnum {
	values := make([]EntityShapeFromTableEntityTypeEnum, 0)
	for _, v := range mappingEntityShapeFromTableEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityShapeFromTableEntityTypeEnumStringValues Enumerates the set of values in String for EntityShapeFromTableEntityTypeEnum
func GetEntityShapeFromTableEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
		"DATA_STORE",
	}
}

// GetMappingEntityShapeFromTableEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityShapeFromTableEntityTypeEnum(val string) (EntityShapeFromTableEntityTypeEnum, bool) {
	enum, ok := mappingEntityShapeFromTableEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
