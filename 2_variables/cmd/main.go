package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 5
	fmt.Println(intNum)

	var str1 string = "five"
	fmt.Println(str1)

	var str2 string = `
give
  me 
    five`
	fmt.Println(str2)

	fmt.Println(len("a"))
	fmt.Println(len("Ф"))
	fmt.Println(utf8.RuneCountInString("Ф"))

	var str3 = "goll"
	str4 := "fall"
	fmt.Println(str3 + " " + str4)

	const str5 string = "can not change!"
	//str5 = "maoa"
	fmt.Println(str5)

}
