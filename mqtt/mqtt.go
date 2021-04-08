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
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	twin "github.com/canonical/iot-devicetwin/domain"
	"github.com/canonical/iot-identity/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

// Constants for connecting to the MQTT broker
const (
	quiesce        = 250
	QOSAtMostOnce  = byte(0)
	QOSAtLeastOnce = byte(1)
	//QOSExactlyOnce = byte(2)
)

// Connection for MQTT protocol
type Connection struct {
	client         MQTT.Client
	clientID       string
	organisationID string
}

var conn *Connection
var client MQTT.Client

// GetConnection fetches or creates an MQTT connection
func GetConnection(enroll *domain.Enrollment) (*Connection, error) {
	if conn == nil {
		// Create the client
		client, err := newClient(enroll)
		if err != nil {
			return nil, err
		}

		// Create a new connection
		conn = &Connection{
			client:         client,
			clientID:       enroll.ID,
			organisationID: enroll.Organization.ID,
		}
	}

	// Check that we have a live connection
	if conn.client.IsConnectionOpen() {
		return conn, nil
	}

	// Connect to the MQTT broker
	if token := conn.client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	// Subscribe to the actions topic
	err := conn.SubscribeToActions()

	return conn, err
}

// newClient creates a new MQTT client
func newClient(enroll *domain.Enrollment) (MQTT.Client, error) {
	// Return the active client, if we have one
	if client != nil {
		return client, nil
	}

	// Generate a new MQTT client
	url := fmt.Sprintf("ssl://%s:%s", enroll.Credentials.MQTTURL, enroll.Credentials.MQTTPort)
	log.Println("Connect to the MQTT broker", url)

	// Generate the TLS config from the enrollment credentials
	tlsConfig, err := newTLSConfig(enroll)
	if err != nil {
		return nil, err
	}

	// Set up the MQTT client options
	opts := MQTT.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(enroll.ID)
	opts.SetTLSConfig(tlsConfig)
	opts.AutoReconnect = true
	//opts.SetOnConnectHandler(connectHandler)

	client = MQTT.NewClient(opts)
	return client, nil
}

// newTLSConfig sets up the certificates from the enrollment record
func newTLSConfig(enroll *domain.Enrollment) (*tls.Config, error) {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(enroll.Organization.RootCert)

	// Import client certificate/key pair
	cert, err := tls.X509KeyPair(enroll.Credentials.Certificate, enroll.Credentials.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Create tls.Config with desired TLS properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certPool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}, nil
}

// SubscribeToActions subscribes to the action topic
func (c *Connection) SubscribeToActions() error {
	t := fmt.Sprintf("devices/sub/%s", c.clientID)
	token := c.client.Subscribe(t, QOSAtLeastOnce, c.SubscribeHandler)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error subscribing to topic `%s`: %v", t, token.Error())
		return fmt.Errorf("error subscribing to topic `%s`: %v", t, token.Error())
	}
	return nil
}

// SubscribeHandler is the handler for the main subscription topic
func (c *Connection) SubscribeHandler(client MQTT.Client, msg MQTT.Message) {
	s, err := deserializePayload(msg)
	if err != nil {
		return
	}

	// The topic to publish the response to the specific action
	t := fmt.Sprintf("devices/pub/%s", c.clientID)

	// Perform the action
	response, err := c.performAction(s)
	if err != nil {
		log.Printf("Error with action `%s`: %v", s.Action, err)
	}

	// Publish the response to the action to the broker
	client.Publish(t, QOSAtLeastOnce, false, response)
}

// Health publishes a health message to indicate that the device is still active
func (c *Connection) Health() {
	// Serialize the device health details
	h := twin.Health{
		OrganizationID: c.organisationID,
		DeviceID:       c.clientID,
		Refresh:        time.Now(),
	}
	data, err := json.Marshal(&h)
	if err != nil {
		log.Printf("Error serializing the health data: %v", err)
		return
	}

	// The topic to publish the response to the specific action
	t := fmt.Sprintf("devices/health/%s", c.clientID)
	c.client.Publish(t, QOSAtMostOnce, false, data)
}

// Close closes the connection to the MQTT broker
func (c *Connection) Close() {
	c.client.Disconnect(quiesce)
}
