package main

import (
	"fmt"
	"mathlib/mathlib"
)

// func main(){
// 	var a int = 11
// 	fmt.Println(a)
// }
func main(){
// var result int = AddNumber(a,6)
var ages int = mathlib.AddAges(10,100)
fmt.Println("main",ages,)
}
func init(){
// var result int = AddNumber(a,6)
var ages int = mathlib.AddAges(10,100)
fmt.Println("ages init111111111111",ages,)
}