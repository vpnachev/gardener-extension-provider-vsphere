/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// Object to identify an uplink based on its type and name
type Uplink struct {

	// Name of this uplink
	UplinkName string `json:"uplink_name"`

	// Type of the uplink
	UplinkType string `json:"uplink_type"`
}
