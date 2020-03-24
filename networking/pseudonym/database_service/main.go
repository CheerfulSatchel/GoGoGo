package main

import (
	"fmt"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database_service/database"
)

func main() {
	fmt.Println("Greetings~~")

	database.CreateTables()
}
