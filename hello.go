package main

import "fmt"

func init() {
	fmt.Println("init")
}

func hoge() {
	fmt.Println("hoge")
}

func main() {
	fmt.Println("hello world")
	hoge()
}
