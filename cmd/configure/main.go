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
	"flag"
	"fmt"
	"github.com/canonical/iot-agent/config"
	"os"
)

func main() {
	var (
		url       string
		credsPath string
	)
	flag.StringVar(&url, "url", config.DefaultIdentityURL, "The URL of the Identity Service")
	flag.StringVar(&credsPath, "path", config.GetPath(config.DefaultCredentialsPath), "The full path to the credentials file")
	flag.Parse()

	// Store the URL (let the other parameters be defaulted)
	err := config.StoreParameters(config.Settings{
		IdentityURL: url,
	})
	if err != nil {
		fmt.Println("Error saving parameters:", err)
		os.Exit(1)
	}
}
