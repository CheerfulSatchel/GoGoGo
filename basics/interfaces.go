package main

import (
	"fmt"
	"time"
)

type Receipe interface {
	BakeTime() int
}

type CookingError struct {
	What string
	When time.Time
}

type Food struct {
	Name            string
	CookTimeMinutes int
}

func (food *Food) BakeTime() int {
	food.CookTimeMinutes = 22
	return food.CookTimeMinutes
}

func (e *CookingError) Error() string {
	return fmt.Sprintf("Time: %v, Reason: %v", e.When, e.What)
}

func (f *Food) String() string {
	return fmt.Sprintf("Yummy yum yum, it's %v and takes %v minutes to cook", f.Name, f.CookTimeMinutes)
}

func testError() error {
	return &CookingError{
		What: "The stove was never turned off!",
		When: time.Now(),
	}
}

func CookFood(receipe Receipe) {
	receipe.BakeTime()
	fmt.Println("All cooked up~")
}

func main() {
	food := Food{
		Name:            "Soup",
		CookTimeMinutes: 5,
	}
	fmt.Println(&food)
	// fmt.Println("Yay " + food.String())

	if err := testError(); err != nil {
		fmt.Printf("%T\n, %v", err, err)
	}

	fmt.Printf("Before, food's cooking time is: %v\n", food.CookTimeMinutes)
	CookFood(&food)
	fmt.Printf("Now, food's cooking time is: %v\n", food.CookTimeMinutes)
}
