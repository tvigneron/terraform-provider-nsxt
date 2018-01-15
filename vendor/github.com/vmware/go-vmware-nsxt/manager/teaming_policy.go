/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// Uplink Teaming Policy
type TeamingPolicy struct {

	// List of Uplinks used in active list
	ActiveList []Uplink `json:"active_list"`

	// Teaming policy
	Policy string `json:"policy"`

	// List of Uplinks used in standby list
	StandbyList []Uplink `json:"standby_list,omitempty"`
}