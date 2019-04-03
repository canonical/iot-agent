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

package config

import "flag"

// Default settings
const (
	DefaultIdentityURL     = "http://localhost:8030/"
	DefaultCredentialsPath = ".secret"
)

// Settings defines the application configuration
type Settings struct {
	IdentityURL     string
	CredentialsPath string
}

// ParseArgs checks the command line arguments
func ParseArgs() *Settings {
	var (
		url      string
		credPath string
	)
	flag.StringVar(&url, "url", DefaultIdentityURL, "The URL of the Identity Service")
	flag.StringVar(&credPath, "credentials", DefaultCredentialsPath, "The full path to the credentials file")
	flag.Parse()

	return &Settings{
		IdentityURL:     url,
		CredentialsPath: credPath,
	}
}
