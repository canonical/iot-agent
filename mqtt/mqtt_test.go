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
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"testing"
)

func TestConnection_Workflow(t *testing.T) {
	m1a := `{"id": "abc123", "action":"install", "snap":"helloworld"}`
	m1b := `{"id": "abc123", "action":"install"}`
	m1c := `{"id": "abc123", "action":"install", "snap":"invalid"}`
	m2a := `{"id": "abc123", "action":"invalid", "snap":"helloworld"}`
	m2b := `\u1000`
	m3a := `{"id": "abc123", "action":"remove", "snap":"helloworld"}`
	m3b := `{"id": "abc123", "action":"remove"}`
	m3c := `{"id": "abc123", "action":"remove", "snap":"invalid"}`
	m4a := `{"id": "abc123", "action":"refresh", "snap":"helloworld"}`
	m4b := `{"id": "abc123", "action":"refresh"}`
	m4c := `{"id": "abc123", "action":"refresh", "snap":"invalid"}`
	m5a := `{"id": "abc123", "action":"revert", "snap":"helloworld"}`
	m5b := `{"id": "abc123", "action":"revert"}`
	m5c := `{"id": "abc123", "action":"revert", "snap":"invalid"}`
	m6a := `{"id": "abc123", "action":"list"}`

	enroll := &domain.Enrollment{
		Credentials: domain.Credentials{
			MQTTURL:  "localhost",
			MQTTPort: "8883",
		},
	}
	client = &MockClient{}
	snapd = &snapdapi.MockClient{}
	tests := []struct {
		name    string
		open    bool
		message MQTT.Message
		withErr bool
	}{
		{"valid-closed", false, &MockMessage{[]byte(m1a)}, false},
		{"valid-open", true, &MockMessage{[]byte(m1a)}, false},
		{"no-snap", true, &MockMessage{[]byte(m1b)}, false},
		{"invalid-install", true, &MockMessage{[]byte(m1c)}, false},

		{"invalid-action", true, &MockMessage{[]byte(m2a)}, true},
		{"bad-data", true, &MockMessage{[]byte(m2b)}, true},

		{"valid-remove", true, &MockMessage{[]byte(m3a)}, false},
		{"no-snap-remove", true, &MockMessage{[]byte(m3b)}, false},
		{"invalid-remove", true, &MockMessage{[]byte(m3c)}, false},

		{"valid-refresh", true, &MockMessage{[]byte(m4a)}, false},
		{"no-snap-refresh", true, &MockMessage{[]byte(m4b)}, false},
		{"invalid-refresh", true, &MockMessage{[]byte(m4c)}, false},

		{"valid-revert", true, &MockMessage{[]byte(m5a)}, false},
		{"no-snap-revert", true, &MockMessage{[]byte(m5b)}, false},
		{"invalid-revert", true, &MockMessage{[]byte(m5c)}, false},

		{"valid-list", true, &MockMessage{[]byte(m6a)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.open {
				client.Connect()
			}
			c, err := GetConnection(enroll)
			if err != nil {
				t.Error("TestConnection_Workflow: error creating connection")
				return
			}
			if c.client == nil {
				t.Error("TestConnection_Workflow: no client created")
			}

			// Publish the health
			c.Health()

			// Subscribe action
			c.SubscribeHandler(client, tt.message)

			// Check again with the action
			sa, err := deserializePayload(tt.message)
			if err != nil && !tt.withErr {
				t.Error("TestConnection_Workflow: payload - expected error got none")
				return
			}
			_, err = performAction(sa)
			if err != nil && !tt.withErr {
				t.Error("TestConnection_Workflow: action - expected error got none")
				return
			}
		})
	}
}
