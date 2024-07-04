package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	}

	switch time.Now().Weekday() {
	//can use commas to separate multiple expresions in the same case statement
	case time.Friday, time.Saturday:
		fmt.Println("Its the Weekend")
	default:
		fmt.Println("Its a Weekday")
	}

	t := time.Now()
	//switch without an expression
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}
	// a type switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
	whatAmI(2e9)
}
