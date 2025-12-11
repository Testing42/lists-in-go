package main

import "fmt"

type Product struct {
	Title string
	ID    int
	Price float64
}

func main() {

	//first question
	hobbies := [3]string{"judo", "karate", "technology"}

	fmt.Println("the entire array")
	fmt.Println(hobbies)

	//second question
	firstItem := hobbies[0]

	newSlice := hobbies[1:3]

	fmt.Println("first item:", firstItem)
	fmt.Println("new slice:", newSlice)

	//third question
	firstprtSlice := hobbies[0:2]
	firstprtSliceV2 := hobbies[:2]

	fmt.Println("first half of slice:", firstprtSlice)
	fmt.Println("first half of sliceV2:", firstprtSliceV2)

	//4th slice

	newReSlice := []string{hobbies[1], hobbies[len(hobbies)-1]}

	fmt.Println("first half of slice:", newReSlice)

	//5th dynamic array

	courseGoals := []string{
		"Master Go's concurrency model (goroutines and channels).",
		"Build a fully functional REST API using the standard library.",
		"Achieve a deep understanding of interface composition.",
	}
	fmt.Println("Coursegoals before change", courseGoals)

	courseGoals[1] = "api information for golang."

	courseGoals = append(courseGoals, "sleeping golang microservices.")

	fmt.Println("Coursegoals after change", courseGoals)

	products := []Product{
		{
			Title: "Laptop",
			ID:    101,
			Price: 1200.50,
		},
		{
			Title: "Mouse",
			ID:    102,
			Price: 25.99,
		},
	}
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
