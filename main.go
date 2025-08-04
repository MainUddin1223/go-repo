package main

import (
	"fmt"
	// "mathlib/mathlib"
)
var a = 100
// func main(){
// 	var a int = 11
// 	fmt.Println(a)
// }
func main(){
// var result int = AddNumber(a,6)
if a == 100{
	a=50

	fmt.Println("if",a,)
}
fmt.Println("out side if",a,)

add := func(a int , b int){
	c:= a+b
	fmt.Println(c)
	fmt.Println(c)
}
add(10,20)
// var ages int = mathlib.AddAges(10,100)
}
// func init(){
// // var result int = AddNumber(a,6)
// var ages int = mathlib.AddAges(10,100)
// fmt.Println("ages init111111111111",ages,)
// }