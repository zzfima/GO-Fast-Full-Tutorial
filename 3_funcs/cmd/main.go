package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(getName())

	printMsg("my message")

	printMsg(concatinator("hello ", "work"))

	var i1, i2 int = addSub(6, 3)
	fmt.Printf("%v, %v\n", i1, i2)

	i3, err1 := div(6, 1)
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println("no errors!")
	}
	fmt.Printf("%v\n", i3)

	i4, err2 := div(6, 0)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println("no errors!")
	}
	fmt.Printf("%v\n", i4)

	i2 = -9
	if i1 > 0 && i2 > 0 {
		fmt.Println("all nums are positive")
	} else if i1 > 0 {
		fmt.Println("only i1 is positive")
	} else if i2 > 0 {
		fmt.Println("only i2 is positive")
	} else {
		fmt.Println("no positive nums")
	}

	switch {
	case i1 > 0 && i2 > 0:
		fmt.Println("all nums are positive")
	case i1 > 0:
		fmt.Println("only i1 is positive")
	case i2 > 0:
		fmt.Println("only i2 is positive")
	default:
		fmt.Println("no positive nums")
	}

	switch i1 {
	case 0:
		fmt.Println("i1 is 0")
	case 1, 2:
		fmt.Println("i1 is 1 or 2")
	default:
		fmt.Println("i1 isn't 0 or 1 or 2")
	}
}

func getName() string {
	return "Vasa"
}

func printMsg(msg string) {
	fmt.Println(msg)
}

func concatinator(str1 string, str2 string) string {
	return str1 + str2
}

func addSub(int1 int, int2 int) (int, int) {
	return int1 + int2, int1 - int2
}

func div(int1 int, int2 int) (int, error) {
	var err error

	if int2 == 0 {
		err = errors.New("can not divide by zero")
		return 0, err
	}
	return int1 / int2, err
}
