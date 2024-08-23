package main

import "fmt"

func main() {

	//array - fixed length, same type, indexable, contiguous in memory
	var intArr [3]int32
	intArr[0] = 12
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])

	//addresses are contiguous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{4, 5, 6}
	fmt.Println(intArr2[0:3])

	intArr3 := [3]int32{14, 15, 16}
	fmt.Println(intArr3[0:3])

	intArr4 := []int32{14, 15, 16, 17, 18}
	fmt.Println(intArr4[0:5])
	fmt.Println(intArr4)

	//slices
	fmt.Printf("capacity before %v, len before %v\n", cap(intArr4), len(intArr4))
	intArr4 = append(intArr4, 77)
	fmt.Printf("capacity after %v, len after %v\n", cap(intArr4), len(intArr4))
	fmt.Println(intArr4)

	//spread operator ...
	var intSlice1 []int32 = []int32{3, 5, 6}
	var intSlice2 []int32 = []int32{8, 9}
	intSlice2 = append(intSlice2, intSlice1...)
	fmt.Println(intSlice2)
	fmt.Printf("capacity %v, len %v\n", cap(intSlice2), len(intSlice2))

	//create array using make
	var intSlice3 []int32 = make([]int32, 3)
	fmt.Printf("capacity %v, len %v\n", cap(intSlice3), len(intSlice3))
	var intSlice4 []int32 = make([]int32, 3, 7)
	fmt.Printf("capacity %v, len %v\n", cap(intSlice4), len(intSlice4))
}
