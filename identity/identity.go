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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/CanonicalLtd/iot-agent/config"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	"github.com/CanonicalLtd/iot-identity/web"
)

// Default parameters
const (
	mediaType          = "application/x.ubuntu.assertion"
	commonDataEnvVar   = "SNAP_COMMON"
	deviceDataFileName = "device-data.bin"
)

// UseCase is the interface for the identity service use cases
type UseCase interface {
	CheckEnrollment() (*domain.Enrollment, error)
}

// Service implements the identity service use cases
type Service struct {
	Settings *config.Settings
	Snapd    snapdapi.SnapdClient
}

// NewService creates a new identity service connection
func NewService(settings *config.Settings, snapd snapdapi.SnapdClient) *Service {
	return &Service{
		Settings: settings,
		Snapd:    snapd,
	}
}

// CheckEnrollment verifies that the device is enrolled with the identity service
func (srv *Service) CheckEnrollment() (*domain.Enrollment, error) {
	// Get the credentials from the filesystem
	en, err := srv.getCredentials()
	if err == nil {
		return en, nil
	}

	// No credentials stored, so enroll the device
	// Enroll the device with the identity service
	return srv.enrollDevice()
}

// enroll registers the device with the identity service
func (srv *Service) enrollDevice() (*domain.Enrollment, error) {
	// Get the model and serial assertions
	data, err := srv.Snapd.GetEncodedAssertions()
	if err != nil {
		return nil, err
	}

	// Format the URL for the identity service
	resp, err := sendEnrollmentRequest(srv.Settings.IdentityURL, data)
	if err != nil {
		return nil, err
	}

	// Store the enrollment credentials
	err = srv.storeCredentials(resp.Enrollment)
	if err != nil {
		return nil, err
	}

	// Store device data in a separate file
	if resp.Enrollment.DeviceData != "" {
		err = storeDeviceData(resp.Enrollment.DeviceData)
		if err != nil {
			return nil, err
		}
	}

	return &resp.Enrollment, err
}

func storeDeviceData(dataBase64 string) error {
	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return fmt.Errorf("Cannot decode device data: %v", err)
	}

	err = ioutil.WriteFile(path.Join(os.Getenv(commonDataEnvVar), deviceDataFileName), data, 0600)
	if err != nil {
		return fmt.Errorf("Cannot write device data: %v", err)
	}

	return nil
}

func sendEnrollmentRequest(idURL string, data []byte) (*web.EnrollResponse, error) {
	// Format the URL for the identity service
	u, err := url.Parse(idURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "v1", "device", "enroll")

	// Send the request to get the credentials from the identity service
	resp, err := sendPOSTRequest(u.String(), data)
	if err != nil {
		return nil, err
	}

	if len(resp.StandardResponse.Code) > 0 {
		return nil, fmt.Errorf("(%s) %s", resp.StandardResponse.Code, resp.StandardResponse.Message)
	}

	return resp, nil
}

func parseEnrollResponse(r io.Reader) (*web.EnrollResponse, error) {
	// Parse the response
	result := web.EnrollResponse{}
	err := json.NewDecoder(r).Decode(&result)
	return &result, err
}

var sendPOSTRequest = func(u string, data []byte) (*web.EnrollResponse, error) {
	// Send the request to get the credentials from the identity service
	w, err := http.Post(u, mediaType, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer w.Body.Close()
	return parseEnrollResponse(w.Body)
}
