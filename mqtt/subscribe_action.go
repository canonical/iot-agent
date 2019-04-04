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

package mqtt

import (
	"fmt"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"log"
)

type SubscribeAct interface {
	SnapInstall()
}

// SubscribeAction
type SubscribeAction struct {
	ID     string
	Action string
	Snap   string
}

// SnapInstall installs new snaps
func (act *SubscribeAction) SnapInstall() (string, error) {
	if len(act.Snap) == 0 {
		log.Println("Error: no snap name provided for install")
		return "", fmt.Errorf("no snap name provided for install")
	}

	snapd := snapdapi.NewClientAdapter()
	return snapd.Install(act.Snap, nil)
}
