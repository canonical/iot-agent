[![Build Status][travis-image]][travis-url]
[![Go Report Card][goreportcard-image]][goreportcard-url]
[![codecov][codecov-image]][codecov-url]
[![Snap Status](https://build.snapcraft.io/badge/canonical/iot-agent.svg)](https://build.snapcraft.io/user/canonical/iot-agent)
# IoT Agent

The IoT Agent enrolls a device with the [IoT Identity](https://github.com/canonical/iot-identity) service and 
receives credentials to access the MQTT broker. Via MQTT, it establishes communication with
an [IoT Management](https://github.com/canonical/iot-management) service, so the device can be remotely monitored and managed over a
secure connection. The state of the device is mirrored in the cloud by the [IoT Device Twin](https://github.com/canonical/iot-devicetwin) service.

The agent is intended to operate on a device running Ubuntu or Ubuntu Core with snapd enabled. 
The device management features are implemented using the snapd REST API.

  ![IoT Management Solution Overview](docs/IoTManagement.svg)
  

## Build
The project uses vendorized dependencies using govendor. Development has been done on minimum Go version 1.12.1.
```bash
$ go get github.com/canonical/iot-agent
$ cd iot-agent
$ ./get-deps.sh
$ go build ./...
```

## Connect Interfaces
iot-agent uses snapd-control interface which is super powerful. Only the private IoT App Store (aka Brand Store) owners can 
automatically connect the mentioned interface. To use the iot-agent, one can either install it with --devmode or manually connect
snapd-control interface. Please note that a snap with snapd-control interface can not be uploaded to the Global Snap Store. iot-agent
snap is an exception for demonstrating purposes.

```bash
$ snap connect iot-agent:snapd-control :snapd-control
```


## Run
```bash
go run cmd/agent/main.go -help
  -credentials string
        The full path to the credentials file (default ".secret")
  -url string
        The URL of the Identity Service (default "http://localhost:8030/")
```

## Contributing
Before contributing you should sign [Canonical's contributor agreement](https://www.ubuntu.com/legal/contributors), it’s the easiest way for you to give us permission to use your contributions.

[travis-image]: https://travis-ci.org/canonical/iot-agent.svg?branch=master
[travis-url]: https://travis-ci.org/canonical/iot-agent
[goreportcard-image]: https://goreportcard.com/badge/github.com/canonical/iot-agent
[goreportcard-url]: https://goreportcard.com/report/github.com/canonical/iot-agent
[codecov-url]: https://codecov.io/gh/canonical/iot-agent
[codecov-image]: https://codecov.io/gh/canonical/iot-agent/branch/master/graph/badge.svg