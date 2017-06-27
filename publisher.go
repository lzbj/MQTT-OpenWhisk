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
	arg1 := os.Args[1]
	var obj map[string]interface{}
	json.Unmarshal([]byte(arg1), &obj)
	url, ok := obj["url"].(string)
	if !ok {
		url = "tcp://localhost:1883"
	}

	arg2 := os.Args[2]
	json.Unmarshal([]byte(arg2), &obj)
	clientID, ok := obj["clientID"].(string)
	if !ok {
		clientID = "publisher"
	}

	arg3 := os.Args[3]
	json.Unmarshal([]byte(arg3), &obj)
	message, ok := obj["message"].(string)
	if !ok {
		message = "world!"
	}

	arg4 := os.Args[4]
	json.Unmarshal([]byte(arg4), &obj)
	topic, ok := obj["topic"].(string)
	if !ok {
		topic = "topic/sample"
	}
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(url).SetClientID(clientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	msg := map[string]string{"msg": ("Hello, " + message + "!")}
	res, _ := json.Marshal(msg)
	token := c.Publish(topic, 0, false, res)
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
	}
}
