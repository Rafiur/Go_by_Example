package main

import "fmt"

func main() {

	i := 1
	//conventional
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	//using only range
	for i := range 3 {
		fmt.Println("Range :", i)
	}
	//without condition loop is an infinte loop
	for {
		fmt.Println("loop")
		//hence we call break to break from the loop
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			//continue can be used to skip through
			//for example all even numbers are skipped through and not printed
			continue
		}
		fmt.Println(n)
	}

}
