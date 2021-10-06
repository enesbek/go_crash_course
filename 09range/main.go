package main

import "fmt"

func main() {
	ids := []int{23, 15, 65, 78, 99}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	//Range with map
	nameSurname := map[string]string{"Ahmet": "Ak", "Bob": "Bee"}
	for k, v := range nameSurname {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range nameSurname {
		fmt.Println("Name: ", k)
	}
}
