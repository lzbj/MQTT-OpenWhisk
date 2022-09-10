# MQTT-OpenWhisk.

This repo is to demo that how MQTT pub-sub use case could be operated as actions inside
OpenWhisk platform. Since MQTT is a commonly used IoT protocol, so OpenWhisk could be
integrated to IoT solutions seamlessly.

[OpenWhisk](http://openwhisk.incubator.apache.org/) is an Open Source Serverless 
computing platform donated to Apache foundation by IBM.

# About MQTT.
As it's website illustrated, [MQTT](https://mqtt.org) is a machine-to-machine 
(M2M)/"Internet of Things" connectivity protocol. It was designed as an extremely 
lightweight publish/subscribe messaging transport.

# Installation.
This section will guide you through the installation of this package.

## Mosquitto.
We used [Mosquitto](https://mosquitto.org/) as the Open Source broker, you could find the
installation related details [here](https://mosquitto.org/download/).
After installation, please run 
`mosquitto -d`, this command will start a MQTT broker on your local host.

## Paho.
We used the [Paho](http://www.eclipse.org/paho/) library as the client solution, since we used Golang 
client, run the following command to get the library installed.
`go get github.com/eclipse/paho.mqtt.golang`

## OpenWhisk.
You could run this demo against IBM [Bluemix](https://console.bluemix.net/) OpenWhisk. If
you choose to run on a local OpenWhisk deployment, you could find the doc [here](https://github.com/apache/incubator-openwhisk#native-development)

## Golang or Docker.
If you choose to compile the code locally, you need to install [golang](https://golang.org) on your host.
We recommend to run the demo with [Docker](https://www.docker.com). So you will need a docker 
daemon running on your local host.

### Publisher.
After clone the code repo to your local host. Run the following command.
`cp publisher.go exec.go`
And then build it with the following command.
`env GOOS=linux GOARCH=amd64 go build exec.go`
After that let's zip it and upload it to OpenWhisk.
`zip publisher.zip exec`
`wsk action create publisher publisher.zip --docker`
After these commands, you have successfully created the publisher action on OpenWhisk.

### Subscriber.
For the subscriber, you can reference the above section. There is a good [blog](http://jamesthom.as/blog/2017/01/17/openwhisk-and-go/) 
about how to build and run Golang action on OpenWhisk.

## Run.
TBD, stay tuned.





