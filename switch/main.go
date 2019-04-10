package main

import (
	"log"
)

func main() {
	i := 10
	switch i {
	case 0:
		log.Println("0")
	case 1, 10:
		log.Println("1")
	case 2, 20:
		log.Println("2")
	}

	switch {
	case i < 10:
		log.Println("<10")
	case i == 10, i > 10:
		log.Println(">=10")
	}
}
