package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	//Short method
	for j := 1; j <= 10; j++ {
		fmt.Print(j)
	}

	fmt.Println()

	//FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
