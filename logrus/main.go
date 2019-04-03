package main

import (
	"strconv"
	"time"
)

func main() {

	log := NewLoggerWithRotate()

	for i := 0; i < 180; i++ {
		log.Infoln("this is a info log , id : " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	log.Panicln("this is a panic log!")
}
