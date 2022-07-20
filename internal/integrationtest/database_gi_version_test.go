// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseGiVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `ExadataCC.Quarter3.100`},
	}

	DatabaseGiVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseGiVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseGiVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_gi_versions.test_gi_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_gi_versions", "test_gi_versions", acctest.Required, acctest.Create, DatabaseDatabaseGiVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseGiVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "gi_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "gi_versions.0.version"),
			),
		},
	})
}
