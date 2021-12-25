package main

import "fmt"

type Component interface {
	Operation(string)
}

// 메인 기능
type SendComponent struct{}

// 부가 기능
type ZipComponent struct{}

type EncrytComponent struct{}

func main() {
	fmt.Println("Decorator")
}
