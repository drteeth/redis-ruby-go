package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"encoding/json"
)

type Message struct {
	User string
	Msg string
}

func connect() (redis.PubSubConn, error) {
	var conn redis.PubSubConn
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		return conn, err
	}

	conn = redis.PubSubConn{c}
	return conn, nil
}

func main() {
	// Connect
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe
	conn.Subscribe("chanlol")

	// Listen
	for {
		switch v := conn.Receive().(type) {
		case redis.Message:
			onMessage(v)
		case redis.Subscription:
			onSubscription(v)
		case error:
			log.Fatal(v)
		}
	}
}

func parseMessage(data []byte) (Message, error) {
var m Message
	err := json.Unmarshal(data, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}

func onMessage(msg redis.Message) {
	m, err := parseMessage(msg.Data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("#%s - [%s]: %s\n", msg.Channel, m.User, m.Msg)
}

func onSubscription(sub redis.Subscription) {
	fmt.Printf("%s: %s %d\n", sub.Channel, sub.Kind, sub.Count)
}
