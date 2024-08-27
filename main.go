package main

import (
	"log"
	"time"
)

func main() {
	teste, err := time.Parse("2006-01-02", "2024-08-27")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(teste)
}
