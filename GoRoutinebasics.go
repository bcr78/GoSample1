package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	f("main Routine 1")
	go f("GoRoutine")
	f("main Routine 2")

	go func(msg string) {
		fmt.Println(msg)
	}("Direct GoRoutine call")

	intChnl := make(chan int)
	go func() {
		intChnl <- 5
	}()

	fmt.Println("value from channel = ", <-intChnl)
	var input string
	fmt.Scanln(&input)
}
