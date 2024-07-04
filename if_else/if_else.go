package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//go supports if without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 7 or 8 is even")
	}
	//any variables declared in the preceding statement of a condition is available in current and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	//num since defined in the if is not available outside that if
	//fmt.Println(num) //--> throws undefined error
}
