package main

import (
	"fmt"
	// "os"
	"strconv"
	"testGo/alib"
	"testGo/foo"
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
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// defer file.Close()
}

func double(i *int) {
	*i = *i * 2
}

func pointerFunc() {
	var n int = 300
	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 200
	fmt.Println(n)

	double(p)
	fmt.Println(n)

}

type Point struct {
	A string
	B int
	C bool
}

type BigPoint struct {
	Point
}

func Update(p *Point) {
	p.A = "test"
	p.B = 1
	p.C = true
}

func (p *Point) Set(i int) {
	p.B = i
}

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func structFunc() {
	p := Point{"hoge", 0, false}
	fmt.Println(p)
	p2 := &Point{}
	Update(p2)
	fmt.Println(p2)
	Update(&p)
	fmt.Println(p)

	p3 := &Point{}
	p3.Set(2)
	fmt.Println(*p3)

	p4 := BigPoint{}
	p4.Point.Set(2)
	fmt.Println(p4)

	p5 := Person{"Ko"}
	p6 := Person{"Jon"}
	p7 := Person{"Mike"}

	ps := Persons{}
	ps.Persons = append(ps.Persons, &p5, &p6, &p7)

	for _, p := range ps.Persons {
		fmt.Println(p)
	}
}

type Stringfy interface {
	ToString() string
}

type Man struct {
	Name string
	Age  int
}

func (m *Man) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", m.Name, m.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type MyErr struct {
	Message   string
	ErrorCode int
}

func (e *MyErr) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyErr{Message: "occurred error", ErrorCode: 1234}
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v,%v, %v>>", p.A, p.B, p.C)
}

func interfaceFunc() {
	vs := []Stringfy{
		&Man{Name: "jon", Age: 32},
		&Car{Number: "1111", Model: "A-1"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	err := RaiseError()
	fmt.Print(err.Error())

	e, ok := err.(*MyErr)
	if ok {
		fmt.Println(e.ErrorCode)
	}

	p := &Point{"abc", 1234, true}
	fmt.Println(p)
}

func devideFileFunc() {
	fmt.Println(foo.Max)
	fmt.Println(foo.ReturnMin())
}

func scopeFunc(s string) (b string) {
	b = s
	{
		b := "bbb"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
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

	countGenerator := Generator()
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())
	fmt.Println(countGenerator())

	strConv()

	mapRoop()

	switchFunc()

	deferFunc()

	pointerFunc()

	structFunc()

	interfaceFunc()

	devideFileFunc()

	scopeFunc("AAA")

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
