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

package main

import (
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"log"
	"time"

	"github.com/CanonicalLtd/iot-agent/config"
	"github.com/CanonicalLtd/iot-agent/identity"
)

const tickInterval = 60

func main() {
	log.Println("Starting IoT agent")

	// Set up the dependency chain
	settings := config.ParseArgs()
	snap := snapdapi.NewClientAdapter()

	// On an interval...
	ticker := time.NewTicker(time.Second * tickInterval)
	for range ticker.C {
		// Check that we are enrolled with the identity service
		idSrv := identity.NewService(settings, snap)
		enroll, err := idSrv.CheckEnrollment()
		if err != nil {
			log.Printf("Error with enrollment: %v", err)
			continue
		}

		log.Println(enroll)

		// Publish scheduled messages
	}
	ticker.Stop()
}
