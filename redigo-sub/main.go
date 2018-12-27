package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	log.Println("测试redis发布订阅")

	conn, err := redis.Dial("tcp", "192.168.2.99:6379")
	checkErr(err)
	defer conn.Close()

	go subScribe()
	go subScribe()
	go subScribe()

	for {
		var s string
		fmt.Scanln(&s)
		_, err := conn.Do("PUBLISH", "testRoom", s)
		if err != nil {
			fmt.Println("pub err: ", err)
			return
		}
	}
}

func subScribe() {
	conn, err := redis.Dial("tcp", "192.168.2.99:6379")
	checkErr(err)
	defer conn.Close()

	psc := redis.PubSubConn{
		Conn: conn,
	}
	psc.Subscribe("testRoom")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			log.Fatalln(v)
		}
	}
}

// checkErr 错误检查
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
