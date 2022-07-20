// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CrossConnectMapping For use with Oracle Cloud Infrastructure FastConnect. Each
// VirtualCircuit runs on one or
// more cross-connects or cross-connect groups. A `CrossConnectMapping`
// contains the properties for an individual cross-connect or cross-connect group
// associated with a given virtual circuit.
// The mapping includes information about the cross-connect or
// cross-connect group, the VLAN, and the BGP peering session.
// If you're a customer who is colocated with Oracle, that means you own both
// the virtual circuit and the physical connection it runs on (cross-connect or
// cross-connect group), so you specify all the information in the mapping. There's
// one exception: for a public virtual circuit, Oracle specifies the BGP IPv4
// addresses.
// If you're a provider, then you own the physical connection that the customer's
// virtual circuit runs on, so you contribute information about the cross-connect
// or cross-connect group and VLAN.
// Who specifies the BGP peering information in the case of customer connection via
// provider? If the BGP session goes from Oracle to the provider's edge router, then
// the provider also specifies the BGP peering information. If the BGP session instead
// goes from Oracle to the customer's edge router, then the customer specifies the BGP
// peering information. There's one exception: for a public virtual circuit, Oracle
// specifies the BGP IPv4 addresses.
type CrossConnectMapping struct {

	// The key for BGP MD5 authentication. Only applicable if your system
	// requires MD5 authentication. If empty or not set (null), that
	// means you don't use BGP MD5 authentication.
	BgpMd5AuthKey *string `mandatory:"false" json:"bgpMd5AuthKey"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect or cross-connect group for this mapping.
	// Specified by the owner of the cross-connect or cross-connect group (the
	// customer if the customer is colocated with Oracle, or the provider if the
	// customer is connecting via provider).
	CrossConnectOrCrossConnectGroupId *string `mandatory:"false" json:"crossConnectOrCrossConnectGroupId"`

	// The BGP IPv4 address for the router on the other end of the BGP session from
	// Oracle. Specified by the owner of that router. If the session goes from Oracle
	// to a customer, this is the BGP IPv4 address of the customer's edge router. If the
	// session goes from Oracle to a provider, this is the BGP IPv4 address of the
	// provider's edge router. Must use a /30 or /31 subnet mask.
	// There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.
	// Example: `10.0.0.18/31`
	CustomerBgpPeeringIp *string `mandatory:"false" json:"customerBgpPeeringIp"`

	// The IPv4 address for Oracle's end of the BGP session. Must use a /30 or /31
	// subnet mask. If the session goes from Oracle to a customer's edge router,
	// the customer specifies this information. If the session goes from Oracle to
	// a provider's edge router, the provider specifies this.
	// There's one exception: for a public virtual circuit, Oracle specifies the BGP IPv4 addresses.
	// Example: `10.0.0.19/31`
	OracleBgpPeeringIp *string `mandatory:"false" json:"oracleBgpPeeringIp"`

	// The number of the specific VLAN (on the cross-connect or cross-connect group)
	// that is assigned to this virtual circuit. Specified by the owner of the cross-connect
	// or cross-connect group (the customer if the customer is colocated with Oracle, or
	// the provider if the customer is connecting via provider).
	// Example: `200`
	Vlan *int `mandatory:"false" json:"vlan"`
}

func (m CrossConnectMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CrossConnectMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
