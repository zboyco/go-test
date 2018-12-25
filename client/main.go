package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("start")
	for i := 0; i < 100; i++ {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9043")
		checkErr(err)
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		checkErr(err)
		defer conn.Close()

		var headSize int
		var headBytes = make([]byte, 4)
		headBytes[0] = '$'
		headBytes[3] = '#'
		s := "hello world"
		content := []byte(s)
		headSize = len(content)
		binary.BigEndian.PutUint16(headBytes[1:], uint16(headSize))
		conn.Write(headBytes)
		conn.Write(content)

		s = "hello go"
		content = []byte(s)
		headSize = len(content)
		binary.BigEndian.PutUint16(headBytes[1:], uint16(headSize))
		conn.Write(headBytes)
		conn.Write(content)

		s = "hello tcp"
		content = []byte(s)
		headSize = len(content)
		binary.BigEndian.PutUint16(headBytes[1:], uint16(headSize))
		conn.Write(headBytes)
		conn.Write(content)
		// conn.Close()
		// time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 120)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
