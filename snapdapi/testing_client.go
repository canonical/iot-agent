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
	"bytes"
	"fmt"
	"github.com/snapcore/snapd/asserts"
	"github.com/snapcore/snapd/client"
)

const model1 = `type: model
authority-id: canonical
series: 16
brand-id: canonical
model: ubuntu-core-18-amd64
architecture: amd64
base: core18
display-name: Ubuntu Core 18 (amd64)
gadget: pc=18
kernel: pc-kernel=18
timestamp: 2018-08-13T09:00:00+00:00
sign-key-sha3-384: 9tydnLa6MTJ-jaQTFUXEwHl1yRx7ZS4K5cyFDhYDcPzhS7uyEkDxdUjg9g08BtNn
AcLBXAQAAQoABgUCW37NBwAKCRDgT5vottzAEut9D/4u9lD3lFWXoHx1VQT+mUCROcFHdXQBY/PJ
NriRiDwBaOjEo5mvHMRJ2UulWvHnwqyMJctJKBP+RCKlrJEPX8eaLP/lmihwIiFfmzm49BLaNwli
si0entond1sVWfiNr7azXoEuAIgYvxmJIvE+GZADDT0/OTFQRcLU69bhNEAQKBnkT0y/HTpuXwlJ
TuwwJtDR0vZuFtwzj6Bdx7W42+vGmuXE7M4Ni6HUySNKYByB5BsrDf3/79p8huXyBtnWp+HBsHtb
fgjzQoBcspj65Gi+crBrJ4jS+nfowRRVXLL1clXJOJLz12za+kN0/FC0PhussiQb5UI7USXJ+RvA
Y8U1vrqG7bG5GYGqe1KB9GbLEm+GBPQZcZI3jRmm9V7tm9OWQzK98/uPwTD73IW7LrDT35WQrIYM
fBfThJcRqpgzwZD/CBx82maLB9tmsRF5Mhcj2H1v7cn8nSkbv7+cCzh25lKv48Vqz1WTgO3HMPWW
0kb6BSoC+YGpstSUslqtpLdY/MfFI0DhshH2Y+h0c9/g4mux/Zb8Gs9V55HGn9mr2KKDmHsU2k+C
maZWcXOxRpverZ2Pi9L4fZxhZ9H+FDcMGiHn2vJFQhI3u+LiK3aUUAov4k3vNRPGSvi1AGhuEtUa
NG54bznx12KgOT3+YiHtfE95WiXUcJUrEXAgfVBVoA==`

const serial1 = `type: serial
authority-id: canonical
brand-id: canonical
model: ubuntu-core-18-amd64
serial: d75f7300-abbf-4c11-bf0a-8b7103038490
device-key:
    AcbBTQRWhcGAARAA05GC1FmdsBVDxd2DbolPLiqnQXDDwW0RScEcuG5ONGMmvolfS4DJxS5ONBq2
    ZdvGYoCzuSE4P/fruKwrfnR+DRn+frA2YAQOagHy2xmSYlXBz1wyDAvKVmJdv7Q2EjGK4K6vgVMn
    v8No+9/fecoIF7oa9kF7EwcnDrN89VGR+jOljGvwJ3QKHh8Tq5szL3ETlhdv4E6GEt4lEjcw3hDM
    rjGezRwM9riypbJp3paNWygff03sC6Q5esZk9U2ijF7tEF7CT5zCZEaLs+OdOQxYL6R4Bw7lp2h2
    xj/0G6pX3AH/VtijIJj/aOn6fBQB9kzGEghjUemHKqfpJ7lEH/TQ0JIMj9z/Tgj5KDPXEgtwgf78
    37TYbDxcfoFJbi4sMoXFoKq2d2b8ufnQ1UlxMiCxr/z3GtraxDhMRx34vxIr1RqhHGt48as0rLjF
    mnsOAxSOhyloVgd9V5jdK7gzCi6aTtNZTMJV5TkGo3HyMEmDmj+TLAmPrENVt2A/EnKEyORz+0o1
    5qtauqdcypOyAQc1aPmbGtqX5adI8tuj6JLxXdcQgCsQp+F5j+NM9TZnNnbwjkWZam1G8seGH+GZ
    QpeT5+5VqhXIkmlk8Mfqgn5br/1D7dfjBrzAumBpOmcOIeCCYrBtlpva4+nnO3Hp6bmkfuYBNXZe
    jJJS3M6FTNApbr0AEQEAAQ==
device-key-sha3-384: xm9bu3yCuJguaB233yCAnXDE9zgOu8V39-2j8c-Rk0R27HjQpruF8ce_vGZDEm-G
timestamp: 2019-01-10T17:40:44.771564Z
sign-key-sha3-384: BWDEoaqyr25nF5SNCvEv2v7QnM9QsfCc0PBMYD_i2NGSQ32EF2d4D0hqUel3m8ul
AcLBUgQAAQoABgUCXDeDnAAAnLMQAG6jJOffkqDrUhbgMP6VBmGr9nTm54fUg+pMYvxVxex6o4vH
thA5qtQE9of1UVAK5qX7qwwl3rsIZ1/ESagW1ME1hyrCcVxcZ63BQrLODj9VX0kp8VmBvgUWGIsw
sS/ZidF4lbsanWyzFefCErgzAncjxGN9cpMUsJPd5ai2c6Iq9+8qvJoT6ubWWg0Nh/Fe+jURKTs8
Sfzfz0vaySoSmuH4cOYShz2tYvVEVvJyaoNt5vLUrG2TKgA5tz1S0mKwhwDbGRwKFL6mQSlJ/L5N
P6UKSpZKfin+/ziH5YV0PoY3pTeTbuoMQWknYqQUBN/rHzd1y6xmY6rcWsZkFN2sPqA57ZgxUW4C
h/3TZDyRUNXSGqiam5lKEx1EUWiWHhZG6TtOG8+pOW+Y+uW8v1c2qKKHIghQHAgZjUzaNyec2Ylw
PfZW5UO8ua37jvSDV4aYcDXLlumD76mCQkXslltXATOnH9ZDMaf7/MRnx7Dwaqu0kuYUCNSWN/kJ
oe5AnCaMg/yTp0EbV9ZlHNeQYGesUkhT9ULXzsUEfhs3S6mQtnC12O1C/F7fsv1x7lSa4WvPzlb7
Azds7xIR91OzXGFMx/PO7ZwflxBRIZw7+iFXEXWzfhzVlrUFDLr8K++g1g563UzY9P86XwGDlS7l
/PVxRaD/Ruiw0ey94zCcn3ROBEs/`

