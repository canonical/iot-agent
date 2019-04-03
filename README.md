# IoT Agent

The IoT Agent enrolls a device with the [IoT Identity](https://github.com/CanonicalLtd/iot-identity) Service and 
receives credentials to access the MQTT broker. Via MQTT, it establishes communication with
an IoT management service, so the device can be remotely monitored and managed over a
secure connection.

The agent is intended to operate on a device running Ubuntu or Ubuntu Core with snapd enabled. 
The device management features are implemented using the snapd REST API.

## Build
The project uses vendorized dependencies using govendor. Development has been done on minimum Go version 1.12.1.
```bash
$ go get github.com/CanonicalLtd/iot-agent
$ cd iot-agent
$ ./get-deps.sh
$ go build ./...
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
