package main

import (
	"github.com/talyx/TaskManagerApi/internal/app"

	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