// MockClient is a mock snapd API client for testing
type MockClient struct {
	WithError bool
}

// Snap mocks the details of a snap
func (c *MockClient) Snap(name string) (*client.Snap, *client.ResultInfo, error) {
	if name == "invalid" {
		return nil, nil, fmt.Errorf("MOCK error install")
	}
	return &client.Snap{
		ID:      "1",
		Title:   "helloworld",
		Summary: "Welcomes the world",
	}, &client.ResultInfo{}, nil
}

// List lists installed snaps
func (c *MockClient) List(names []string, opts *client.ListOptions) ([]*client.Snap, error) {
	if c.WithError {
		return nil, fmt.Errorf("MOCK error snap list")
	}
	return []*client.Snap{
		{
			ID:      "1",
			Title:   "helloworld",
			Summary: "Welcomes the world",
		},
	}, nil
}

// Install mocks the snap installation
func (c *MockClient) Install(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error install")
	}
	return "100", nil
}

// Refresh refreshes an existing snap
func (c *MockClient) Refresh(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error refresh")
	}
	return "102", nil
}

// Revert reverts an existing snap
func (c *MockClient) Revert(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error revert")
	}
	return "103", nil
}

// Remove mocks a snap removal
func (c *MockClient) Remove(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error remove")
	}
	return "101", nil
}

// Enable mocks a snap enable
func (c *MockClient) Enable(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error enable")
	}
	return "104", nil
}

// Disable mocks a snap disable
func (c *MockClient) Disable(name string, options *client.SnapOptions) (string, error) {
	if name == "invalid" {
		return "", fmt.Errorf("MOCK error disable")
	}
	return "105", nil
}

// ServerVersion mocks the server version call
func (c *MockClient) ServerVersion() (*client.ServerVersion, error) {
	if c.WithError {
		return nil, fmt.Errorf("MOCK error server version")
	}
	return &client.ServerVersion{
		Version:       "1.0",
		KernelVersion: "generic-kernel.5.0",
	}, nil
}

// Ack mocks adding an assertion
func (c *MockClient) Ack(b []byte) error {
	if bytes.Equal(b, []byte("invalid")) {
		return fmt.Errorf("MOCK error ack")
	}
	return nil
}

// Known returns the known assertions for a given type
func (c *MockClient) Known(assertTypeName string, headers map[string]string) ([]asserts.Assertion, error) {
	panic("implement me")
}

// Conf mocks returning config
func (c *MockClient) Conf(name string) (map[string]interface{}, error) {
	if name == "invalid" {
		return nil, fmt.Errorf("MOCK error conf")
	}
	return map[string]interface{}{"setting": "value"}, nil
}

// SetConf mocks setting the config
func (c *MockClient) SetConf(name string, patch map[string]interface{}) (string, error) {
	if name == "invalid" {
		return "106", fmt.Errorf("MOCK error set conf")
	}
	return "106", nil
}

// GetEncodedAssertions returns the encoded model and serial assertions
func (c *MockClient) GetEncodedAssertions() ([]byte, error) {
	if c.WithError {
		return nil, fmt.Errorf("MOCK error known")
	}
	return []byte(fmt.Sprintf("%s\n%s", model1, serial1)), nil
}

// DeviceInfo returns a mock device details
func (c *MockClient) DeviceInfo() (ActionDevice, error) {
	if c.WithError {
		return ActionDevice{}, fmt.Errorf("MOCK error device info")
	}
	return ActionDevice{
		Brand:        "example",
		Model:        "drone-1000",
		SerialNumber: "A11111111",
		DeviceKey:    "AAAAAAAAAA",
		StoreID:      "example-store",
	}, nil
}
