package main

import (
	"fmt"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/gitcrawler"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/server_service/server"
)

func main() {
	fmt.Println("Greetings~~")

	go gitcrawler.GetRandomUsers()

	server.StartServer()
}
