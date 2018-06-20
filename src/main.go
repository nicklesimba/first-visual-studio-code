package main

import "fmt"

func main() {
	fmt.Println("Welcome to Visual Studio Code!")
	//now a random calculation to test how the debugger works...
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
