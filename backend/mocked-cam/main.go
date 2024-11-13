package main

import (
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
