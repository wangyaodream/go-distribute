package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id    int
	name  string
	age   int
	score float32
}

func fillInfo(s *student) *student {
	s.id = 20
	s.age = 18
	s.name = "wangyao"
	s.score = 96.5
	return s
}

func main() {

	a := student{}
	tmp := fillInfo(&a)

	fmt.Println(reflect.TypeOf(a))
	result := reflect.TypeOf(&tmp)
	fmt.Println(result)

}
