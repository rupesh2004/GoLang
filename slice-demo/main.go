package main

import "fmt"

func main() {
	fmt.Println("Slice demo")
	// Declare a slice
	slice1:= []int{1,2,3,4,5}
	fmt.Println(slice1) 
	fmt.Println(len(slice1))
	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	for _,num := range slice1 {
		fmt.Println(num)
	}

	// declare slice using make function
	fmt.Println("slice using make function")
	
	slice2 :=make([]int, 3,5)
	
	slice2 = append(slice2, 1)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 2)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 3)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 4)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 5)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 6)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 7)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))
	
	slice2 = append(slice2, 8)
	fmt.Println("Slice: ",slice2)
	fmt.Println("Length: ",len(slice2))
	fmt.Println("Capacity: ",cap(slice2))

	// declaring another slice
	fmt.Println("declaring another slice")
	slice3 := make([]int, 0)
	
	fmt.Println("Slice: ",slice3)
	fmt.Println("Length: ",len(slice3))
	fmt.Println("Capacity: ",cap(slice3))

	slice3 = append(slice3, 1)
	fmt.Println("Slice: ",slice3)
	fmt.Println("Length: ",len(slice3))
	fmt.Println("Capacity: ",cap(slice3))

	slice3 = append(slice3, 2)
	fmt.Println("Slice: ",slice3)
	fmt.Println("Length: ",len(slice3))
	fmt.Println("Capacity: ",cap(slice3))

	slice3 = append(slice3, 3)
	fmt.Println("Slice: ",slice3)
	fmt.Println("Length: ",len(slice3))
	fmt.Println("Capacity: ",cap(slice3))

}