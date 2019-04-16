// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Agent
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mqtt

import (
	"encoding/json"
	"time"
)

// PublishResponse is the published message showing the result of an action
type PublishResponse struct {
	ID      string      `json:"id"`
	Action  string      `json:"action"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Health update contains enough details to record a device
type Health struct {
	OrganizationID string    `json:"orgId"`
	DeviceID       string    `json:"deviceId"`
	Refresh        time.Time `json:"refresh"`
}

// ActionDevice details
type ActionDevice struct {
	OrganizationID string `json:"orgId"`
	DeviceID       string `json:"deviceId"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	SerialNumber   string `json:"serial"`
	StoreID        string `json:"store"`
	DeviceKey      string `json:"deviceKey"`
}

func serializeResponse(resp interface{}) ([]byte, error) {
	return json.Marshal(resp)
}
