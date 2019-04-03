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
	"fmt"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	"github.com/CanonicalLtd/iot-identity/web"
	"os"
	"testing"

	"github.com/CanonicalLtd/iot-agent/config"
)

func TestService_CheckEnrollment(t *testing.T) {
	settings := config.ParseArgs()
	_ = os.Remove(settings.CredentialsPath)
	snapd := &snapdapi.MockClient{}
	tests := []struct {
		name    string
		sendErr bool
		wantErr bool
		cleanUp bool
	}{
		{"valid", false, false, false},
		{"valid-secret", false, false, true},
		{"send-error", true, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sendErr {
				sendEnrollmentRequest = mockSendRequestError
			} else {
				sendEnrollmentRequest = mockSendRequest
			}

			srv := NewService(settings, snapd)
			got, err := srv.CheckEnrollment()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckEnrollment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got.ID) == 0 {
					t.Error("Service.CheckEnrollment() error = empty enrollment")
				}
			}

			// Clean up
			if tt.cleanUp {
				_ = os.Remove(settings.CredentialsPath)
			}
		})
	}
}

func mockSendRequest(idURL string, data []byte) (*web.EnrollResponse, error) {
	return &web.EnrollResponse{Enrollment: domain.Enrollment{ID: "abc123"}}, nil
}

func mockSendRequestError(idURL string, data []byte) (*web.EnrollResponse, error) {
	return nil, fmt.Errorf("mock send request error")
}
