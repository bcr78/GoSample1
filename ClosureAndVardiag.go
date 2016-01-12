package main

import (
	"fmt"
)

const Pi float64 = 3.14

//Vardiac function
func sum(Name string, nums ...int) {
	fmt.Println("name = ", Name)
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func makeEvenGenerator(i int) func(i int) int {

	return func(i int) (ret int) {
		ret = i + 2
		return
	}
}

func main() {
	sum("AAA", 1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum("BBB", nums...)

	nextEven := makeEvenGenerator(2)
	fmt.Println(nextEven(2)) // 4
	fmt.Println(nextEven(3)) // 5
	fmt.Println(nextEven(4)) // 6

	fmt.Printf("%T", Pi)
	//fmt.printf("",&Pi) // Compile time error--> cannot take the address of Pi
}

/*
name =  AAA
[1 2 3 4 5] 15
name =  BBB
[1 2 3 4] 10
4
5
6
float64
[Finished in 0.334s]
*/
