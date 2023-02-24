package main

import "fmt"

func main() {

	var (
		name string
		age  int
	)
	fmt.Print("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Your input is:", name)
	fmt.Print("Enter your age:")
	fmt.Scan(&age)
	fmt.Println("Your input is:", age)
	fmt.Println("Your name is:", name, "and your age is:", age)

}
