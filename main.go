package main

import (
	"mysort"
)

func main() {
	// fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	var slice []int
	slice = []int{1, 3, 3, 11, 5, 6}
	slice = mysort.BubbleSort(slice)
	// fmt.Printf("Hello, world.  bubblesort = %v\n", slice)

	// http.HandleFunc("/", Redirect)
	// http.HandleFunc("/add", Add)
	// http.ListenAndServe(":8080", nil)
}
