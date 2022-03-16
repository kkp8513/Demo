package main

import (
	"fmt"
)

type Api interface {
	println()
}

type s1 struct {}

func (s *s1) println() {
	fmt.Println("s1")
}

type s2 struct {}

func (s *s2) println() {
	fmt.Println("s2")
}

func newApi(i int) Api {
	switch {
	case i == 1:
		return &s1{}
	case i == 2:
		return &s2{}
	default:
		return nil
	}
}

func main() {
	s := newApi(1)
	s.println()

	s2 := newApi(2)
	s2.println()
}