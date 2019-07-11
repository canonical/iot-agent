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

import (
	"os"
	"testing"
)

func TestReadParameters(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default-settings-create"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{
				got := ReadParameters()
				if got.IdentityURL != DefaultIdentityURL {
					t.Errorf("Config.ReadParameters() got = %v, want %v", got.IdentityURL, DefaultIdentityURL)
				}
				if got.CredentialsPath != DefaultCredentialsPath {
					t.Errorf("Config.ReadParameters() got = %v, want %v", got.CredentialsPath, DefaultCredentialsPath)
				}

				_ = os.Remove(paramsFilename)
			}
		})
	}
}

func TestStoreParameters(t *testing.T) {
	tests := []struct {
		name    string
		args    Settings
		wantErr bool
	}{
		{"valid", Settings{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StoreParameters(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("StoreParameters() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := ReadParameters()
			if got.IdentityURL != DefaultIdentityURL {
				t.Errorf("Config.ReadParameters() got = %v, want %v", got.IdentityURL, DefaultIdentityURL)
			}
			if got.CredentialsPath != DefaultCredentialsPath {
				t.Errorf("Config.ReadParameters() got = %v, want %v", got.CredentialsPath, DefaultCredentialsPath)
			}

			_ = os.Remove(paramsFilename)
		})
	}
}
