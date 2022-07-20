// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_em_warehouse.EmDataLakeClient", &OracleClient{InitClientFn: initEmwarehouseEmDataLakeClient})
}

func initEmwarehouseEmDataLakeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_em_warehouse.NewEmDataLakeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) EmDataLakeClient() *oci_em_warehouse.EmDataLakeClient {
	return m.GetClient("oci_em_warehouse.EmDataLakeClient").(*oci_em_warehouse.EmDataLakeClient)
}
