package main

import "fmt"

func main() {
	var myMap1 map[string]int32 = map[string]int32{"a": 1, "b": 2}
	fmt.Println(myMap1)

	var myMap2 map[string]int32 = make(map[string]int32)
	myMap2["G"] = 22
	fmt.Println(myMap2)
	//no error - return 0
	fmt.Println(myMap2["K"])

	//check if exists
	var code, isExit = myMap2["K"]
	fmt.Printf("value: %v, isExists: %v\n", code, isExit)

	//delete value
	myMap2["K"] = 33
	myMap2["L"] = 356
	fmt.Println(myMap2)
	delete(myMap2, "K")
	fmt.Println(myMap2)

	//iterate map
	for k1, v1 := range myMap2 {
		fmt.Println(k1, v1)
	}

	//iterate array
	var myArr1 []int32 = []int32{5, 6, 7}
	for i1, v1 := range myArr1 {
		fmt.Println(i1, v1)
	}
}
