package main

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello world")

	log := NewLogger()

	for i := 0; i < 10; i++ {
		log.Infoln("this is a info log , id : " + strconv.Itoa(i))
	}

	log.Panicln("this is a panic log!")
}
