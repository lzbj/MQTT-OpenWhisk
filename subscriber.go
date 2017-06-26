package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
	"fmt"
	"encoding/json"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//topic:=msg.Topic()
	pld := msg.Payload()
	//fmt.Printf("%s\n", topic)
	fmt.Printf("%s\n", pld)
}

// args[1] the broker url, such as "tcp://localhost:1883"
// args[2] the client id, such as "client2"
// args[3] the topic name, such as "topic/sample"
func main() {
	checkArgs()
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(os.Args[1]).SetClientID(os.Args[2])
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(os.Args[3], 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		time.Sleep(5 * time.Second)
	}
}

func checkArgs() {
	if (len(os.Args[1:]) != 3) {
		log.Println("args not correct")
		msg := "args not correct"
		ms, _ := json.Marshal(msg)
		fmt.Println(string(ms))
		os.Exit(1)
	}
}
