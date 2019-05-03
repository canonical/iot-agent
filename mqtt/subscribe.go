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
	"github.com/CanonicalLtd/iot-devicetwin/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

// performAction acts on the topic and returns a response to publish back
func (c *Connection) performAction(s *SubscribeAction) ([]byte, error) {
	var result domain.PublishResponse
	// Act based on the message action
	switch s.Action {
	case "device":
		result = s.Device(c.organisationID, c.clientID)
	case "list":
		result = s.SnapList()
	case "install":
		result = s.SnapInstall()
	case "remove":
		result = s.SnapRemove()
	case "refresh":
		result = s.SnapRefresh()
	case "revert":
		result = s.SnapRevert()
	case "enable":
		result = s.SnapEnable()
	case "disable":
		result = s.SnapDisable()
	case "conf":
		result = s.SnapConf()
	case "setconf":
		result = s.SnapSetConf()
	case "info":
		result = s.SnapInfo()
	case "ack":
		result = s.SnapAck()
	case "server":
		result = s.SnapServerVersion(c.clientID)

	default:
		return nil, fmt.Errorf("unhandled action: %s", s.Action)
	}

	result.Action = s.Action
	return serializeResponse(result)
}

func deserializePayload(msg MQTT.Message) (*SubscribeAction, error) {
	s := SubscribeAction{}

	// Decode the message payload - the list of snaps
	err := json.Unmarshal(msg.Payload(), &s)
	if err != nil {
		log.Println("Error decoding the subscribed message:", err)
	}
	return &s, err
}

func serializeResponse(resp interface{}) ([]byte, error) {
	return json.Marshal(resp)
}
