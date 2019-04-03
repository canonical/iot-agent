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
	"fmt"
	"github.com/snapcore/snapd/asserts"
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
type MockClient struct{}

//func (c *MockClient) Snap(name string) (*client.Snap, *client.ResultInfo, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) List(names []string, opts *client.ListOptions) ([]*client.Snap, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Install(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Refresh(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Revert(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Remove(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Enable(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Disable(name string, options *client.SnapOptions) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) ServerVersion() (*client.ServerVersion, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Ack(b []byte) error {
//	panic("implement me")
//}

// Known returns the known assertions for a given type
func (c *MockClient) Known(assertTypeName string, headers map[string]string) ([]asserts.Assertion, error) {
	panic("implement me")
}

//func (c *MockClient) Conf(name string) (map[string]interface{}, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) SetConf(name string, patch map[string]interface{}) (string, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) Find(opts *client.FindOptions) ([]*client.Snap, *client.ResultInfo, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) FindOne(name string) (*client.Snap, *client.ResultInfo, error) {
//	panic("implement me")
//}
//
//func (c *MockClient) FindSnaps(query, section string, private bool) ([]*client.Snap, *client.ResultInfo, error) {
//	panic("implement me")
//}

// GetEncodedAssertions returns the encoded model and serial assertions
func (c *MockClient) GetEncodedAssertions() ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n%s", model1, serial1)), nil
}
