package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	fmt.Println(fruitArr)

	//Declare and Assign
	names := [2]string{"Enes", "Ahmet"}

	fmt.Println(names)

	names2 := []string{"Enes", "Ahmet", "Mehmet"}

	fmt.Println(len(names2))
	fmt.Println(names2[1:2])
}
