package main

import (
	"fmt"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/gitcrawler"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/server"
)

func main() {
	fmt.Println("Greetings~~")
	fmt.Println("Creating tables...")
	go database.CreateTables()

	go gitcrawler.GetRandomUsers()

	server.StartServer()
}
