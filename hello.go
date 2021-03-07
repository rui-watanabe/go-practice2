package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("init")
}

func hoge() {
	fmt.Println("hoge", time.Now())
	fmt.Println(user.Current)
}

func main() {
	fmt.Println("hello world")
	hoge()
}
