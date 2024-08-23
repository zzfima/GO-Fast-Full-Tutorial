package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (ge gasEngine) milesLeft() uint16 {
	var v uint16 = uint16(ge.gallons) * uint16(ge.mpg)
	return v
}

func milesLeft1(ge gasEngine) uint16 {
	var v uint16 = uint16(ge.gallons) * uint16(ge.mpg)
	return v
}

func main() {
	var myGasEngine gasEngine
	myGasEngine.gallons = 45
	myGasEngine.mpg = 21
	fmt.Println(myGasEngine)

	gasEngine1 := gasEngine{mpg: 19, gallons: 55}
	fmt.Println(gasEngine1)

	fmt.Println("miles left: ", gasEngine1.milesLeft())
	fmt.Println("miles left: ", milesLeft1(gasEngine1))

	//anonymous struct
	var myElecEngine = struct {
		kv       uint8
		capacity uint16
	}{kv: 8, capacity: 40000}
	fmt.Println(myElecEngine)
}
