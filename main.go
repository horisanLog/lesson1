package main

import (
	"fmt"
	"strconv"
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
	i2 := 3.4
	fmt.Println(i2)

	i2 = 5.2
	fmt.Println(i2)

	// i2 := 30 再度暗黙的な定義はできない
	// fmt.Println(i2)
	fmt.Println(i3)

	outer()

	// 型をしらべることができる
	fmt.Printf("%T\n", i2)

	// floatで型変換
	var fl64 float64 = 2.4
	fmt.Printf("%T, %T\n", fl64, i2)

	// ゼロで割る
	zero := 1.0
	pinf := -1.0 / zero
	fmt.Println(pinf)

	// 配列型
	var a1 [3]int
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	var a2 [3]string = [3]string{"A", "B"}
	fmt.Println(a2)

	a3 := [3]int{1, 2, 3}
	fmt.Println(a3)

	a4 := [...]string{"C", "D"}
	fmt.Println(a4)
	fmt.Printf("%T\n", a4)

	fmt.Println(a4[0])
	fmt.Println(a2[1])
	a2[2] = "C"
	fmt.Println(a2)
	// 配列の要素数は変えられない
	// a2[3] = "g"
	// fmt.Println(a2)

	// interface すべての値を格納できる
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = "a"
	fmt.Println(x)

	// 型変換
	var it int = 1
	f64 := float64(it)
	fmt.Println(f64)
	fmt.Printf("it = %T\n", it)
	fmt.Printf("f64 = %T\n", f64)

	it2 := int(f64)
	fmt.Printf("it2 = %T\n", it2)

	var st string = "100"
	fmt.Printf("st = %T\n", st)

	// 数値変換
	sti, _ := strconv.Atoi(st)
	fmt.Println(sti)
	fmt.Printf("it2 = %T\n", sti)

	var int2 int = 200
	sti2 := strconv.Itoa(int2)
	fmt.Println(sti2)
	fmt.Printf("sti2 = %T\n", sti2)

	// 論理値型
	t, f := true, false

	// T : 型、 v :データ、 t\n : 結果
	// bool 1 true
	// bool 0 false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	// map
	m := map[string]int{"a": 100, "b": 200}
	fmt.Println(m)
	v, ok := m["a"]
	fmt.Println(v, ok)

	// 1.11をint型に変換

	u := 1.11
	reslut := int(u)
	fmt.Println(reslut)

	// fmt.Printf("%T %v", m, m)
	se := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}

	fmt.Printf("%T %v", se, se)

	// map[string]int map[Mike:20 Nancy:24 Messi:30]

}
