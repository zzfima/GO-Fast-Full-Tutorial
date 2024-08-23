package main

import "fmt"

func main() {
	var p1 *int32 = nil
	p1 = new(int32)
	fmt.Printf("pointer value: %v\n", *p1)
	fmt.Printf("pointer address: %v\n", p1)

	*p1 = 10
	fmt.Printf("pointer value: %v\n", *p1)
	fmt.Printf("pointer address: %v\n", p1)

	var v1 int32 = 45
	fmt.Printf("var address: %v\n", &v1)
	p1 = &v1
	fmt.Printf("pointer value: %v\n", *p1)
	fmt.Printf("pointer address: %v\n", p1)
}
