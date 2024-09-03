package main

import (
	"log"
	"slogv2/src/main/entity"
)

func main() {
	entity.DbInit()

	log.Println("END")
}
