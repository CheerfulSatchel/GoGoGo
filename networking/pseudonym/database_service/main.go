package main

import (
	"fmt"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/databvase_service/database"
)

func main() {
	fmt.Println("Greetings~~")

	database.GetRandomUsers()
}
