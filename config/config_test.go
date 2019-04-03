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
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default-settings-create"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			{
				got := ParseArgs()
				if got.IdentityURL != DefaultIdentityURL {
					t.Errorf("Config.ParseArgs() got = %v, want %v", got.IdentityURL, DefaultIdentityURL)
				}
			}
		})
	}
}
