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

package snapdapi

import (
	"log"

	"github.com/snapcore/snapd/asserts"
)

// GetEncodedAssertions fetches the encoded model and serial assertions
func (a *ClientAdapter) GetEncodedAssertions() ([]byte, error) {
	// Get the model assertion
	modelAssertions, err := a.Known(asserts.ModelType.Name, map[string]string{})
	if err != nil || len(modelAssertions) == 0 {
		log.Printf("error retrieving the model assertion: %v", err)
		return nil, err
	}
	dataModel := asserts.Encode(modelAssertions[0])

	// Get the serial assertion
	serialAssertions, err := a.Known(asserts.SerialType.Name, map[string]string{})
	if err != nil || len(serialAssertions) == 0 {
		log.Printf("error retrieving the serial assertion: %v", err)
		return nil, err
	}
	dataSerial := asserts.Encode(serialAssertions[0])

	// Bring the assertions together
	data := append(dataModel, []byte("\n")...)
	data = append(data, dataSerial...)
	return data, nil
}
