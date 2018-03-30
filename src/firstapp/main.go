package main

import (
	"bufio"
	"errors"
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

func sumOf1ToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sumOf1ToNMulitplesOf3And5(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func largetNum(arr []int) (int, error) {
	largest := -1000000
	for item := range arr {
		if item > largest {
			largest = item
		}
	}
	err := func() error {
		if largest == -1000000 {
			return errors.New("errror")
		} else {
			return nil
		}
	}()
	return largest, err
}
func main() {
	// helloWorld()
	// askAndGreet()
	// askAndGreetAliceAndBob()
	// fmt.Println(sumOf1ToN(3))
	// fmt.Println(sumOf1ToNMulitplesOf3And5(9))
	fmt.Println(largetNum([]int{1, 2, 3, 4, 5}))
}
