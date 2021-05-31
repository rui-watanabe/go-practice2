package main

import "fmt"

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

func main() {
	localVariable()
	fmt.Println(i, f64, s, t, f)

	f := Later()
	fmt.Println(f("hello"))
	fmt.Println(f("hey"))
	fmt.Println(f("foo"))
	fmt.Println(f("hoge"))

}
