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
	"github.com/CanonicalLtd/iot-agent/config"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"valid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()

			got := config.ReadParameters()
			if got.IdentityURL != config.DefaultIdentityURL {
				t.Errorf("Config.ReadParameters() got = %v, want %v", got.IdentityURL, config.DefaultIdentityURL)
			}
			if got.CredentialsPath != config.DefaultCredentialsPath {
				t.Errorf("Config.ReadParameters() got = %v, want %v", got.CredentialsPath, config.DefaultCredentialsPath)
			}

			_ = os.Remove("params")
		})
	}
}
