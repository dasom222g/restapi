package main

import "fmt"

// bread
type bread struct {
	name string
}

func (b *bread) putJam(j *jam) {
	b.name += " + " + j.name
	// b.name += " + " + j.getSpoon()
}

func (b *bread) String() string {
	return b.name
}

// jam
type jam struct {
	name string
}

type jamMethods interface {
	getSpoon() string
}

// interface에서 정의한 메소드 구현
func (j *jam) getSpoon() string {
	return j.name + " get spoon"
}

func main() {
	// newBread := bread{name: "dounut"}
	// bread := &bread{name: "dounut"}
	appleJam := &jam{name: "appleJam"}
	var j jamMethods = appleJam

	// bread.putJam(appleJam)
	// fmt.Println(bread)
	fmt.Println(j.getSpoon())
}
