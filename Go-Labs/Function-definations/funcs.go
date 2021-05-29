/*
1. Funtion as Variable
2. Function as Argument
3. Anonymous Function
*/

//function as Vairable

package main

import (
	"fmt"
	"strings"
)

var funcvar /*eta function er name*/ func(int) int /*type*/ // var name_of_variable type_of_vairable = var name_of_function_variable func(value_type) return_value_of_func

func intfn(x int) int {
	return x + 1
}

var name_func func(string) string

func your_name(name string) string {
	return "Hello, " + name
}

var name_count func(string) int

func name(u_name string) int {
	words := strings.Fields(u_name)
	fmt.Print(words)
	count := 0
	for i, word := range words {
		fmt.Print(word, "\n")
		fmt.Print(i, "\n")
		count += 1
	}
	return count
}

// Function As a argument:

//example 1

func applyit(afunc func(int) int, val int) int { //taken function as argument
	return afunc(val)
}

func incfn(x int) int {
	return x + 1
}

func decfn(x int) int {
	return x - 1
}

// example 2

//Anonymous Function

//example 1

func applyIt(afunc func(int) int, val int) int {
	return afunc(val)
}

//Function as return value:

//Example 1:

func sqrt(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		negX := x - o_x
		negY := y - o_y
		return negX + negY
	}
	return fn //eta pura function return kortese
}

//Varibable Agumnt call

//Example:

func variable_val(vals ...int) { // ekhane multiple variable nea jabe, ... diye eta mention kora hoy, vals ekhane slice hobe
	for _, val := range vals {
		fmt.Print(val)
	}
}

//main function
func main() {
	//funcation as a argument call
	//a := applyit(incfn, 2)
	//b := applyit(decfn, 2)
	//fmt.Println(a, b)
	//annonymous function call
	//a := applyIt(func(x int) int { return x + 1 }, 2)
	//fmt.Print(a)
	//function retunr call
	//sec_func := sqrt(0, 0) // "sec_func" variable er vitore full function store hoilo ekhon "sec_func" nije e ekta function jeta er vitore fn theke return kora function store kora
	//fmt.Println(sec_func(1.2, 3.3))
	variable_val(1, 2, 3, 4) // number egula function e pass korbe and egula ekta slice r moto count hobe
	vslice := []int{5, 6, 7, 8}
	variable_val(vslice...) //multicple variable hisab e slice o passs kora jay

	//deferred call:
	//defer fmt.Print("world!") // deferred call use korle code serial onujai e execute hobe but call hobe pura main() function complete er pore
	//fmt.Print("Hello,")
	i := 1

	fmt.Print(i)

	i++

	defer fmt.Print(i + 1)

	fmt.Print(i)

}
