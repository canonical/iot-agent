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

package identity

import (
	"encoding/json"
	"github.com/canonical/iot-identity/domain"
	"io/ioutil"
)

// storeCredentials caches the serialized enrollment details
func (srv *Service) storeCredentials(enroll domain.Enrollment) error {
	data, err := json.Marshal(&enroll)
	if err != nil {
		return nil
	}

	return ioutil.WriteFile(srv.Settings.CredentialsPath, data, 0600)
}

// getCredentials fetches the cached enrollment details
func (srv *Service) getCredentials() (*domain.Enrollment, error) {
	enroll := &domain.Enrollment{}

	// Read the credentials from the filesystem
	data, err := ioutil.ReadFile(srv.Settings.CredentialsPath)
	if err != nil {
		return nil, err
	}

	// Deserialize the credentials
	err = json.Unmarshal(data, enroll)
	return enroll, err
}
