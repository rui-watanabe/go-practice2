package main

import (
	"fmt"
	"os"
	"strconv"
)

//関数内で宣言が可能
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "hoge"
	t, f bool    = true, false
)

func localVariable() {
	//関数内のみ宣言が可能
	xi := 2
	xi = 100               //再代入
	var xf32 float32 = 1.3 //宣言をしたい場合はvarを使う
	xs := "hogehoge"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func Generator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func strConv() {
	// var s string = "-42"
	var s string = "A"

	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}

func mapRoop() {
	m := map[string]int{"apple": 100, "banana": 200}

	for _, v := range m {
		fmt.Println(v)
	}
}

func anything(a interface{}) {
	fmt.Println(a)
}

func switchFunc() {
	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "I don't know")
	}
}

func deferFunc() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
}

func main() {
	localVariable()
	fmt.Println(i, f64, s, t, f)

	f := Later()
	fmt.Println(f("hello"))
	fmt.Println(f("hey"))
	fmt.Println(f("foo"))
	fmt.Println(f("hoge"))

	countGenerator := Generator()
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())

	strConv()

	mapRoop()

	switchFunc()

	deferFunc()
}
