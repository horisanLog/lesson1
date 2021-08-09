package main

import (
	"fmt"
)

/*
変数練習
*/
var i3 int = 500

func outer() {
	var s4 string = "test"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "hello"
	fmt.Println(s)
	//暗黙的な定義
	//変数名 := 値
	i2 := 400
	fmt.Println(i2)

	i2 = 500
	fmt.Println(i2)

	// i2 := 30 再度暗黙的な定義はできない
	// fmt.Println(i2)
	fmt.Println(i3)

	outer()
}
