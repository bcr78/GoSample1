package main

import (
	"fmt"
)

func Sum(a *[3]int) (sum int) {
	for i, v := range *a {
		a[i] = 10
		sum += v
	}
	fmt.Println(a)
	return
}

func Sum1(a [3]int) (sum int) {
	for i, v := range a {
		a[i] = 10
		sum += v
	}
	fmt.Println(a)
	return
}

func main() {
	fmt.Println("hello World")
//edit1
//edit2
	//arrays
	// 3 ways to declare arrays
	//	var x [5]float64
	//x[4=100]
	//	a := [2]string{"hello", "world!"}
	//	a := [...]string{"hello", "world!"}

	//	x := [6]float64{1, 2, 3, 4, 5, 6}

	x := [6]float64{
		10,
		10,
		10,
		10,
		10,
		10,
	}
	fmt.Println(x)

	var total float64 = 0
	//Range
	for i, _ := range x {
		total += x[i]
	}
	fmt.Println("total = ", total)
	/*
		for _, value := range x {
			total += value
		}
	*/
	fmt.Println("Avg=", total/float64(len(x))) //typecasting from int to float64

	//Slices --> unlike arrays, it allows to change the size
	var nilSlice []float64 //Nil slice -> cap=0, len=0
	fmt.Printf("nil Slice=%v len=%d cap=%d\n", nilSlice, len(nilSlice), cap(nilSlice))
	s := make([]float64, 5, 10) // 5-> length of the slice, 10-> capacity of the array
	fmt.Printf("s=%v len=%d cap=%d\n", s, len(s), cap(s))

	arr := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[0:5]
	fmt.Println("s1=", s1)

	//slice function - append
	s = arr[:]
	fmt.Println(s)
	s = append(s, 11, 12, 13)
	fmt.Printf("slice= %v	len=%d	cap=%d\n", s, len(s), cap(s))
	s = append(s, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Printf("slice= %v	len=%d	cap=%d\n", s, len(s), cap(s))

	//slice function - copy
	s3 := []int{1, 2, 3, 4, 5}
	s4 := make([]int, 2, 3) //
	copy(s4, s3)            //destination, src
	fmt.Printf("s3=%v s4=%v\n", s3, s4)

	// passing array to a function, its always call by value
	array := [...]int{5, 5, 5}
	a1 := Sum1(array)
	fmt.Println(array, a1) //[7 8.5 9.1] 24.6
	a2 := Sum(&array)      // Note the explicit address-of operator
	fmt.Println(array, a2) //[10 10 10] 24.6
}

/*
output:
hello World
[10 10 10 10 10 10]
total =  60
Avg= 10
nil Slice=[] len=0 cap=0
s=[0 0 0 0 0] len=5 cap=10
s1= [1 2 3 4 5]
[1 2 3 4 5 6 7 8 9 10]
slice= [1 2 3 4 5 6 7 8 9 10 11 12 13]	len=13	cap=20
slice= [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21]	len=21	cap=40
s3=[1 2 3 4 5] s4=[1 2]
[10 10 10]
[5 5 5] 15
&[10 10 10]
[10 10 10] 15
*/
