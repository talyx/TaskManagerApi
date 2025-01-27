package main

import (
	"fmt"
	"github.com/talyx/TaskManagerApi/internal/app"
	"github.com/talyx/TaskManagerApi/internal/utils"
	"log"
)

func main() {
	password := "986532"
	hashedPassword, _ := utils.HashPassword(password) // Хэш из базы данных

	err := utils.ComparePassword(hashedPassword, password)
	if err != nil {
		fmt.Println("Password does not match:", err)
	} else {
		fmt.Println("Password matches!")
	}
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
