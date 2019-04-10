package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	s := []interface{}{1, 2, 3, 4, 5}
	log.Println(s, len(s), cap(s))
	log.Println(s[5:])

	remove(&s, 5)
	log.Println(s, len(s), cap(s))
}

func remove(s *[]interface{}, value interface{}) error {
	for i, v := range *s {
		if v == value {
			log.Println(i)

			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}
	return errors.New("NOT EXITS")
}
