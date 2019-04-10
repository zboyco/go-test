package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner()
}

func char() {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, _ := reader.ReadRune()
		fmt.Println(char)
	}
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
