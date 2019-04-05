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
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
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
	m7a := `{"id": "abc123", "action":"enable", "snap":"helloworld"}`
	m7b := `{"id": "abc123", "action":"enable"}`
	m7c := `{"id": "abc123", "action":"enable", "snap":"invalid"}`
	m8a := `{"id": "abc123", "action":"disable", "snap":"helloworld"}`
	m8b := `{"id": "abc123", "action":"disable"}`
	m8c := `{"id": "abc123", "action":"disable", "snap":"invalid"}`
	m9a := `{"id": "abc123", "action":"conf", "snap":"helloworld"}`
	m9b := `{"id": "abc123", "action":"conf"}`
	m9c := `{"id": "abc123", "action":"conf", "snap":"invalid"}`
	m10a := `{"id": "abc123", "action":"setconf", "snap":"helloworld", "data":"{\"title\": \"Hello World!\"}"}`
	m10b := `{"id": "abc123", "action":"setconf"}`
	m10c := `{"id": "abc123", "action":"setconf", "snap":"invalid", "data":"{\"title\": \"Hello World!\"}"}`
	m10d := `{"id": "abc123", "action":"setconf", "snap":"helloworld", "data":"\u1000"}`

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
		respErr bool
	}{
		{"valid-closed", false, &MockMessage{[]byte(m1a)}, false, false},
		{"valid-open", true, &MockMessage{[]byte(m1a)}, false, false},
		{"no-snap", true, &MockMessage{[]byte(m1b)}, false, true},
		{"invalid-install", true, &MockMessage{[]byte(m1c)}, false, true},

		{"invalid-action", true, &MockMessage{[]byte(m2a)}, true, true},
		{"bad-data", true, &MockMessage{[]byte(m2b)}, true, true},

		{"valid-remove", true, &MockMessage{[]byte(m3a)}, false, false},
		{"no-snap-remove", true, &MockMessage{[]byte(m3b)}, false, true},
		{"invalid-remove", true, &MockMessage{[]byte(m3c)}, false, true},

		{"valid-refresh", true, &MockMessage{[]byte(m4a)}, false, false},
		{"no-snap-refresh", true, &MockMessage{[]byte(m4b)}, false, true},
		{"invalid-refresh", true, &MockMessage{[]byte(m4c)}, false, true},

		{"valid-revert", true, &MockMessage{[]byte(m5a)}, false, false},
		{"no-snap-revert", true, &MockMessage{[]byte(m5b)}, false, true},
		{"invalid-revert", true, &MockMessage{[]byte(m5c)}, false, true},

		{"valid-list", true, &MockMessage{[]byte(m6a)}, false, false},

		{"valid-enable", true, &MockMessage{[]byte(m7a)}, false, false},
		{"no-snap-enable", true, &MockMessage{[]byte(m7b)}, false, true},
		{"invalid-enable", true, &MockMessage{[]byte(m7c)}, false, true},

		{"valid-disable", true, &MockMessage{[]byte(m8a)}, false, false},
		{"no-snap-disable", true, &MockMessage{[]byte(m8b)}, false, true},
		{"invalid-disable", true, &MockMessage{[]byte(m8c)}, false, true},

		{"valid-conf", true, &MockMessage{[]byte(m9a)}, false, false},
		{"no-snap-conf", true, &MockMessage{[]byte(m9b)}, false, true},
		{"invalid-conf", true, &MockMessage{[]byte(m9c)}, false, true},

		{"valid-setconf", true, &MockMessage{[]byte(m10a)}, false, false},
		{"no-snap-setconf", true, &MockMessage{[]byte(m10b)}, false, true},
		{"invalid-setconf", true, &MockMessage{[]byte(m10c)}, false, true},
		{"bad-data-setconf", true, &MockMessage{[]byte(m10d)}, false, true},
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
			resp, err := performAction(sa)
			if err != nil && !tt.withErr {
				t.Error("TestConnection_Workflow: action - expected error got none")
				return
			}

			r, err := deserializePublishResponse(resp)
			if err != nil && !tt.withErr {
				t.Errorf("TestConnection_Workflow: publish response: %v", err)
				return
			}
			if r == nil {
				t.Error("TestConnection_Workflow: publish response is nil")
				return
			}
			if r.Success == tt.respErr {
				t.Errorf("TestConnection_Workflow: publish response unexpected: %s", r.Message)
			}
		})
	}
}

func deserializePublishResponse(data []byte) (*PublishResponse, error) {
	s := PublishResponse{}

	// Decode the message payload - the list of snaps
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println("Error decoding the published message:", err)
	}
	return &s, err
}
