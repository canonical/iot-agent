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
	"github.com/CanonicalLtd/iot-identity/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"testing"
)

func TestConnection_Workflow(t *testing.T) {
	m1 := `{"id": "abc123", "action":"install", "snap":"helloworld"}`
	m2 := `{"id": "abc123", "action":"install"}`
	m3 := `{"id": "abc123", "action":"invalid", "snap":"helloworld"}`
	m4 := `{"id": "abc123", "action":"remove", "snap":"helloworld"}`
	m5 := `{"id": "abc123", "action":"remove"}`
	m6 := `\u1000`

	enroll := &domain.Enrollment{
		Credentials: domain.Credentials{
			MQTTURL:  "localhost",
			MQTTPort: "8883",
		},
	}
	client = &MockClient{}
	tests := []struct {
		name    string
		open    bool
		message MQTT.Message
		withErr bool
	}{
		{"valid-closed", false, &MockMessage{[]byte(m1)}, false},
		{"valid-open", true, &MockMessage{[]byte(m1)}, false},
		{"no-snap", true, &MockMessage{[]byte(m2)}, false},
		{"invalid-action", true, &MockMessage{[]byte(m3)}, true},
		{"valid-remove", true, &MockMessage{[]byte(m4)}, false},
		{"no-snap-remove", true, &MockMessage{[]byte(m5)}, false},
		{"bad-data", true, &MockMessage{[]byte(m6)}, true},
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
			sa, err := subscribePayload(tt.message)
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
