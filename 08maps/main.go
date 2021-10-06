package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)

	//Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Ak"] = "ak@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//Delete
	delete(emails, "Ak")
	fmt.Println(emails)

	//Declare map and key values
	nameSurname := map[string]string{"Ahmet": "Ak", "Bob": "Bee"}
	fmt.Println(nameSurname)
}
