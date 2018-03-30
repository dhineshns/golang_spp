package main

import (
	"bufio"
	"fmt"
	"os"
)

func helloWorld() {
	fmt.Printf("Hello")
}

func askAndGreet() {
	fmt.Printf("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Hello " + input)
}

func askAndGreetAliceAndBob() {
	fmt.Println("Please enter your name : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "Alice" || input == "Bob" {
		fmt.Println("Hello " + input)
	}
}

func main() {
	helloWorld()
	askAndGreet()
	askAndGreetAliceAndBob()
}
