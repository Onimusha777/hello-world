package main //package это пакет, содержащий код. Чтобы достать код - нужно сослаться на пакет

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(1 + 111)
	fmt.Println(true)

	//declaration
	var name string

	//declaration and initialization
	name2 := "Konstantin" //recommended
	var name3 string = "Ivan"

	//initialization
	name = "Sergei"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)

}

func main2() {
	var a bool = true    //Boolean
	var b int = 5        //Integer
	var c float32 = 3.14 //floating point number
	var d string = "пруикаукари"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main3() {
	//+
	//-
	//*
	// /
	// %

	a := 10
	b := 31

	result := a + b
	fmt.Println(result)

}
