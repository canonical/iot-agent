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
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

// performAction acts on the topic and returns a response to publish back
func performAction(s *SubscribeAction) ([]byte, error) {
	// Act based on the message action
	switch s.Action {
	case "install":
		r, err := s.SnapInstall()
		if err != nil {
			return serializeResponse(PublishResponse{ID: s.ID, Success: false, Message: err.Error()})
		}
		return serializeResponse(PublishResponse{ID: s.ID, Success: true, Message: r})

	case "remove":
		r, err := s.SnapRemove()
		if err != nil {
			return serializeResponse(PublishResponse{ID: s.ID, Success: false, Message: err.Error()})
		}
		return serializeResponse(PublishResponse{ID: s.ID, Success: true, Message: r})

	default:
		return nil, fmt.Errorf("unhandled action: %s", s.Action)
	}
}

func subscribePayload(msg MQTT.Message) (*SubscribeAction, error) {
	s := SubscribeAction{}

	// Decode the message payload - the list of snaps
	err := json.Unmarshal(msg.Payload(), &s)
	if err != nil {
		log.Println("Error decoding the subscribed message:", err)
	}
	return &s, err
}
