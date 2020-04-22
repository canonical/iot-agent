// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Agent
 * Copyright 2020 Canonical Ltd.
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
	"fmt"
	"log"
	"math/rand"
)

// Metrics publishes a metrics messages to indicate so the device can be monitored
func (c *Connection) Metrics() {
	payload := fmt.Sprintf("metrics,device=%s cpu=%f,mem=%d", c.clientID, rand.Float64()*100, rand.Intn(2000))
	log.Println(payload)

	// The topic to publish the response to the specific action
	t := fmt.Sprintf("metrics/%s", c.organisationID)
	c.client.Publish(t, QOSAtMostOnce, false, []byte(payload))
}
