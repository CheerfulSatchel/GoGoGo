package main

import (
	"fmt"
)

func main() {
	food := Food{
		Name:            "Soup",
		CookTimeMinutes: 5,
	}
	fmt.Println(&food)

	if err := testError(); err != nil {
		fmt.Printf("%T\n, %v", err, err)
	}

	fmt.Printf("Before, food's cooking time is: %v\n", food.CookTimeMinutes)
	CookFood(&food)
	fmt.Printf("Now, food's cooking time is: %v\n", food.CookTimeMinutes)

	goodResponseBytes, goodResponseErr := GetResponse(true)

	if goodResponseErr != nil {
		fmt.Printf("ERROR: %v", goodResponseErr)
	} else {
		fmt.Printf("%v", string(goodResponseBytes))
	}

	badResponseBytes, badResponseErr := GetResponse(false)

	if badResponseErr != nil {
		fmt.Printf("ERROR: %v", badResponseErr)
	} else {
		fmt.Printf("%v", string(badResponseBytes))
	}

}
