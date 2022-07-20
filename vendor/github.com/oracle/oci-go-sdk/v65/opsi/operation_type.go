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
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeEnableDatabaseInsight                   OperationTypeEnum = "ENABLE_DATABASE_INSIGHT"
	OperationTypeDisableDatabaseInsight                  OperationTypeEnum = "DISABLE_DATABASE_INSIGHT"
	OperationTypeUpdateDatabaseInsight                   OperationTypeEnum = "UPDATE_DATABASE_INSIGHT"
	OperationTypeCreateDatabaseInsight                   OperationTypeEnum = "CREATE_DATABASE_INSIGHT"
	OperationTypeMoveDatabaseInsight                     OperationTypeEnum = "MOVE_DATABASE_INSIGHT"
	OperationTypeDeleteDatabaseInsight                   OperationTypeEnum = "DELETE_DATABASE_INSIGHT"
	OperationTypeCreateEnterpriseManagerBridge           OperationTypeEnum = "CREATE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeUdpateEnterpriseManagerBridge           OperationTypeEnum = "UDPATE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeMoveEnterpriseManagerBridge             OperationTypeEnum = "MOVE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeDeleteEnterpriseManagerBridge           OperationTypeEnum = "DELETE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeEnableHostInsight                       OperationTypeEnum = "ENABLE_HOST_INSIGHT"
	OperationTypeDisableHostInsight                      OperationTypeEnum = "DISABLE_HOST_INSIGHT"
	OperationTypeUpdateHostInsight                       OperationTypeEnum = "UPDATE_HOST_INSIGHT"
	OperationTypeCreateHostInsight                       OperationTypeEnum = "CREATE_HOST_INSIGHT"
	OperationTypeMoveHostInsight                         OperationTypeEnum = "MOVE_HOST_INSIGHT"
	OperationTypeDeleteHostInsight                       OperationTypeEnum = "DELETE_HOST_INSIGHT"
	OperationTypeCreateExadataInsight                    OperationTypeEnum = "CREATE_EXADATA_INSIGHT"
	OperationTypeEnableExadataInsight                    OperationTypeEnum = "ENABLE_EXADATA_INSIGHT"
	OperationTypeDisableExadataInsight                   OperationTypeEnum = "DISABLE_EXADATA_INSIGHT"
	OperationTypeUpdateExadataInsight                    OperationTypeEnum = "UPDATE_EXADATA_INSIGHT"
	OperationTypeMoveExadataInsight                      OperationTypeEnum = "MOVE_EXADATA_INSIGHT"
	OperationTypeDeleteExadataInsight                    OperationTypeEnum = "DELETE_EXADATA_INSIGHT"
	OperationTypeAddExadataInsightMembers                OperationTypeEnum = "ADD_EXADATA_INSIGHT_MEMBERS"
	OperationTypeExadataAutoSync                         OperationTypeEnum = "EXADATA_AUTO_SYNC"
	OperationTypeUpdateOpsiWarehouse                     OperationTypeEnum = "UPDATE_OPSI_WAREHOUSE"
	OperationTypeCreateOpsiWarehouse                     OperationTypeEnum = "CREATE_OPSI_WAREHOUSE"
	OperationTypeMoveOpsiWarehouse                       OperationTypeEnum = "MOVE_OPSI_WAREHOUSE"
	OperationTypeDeleteOpsiWarehouse                     OperationTypeEnum = "DELETE_OPSI_WAREHOUSE"
	OperationTypeRotateOpsiWarehouseWallet               OperationTypeEnum = "ROTATE_OPSI_WAREHOUSE_WALLET"
	OperationTypeUpdateOpsiWarehouseUser                 OperationTypeEnum = "UPDATE_OPSI_WAREHOUSE_USER"
	OperationTypeCreateOpsiWarehouseUser                 OperationTypeEnum = "CREATE_OPSI_WAREHOUSE_USER"
	OperationTypeMoveOpsiWarehouseUser                   OperationTypeEnum = "MOVE_OPSI_WAREHOUSE_USER"
	OperationTypeDeleteOpsiWarehouseUser                 OperationTypeEnum = "DELETE_OPSI_WAREHOUSE_USER"
	OperationTypeUpdateAwrhub                            OperationTypeEnum = "UPDATE_AWRHUB"
	OperationTypeCreateAwrhub                            OperationTypeEnum = "CREATE_AWRHUB"
	OperationTypeMoveAwrhub                              OperationTypeEnum = "MOVE_AWRHUB"
	OperationTypeDeleteAwrhub                            OperationTypeEnum = "DELETE_AWRHUB"
	OperationTypeUpdatePrivateEndpoint                   OperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	OperationTypeCreatePrivateEndpoint                   OperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	OperationTypeMovePrivateEndpoint                     OperationTypeEnum = "MOVE_PRIVATE_ENDPOINT"
	OperationTypeDeletePrivateEndpoint                   OperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	OperationTypeChangePeComanagedDatabaseInsightDetails OperationTypeEnum = "CHANGE_PE_COMANAGED_DATABASE_INSIGHT_DETAILS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"ENABLE_DATABASE_INSIGHT":                      OperationTypeEnableDatabaseInsight,
	"DISABLE_DATABASE_INSIGHT":                     OperationTypeDisableDatabaseInsight,
	"UPDATE_DATABASE_INSIGHT":                      OperationTypeUpdateDatabaseInsight,
	"CREATE_DATABASE_INSIGHT":                      OperationTypeCreateDatabaseInsight,
	"MOVE_DATABASE_INSIGHT":                        OperationTypeMoveDatabaseInsight,
	"DELETE_DATABASE_INSIGHT":                      OperationTypeDeleteDatabaseInsight,
	"CREATE_ENTERPRISE_MANAGER_BRIDGE":             OperationTypeCreateEnterpriseManagerBridge,
	"UDPATE_ENTERPRISE_MANAGER_BRIDGE":             OperationTypeUdpateEnterpriseManagerBridge,
	"MOVE_ENTERPRISE_MANAGER_BRIDGE":               OperationTypeMoveEnterpriseManagerBridge,
	"DELETE_ENTERPRISE_MANAGER_BRIDGE":             OperationTypeDeleteEnterpriseManagerBridge,
	"ENABLE_HOST_INSIGHT":                          OperationTypeEnableHostInsight,
	"DISABLE_HOST_INSIGHT":                         OperationTypeDisableHostInsight,
	"UPDATE_HOST_INSIGHT":                          OperationTypeUpdateHostInsight,
	"CREATE_HOST_INSIGHT":                          OperationTypeCreateHostInsight,
	"MOVE_HOST_INSIGHT":                            OperationTypeMoveHostInsight,
	"DELETE_HOST_INSIGHT":                          OperationTypeDeleteHostInsight,
	"CREATE_EXADATA_INSIGHT":                       OperationTypeCreateExadataInsight,
	"ENABLE_EXADATA_INSIGHT":                       OperationTypeEnableExadataInsight,
	"DISABLE_EXADATA_INSIGHT":                      OperationTypeDisableExadataInsight,
	"UPDATE_EXADATA_INSIGHT":                       OperationTypeUpdateExadataInsight,
	"MOVE_EXADATA_INSIGHT":                         OperationTypeMoveExadataInsight,
	"DELETE_EXADATA_INSIGHT":                       OperationTypeDeleteExadataInsight,
	"ADD_EXADATA_INSIGHT_MEMBERS":                  OperationTypeAddExadataInsightMembers,
	"EXADATA_AUTO_SYNC":                            OperationTypeExadataAutoSync,
	"UPDATE_OPSI_WAREHOUSE":                        OperationTypeUpdateOpsiWarehouse,
	"CREATE_OPSI_WAREHOUSE":                        OperationTypeCreateOpsiWarehouse,
	"MOVE_OPSI_WAREHOUSE":                          OperationTypeMoveOpsiWarehouse,
	"DELETE_OPSI_WAREHOUSE":                        OperationTypeDeleteOpsiWarehouse,
	"ROTATE_OPSI_WAREHOUSE_WALLET":                 OperationTypeRotateOpsiWarehouseWallet,
	"UPDATE_OPSI_WAREHOUSE_USER":                   OperationTypeUpdateOpsiWarehouseUser,
	"CREATE_OPSI_WAREHOUSE_USER":                   OperationTypeCreateOpsiWarehouseUser,
	"MOVE_OPSI_WAREHOUSE_USER":                     OperationTypeMoveOpsiWarehouseUser,
	"DELETE_OPSI_WAREHOUSE_USER":                   OperationTypeDeleteOpsiWarehouseUser,
	"UPDATE_AWRHUB":                                OperationTypeUpdateAwrhub,
	"CREATE_AWRHUB":                                OperationTypeCreateAwrhub,
	"MOVE_AWRHUB":                                  OperationTypeMoveAwrhub,
	"DELETE_AWRHUB":                                OperationTypeDeleteAwrhub,
	"UPDATE_PRIVATE_ENDPOINT":                      OperationTypeUpdatePrivateEndpoint,
	"CREATE_PRIVATE_ENDPOINT":                      OperationTypeCreatePrivateEndpoint,
	"MOVE_PRIVATE_ENDPOINT":                        OperationTypeMovePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                      OperationTypeDeletePrivateEndpoint,
	"CHANGE_PE_COMANAGED_DATABASE_INSIGHT_DETAILS": OperationTypeChangePeComanagedDatabaseInsightDetails,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"enable_database_insight":                      OperationTypeEnableDatabaseInsight,
	"disable_database_insight":                     OperationTypeDisableDatabaseInsight,
	"update_database_insight":                      OperationTypeUpdateDatabaseInsight,
	"create_database_insight":                      OperationTypeCreateDatabaseInsight,
	"move_database_insight":                        OperationTypeMoveDatabaseInsight,
	"delete_database_insight":                      OperationTypeDeleteDatabaseInsight,
	"create_enterprise_manager_bridge":             OperationTypeCreateEnterpriseManagerBridge,
	"udpate_enterprise_manager_bridge":             OperationTypeUdpateEnterpriseManagerBridge,
	"move_enterprise_manager_bridge":               OperationTypeMoveEnterpriseManagerBridge,
	"delete_enterprise_manager_bridge":             OperationTypeDeleteEnterpriseManagerBridge,
	"enable_host_insight":                          OperationTypeEnableHostInsight,
	"disable_host_insight":                         OperationTypeDisableHostInsight,
	"update_host_insight":                          OperationTypeUpdateHostInsight,
	"create_host_insight":                          OperationTypeCreateHostInsight,
	"move_host_insight":                            OperationTypeMoveHostInsight,
	"delete_host_insight":                          OperationTypeDeleteHostInsight,
	"create_exadata_insight":                       OperationTypeCreateExadataInsight,
	"enable_exadata_insight":                       OperationTypeEnableExadataInsight,
	"disable_exadata_insight":                      OperationTypeDisableExadataInsight,
	"update_exadata_insight":                       OperationTypeUpdateExadataInsight,
	"move_exadata_insight":                         OperationTypeMoveExadataInsight,
	"delete_exadata_insight":                       OperationTypeDeleteExadataInsight,
	"add_exadata_insight_members":                  OperationTypeAddExadataInsightMembers,
	"exadata_auto_sync":                            OperationTypeExadataAutoSync,
	"update_opsi_warehouse":                        OperationTypeUpdateOpsiWarehouse,
	"create_opsi_warehouse":                        OperationTypeCreateOpsiWarehouse,
	"move_opsi_warehouse":                          OperationTypeMoveOpsiWarehouse,
	"delete_opsi_warehouse":                        OperationTypeDeleteOpsiWarehouse,
	"rotate_opsi_warehouse_wallet":                 OperationTypeRotateOpsiWarehouseWallet,
	"update_opsi_warehouse_user":                   OperationTypeUpdateOpsiWarehouseUser,
	"create_opsi_warehouse_user":                   OperationTypeCreateOpsiWarehouseUser,
	"move_opsi_warehouse_user":                     OperationTypeMoveOpsiWarehouseUser,
	"delete_opsi_warehouse_user":                   OperationTypeDeleteOpsiWarehouseUser,
	"update_awrhub":                                OperationTypeUpdateAwrhub,
	"create_awrhub":                                OperationTypeCreateAwrhub,
	"move_awrhub":                                  OperationTypeMoveAwrhub,
	"delete_awrhub":                                OperationTypeDeleteAwrhub,
	"update_private_endpoint":                      OperationTypeUpdatePrivateEndpoint,
	"create_private_endpoint":                      OperationTypeCreatePrivateEndpoint,
	"move_private_endpoint":                        OperationTypeMovePrivateEndpoint,
	"delete_private_endpoint":                      OperationTypeDeletePrivateEndpoint,
	"change_pe_comanaged_database_insight_details": OperationTypeChangePeComanagedDatabaseInsightDetails,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"ENABLE_DATABASE_INSIGHT",
		"DISABLE_DATABASE_INSIGHT",
		"UPDATE_DATABASE_INSIGHT",
		"CREATE_DATABASE_INSIGHT",
		"MOVE_DATABASE_INSIGHT",
		"DELETE_DATABASE_INSIGHT",
		"CREATE_ENTERPRISE_MANAGER_BRIDGE",
		"UDPATE_ENTERPRISE_MANAGER_BRIDGE",
		"MOVE_ENTERPRISE_MANAGER_BRIDGE",
		"DELETE_ENTERPRISE_MANAGER_BRIDGE",
		"ENABLE_HOST_INSIGHT",
		"DISABLE_HOST_INSIGHT",
		"UPDATE_HOST_INSIGHT",
		"CREATE_HOST_INSIGHT",
		"MOVE_HOST_INSIGHT",
		"DELETE_HOST_INSIGHT",
		"CREATE_EXADATA_INSIGHT",
		"ENABLE_EXADATA_INSIGHT",
		"DISABLE_EXADATA_INSIGHT",
		"UPDATE_EXADATA_INSIGHT",
		"MOVE_EXADATA_INSIGHT",
		"DELETE_EXADATA_INSIGHT",
		"ADD_EXADATA_INSIGHT_MEMBERS",
		"EXADATA_AUTO_SYNC",
		"UPDATE_OPSI_WAREHOUSE",
		"CREATE_OPSI_WAREHOUSE",
		"MOVE_OPSI_WAREHOUSE",
		"DELETE_OPSI_WAREHOUSE",
		"ROTATE_OPSI_WAREHOUSE_WALLET",
		"UPDATE_OPSI_WAREHOUSE_USER",
		"CREATE_OPSI_WAREHOUSE_USER",
		"MOVE_OPSI_WAREHOUSE_USER",
		"DELETE_OPSI_WAREHOUSE_USER",
		"UPDATE_AWRHUB",
		"CREATE_AWRHUB",
		"MOVE_AWRHUB",
		"DELETE_AWRHUB",
		"UPDATE_PRIVATE_ENDPOINT",
		"CREATE_PRIVATE_ENDPOINT",
		"MOVE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"CHANGE_PE_COMANAGED_DATABASE_INSIGHT_DETAILS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
