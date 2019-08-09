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
	"github.com/CanonicalLtd/iot-agent/config"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type StoreTestSuite struct {
	credentialsPath string
	identityURL     string
	snapdMockClient *snapdapi.MockClient
}

var _ = Suite(&StoreTestSuite{})

func (s *StoreTestSuite) SetUpTest(c *C) {
	s.credentialsPath = ".test"
	s.identityURL = ""
	os.Remove(s.credentialsPath)

	s.snapdMockClient = &snapdapi.MockClient{WithError: false}
}

func (s *StoreTestSuite) TearDownSuite(c *C) {
	os.Remove(s.credentialsPath)
}

func (s *StoreTestSuite) TestStoreCredentialsValid(c *C) {
	settings := &config.Settings{
		IdentityURL:     s.identityURL,
		CredentialsPath: s.credentialsPath,
	}

	srv := &Service{
		Settings: settings,
		Snapd:    s.snapdMockClient,
	}
	enroll := domain.Enrollment{}
	err := srv.storeCredentials(enroll)

	c.Assert(err, IsNil)

	_, err = os.Stat(settings.CredentialsPath)
	c.Assert(err, IsNil)
}
