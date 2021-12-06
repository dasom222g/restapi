package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	String() string
}

// 인자를 interface로 받으므로써 struct타입 신경쓰지 않고 interface의 메소드가 구현된 값인지만 확인
// object의 속성타입은 신경쓰지 않으므로 struct와의 종속성 끊음
func Println(p Printable) {
	fmt.Println(p.String())
}

// StructA
type StructA struct {
	val string
}

func (a *StructA) String() string {
	return "A val: " + a.val
}

// StructB
type StructB struct {
	val int
}

func (b *StructB) String() string {
	return "B val: " + strconv.Itoa(b.val)
}

func main() {
	a := &StructA{val: "aaa"}
	Println(a)

	b := &StructB{val: 4000}
	Println(b)
}
