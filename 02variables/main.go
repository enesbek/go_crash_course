package main

import "fmt"

func main() {
	const name string = "Enes"

	var age = 23

	isTrue := true

	fmt.Println(name, age, isTrue)
	fmt.Printf("%T\n", name)

	var number int32 = 123

	fmt.Printf("%T\n", number)
}
