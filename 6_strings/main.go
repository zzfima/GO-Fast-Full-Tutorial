package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(myString)

	//letter é representation in UTF8: 232
	//so its in range 128-255
	//so need 2 bytes
	//put the number 232 = 11101000 into pattern for 2 bytes: 110xxxxx 10xxxxxx (see utf8 encoding attached image)
	//lets pad 11101000 by 3 zeroes from start because we have 11 x: 00011101000
	//we get 11000011 and 10101000
	//first is 195 second is 169

	var someLetter1 = myString[1]
	fmt.Printf("%v %T\n", someLetter1, someLetter1)
	var someLetter2 = myString[2]
	fmt.Printf("%v %T\n", someLetter2, someLetter2)

	for i1, v1 := range myString {
		fmt.Println(i1, v1)
	}

	fmt.Println("Is num of bytes and not a characters: ", len(myString))

	//runes is unicode point number which represents characters
	var myRune = []rune("résumé")
	var someRune1 = myRune[1]
	fmt.Printf("%v %T\n", someRune1, someRune1)
	for i2, v2 := range myRune {
		fmt.Println(i2, v2)
	}

	//strings are immutable
	//myString[1] = "h"

	//string builder
	var strBld strings.Builder
	var myString1 = "mama"
	for i := range myString1 {
		strBld.WriteByte(myString1[i])
	}
	fmt.Println("Rune format: ", strBld)
	fmt.Println("String format: ", strBld.String())
}
