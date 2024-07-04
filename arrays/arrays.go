package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// "..." makes the compiler count the number of elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//while creating if an index is mentioned then the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b) //--> will return "idx: [100 0 0 400 500]
	//can create multi dimentional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	//declaring 2d array at once
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
