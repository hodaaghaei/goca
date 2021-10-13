/* -------------------------------------------------------------------------- */
/* Copyright 2002-2021, OpenNebula Project, OpenNebula Systems                */
/*                                                                            */
/* Licensed under the Apache License, Version 2.0 (the "License"); you may    */
/* not use this file except in compliance with the License. You may obtain    */
/* a copy of the License at                                                   */
/*                                                                            */
/* http://www.apache.org/licenses/LICENSE-2.0                                 */
/*                                                                            */
/* Unless required by applicable law or agreed to in writing, software        */
/* distributed under the License is distributed on an "AS IS" BASIS,          */
/* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   */
/* See the License for the specific language governing permissions and        */
/* limitations under the License.                                             */
/*--------------------------------------------------------------------------- */

package securitygroup

import (
	"encoding/xml"

	"github.com/hodaaghaei/onego/schemas/shared"
)

// Pool represents an OpenNebula SecurityGroup pool
type Pool struct {
	XMLName        xml.Name        `xml:"SECURITY_GROUP_POOL"`
	SecurityGroups []SecurityGroup `xml:"SECURITY_GROUP"`
}

// SecurityGroup represents an OpenNebula SecurityGroup
type SecurityGroup struct {
	XMLName     xml.Name            `xml:"SECURITY_GROUP"`
	ID          int                 `xml:"ID,omitempty"`
	UID         int                 `xml:"UID,omitempty"`
	GID         int                 `xml:"GID,omitempty"`
	UName       string              `xml:"UNAME,omitempty"`
	GName       string              `xml:"GNAME,omitempty"`
	Name        string              `xml:"NAME"`
	Permissions *shared.Permissions `xml:"PERMISSIONS,omitempty"`
	UpdatedVMs  shared.EntitiesID   `xml:"UPDATED_VMS,omitempty"`
	OutdatedVMs shared.EntitiesID   `xml:"OUTDATED_VMS,omitempty"`
	UpdatingVMs shared.EntitiesID   `xml:"UPDATING_VMS,omitempty"`
	ErrorVMs    shared.EntitiesID   `xml:"ERROR_VMS,omitempty"`
	Template    Template            `xml:"TEMPLATE"`
}
