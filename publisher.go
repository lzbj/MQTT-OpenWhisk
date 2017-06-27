package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
	"encoding/json"
	"fmt"
)

// args[1] the broker url, such as "tcp://localhost:1883"
// args[2] the client id, such as "client2"
// args[3] the message to publish, "message sample"
// args[4] the topic name, such as "topic/sample"
func main() {
	checkArg()
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(os.Args[1]).SetClientID(os.Args[2])
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	msg := map[string]string{"msg": ("Hello, " + os.Args[3] + "!")}
	res, _ := json.Marshal(msg)
	token := c.Publish(os.Args[4], 0, false, res)
	token.Wait()
	time.Sleep(20 * time.Second)
	c.Disconnect(250)
	msg = map[string]string{"msg": ("published")}
	res, _ = json.Marshal(msg)
	fmt.Println(string(res))
}

func checkArg() {
	if (len(os.Args[1:]) != 4) {
		log.Println("args not correct")
		msg := "args not correct"
		ms, _ := json.Marshal(msg)
		fmt.Println(string(ms))
		os.Exit(1)
	}
}
