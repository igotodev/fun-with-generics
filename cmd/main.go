package main

import (
	"fmt"
	"fun-with-generics/pkg/collection"
)

func main() {
	listFloats := collection.NewList[float64]()
	listInts := collection.NewList[int]()
	listStrings := collection.NewList[string]()

	var f1 float64 = 22.2
	var f2 float64 = 13.2
	var f3 float64 = 10.3
	var f4 float64 = 17.5

	listFloats.Add(f1)
	listFloats.Add(f2)
	listFloats.Add(f3)
	listFloats.Add(f4)

	fmt.Println("List is empty:", listFloats.IsEmpty())
	fmt.Println("Unsorted:", listFloats.GetAll())
	listFloats.Sort()
	fmt.Println("Sorted:", listFloats.GetAll())

	listStrings.Add("Dina", "Anna", "Ilya", "Boris")
	listStrings.Sort()
	fmt.Printf("Sorted: %v\n", listStrings.GetAll())

	listInts.Add(17, 12, 5, 78, 34)
	listInts.Remove(0)
	fmt.Printf("Unsorted list: %v, length is %d\n", listInts.GetAll(), listInts.Len())

}
