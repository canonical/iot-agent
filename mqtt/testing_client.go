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
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// MockClient mocks the MQTT client
type MockClient struct {
	open bool
}

// IsConnected mocks the connect status
func (cli *MockClient) IsConnected() bool {
	return cli.open
}

// IsConnectionOpen mocks the connect status
func (cli *MockClient) IsConnectionOpen() bool {
	return cli.open
}

// Connect mocks connecting to broker
func (cli *MockClient) Connect() MQTT.Token {
	cli.open = true
	return &MockToken{}
}

// Disconnect mocks client close
func (cli *MockClient) Disconnect(quiesce uint) {
	cli.open = false
	return
}

// Publish mocks a publish message
func (cli *MockClient) Publish(topic string, qos byte, retained bool, payload interface{}) MQTT.Token {
	return &MockToken{}
}

// Subscribe mocks a subscribe message
func (cli *MockClient) Subscribe(topic string, qos byte, callback MQTT.MessageHandler) MQTT.Token {
	return &MockToken{}
}

// SubscribeMultiple mocks subscribe messages
func (cli *MockClient) SubscribeMultiple(filters map[string]byte, callback MQTT.MessageHandler) MQTT.Token {
	return &MockToken{}
}

// Unsubscribe mocks a unsubscribe message
func (cli *MockClient) Unsubscribe(topics ...string) MQTT.Token {
	return &MockToken{}
}

// AddRoute mocks routing
func (cli *MockClient) AddRoute(topic string, callback MQTT.MessageHandler) {
	return
}

// OptionsReader mocks the options reader (badly)
func (cli *MockClient) OptionsReader() MQTT.ClientOptionsReader {
	return MQTT.NewClient(nil).OptionsReader()
}

// MockToken implements a Token
type MockToken struct{}

// Wait mocks the token wait
func (t *MockToken) Wait() bool {
	return true
}

// WaitTimeout mocks the token wait timeout
func (t *MockToken) WaitTimeout(time.Duration) bool {
	return true
}

// Error mocks a token error check
func (t *MockToken) Error() error {
	return nil
}

// MockMessage implements an MQTT message
type MockMessage struct {
	message []byte
}

// Duplicate mocks a duplicate message check
func (m *MockMessage) Duplicate() bool {
	panic("implement me")
}

// Qos mocks the QoS flag
func (m *MockMessage) Qos() byte {
	panic("implement me")
}

// Retained mocks the retained flag
func (m *MockMessage) Retained() bool {
	panic("implement me")
}

// Topic mocks the topic
func (m *MockMessage) Topic() string {
	panic("implement me")
}

// MessageID mocks the message ID
func (m *MockMessage) MessageID() uint16 {
	return 1000
}

// Payload mocks the payload retrieval
func (m *MockMessage) Payload() []byte {
	return m.message
}

// Ack mocks the message ack
func (m *MockMessage) Ack() {
	panic("implement me")
}
