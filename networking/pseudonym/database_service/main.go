package main

import (
	"fmt"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database_service/server"
)

func main() {
	fmt.Println("Greetings~~")

	server.StartServer()
}
