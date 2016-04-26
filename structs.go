package main

import (
	"fmt"
)

type SuperType interface {
	GetName() string
	GetMessage() string
	CallSecond()
}

type A struct{}

func (a *A) CallFirst() {
	fmt.Println("A CallFirst")
}

func (a *A) CallSecond() {
	a.callSecond(a)
}

func (a *A) callSecond(s SuperType) {
	fmt.Println(s.GetName(), s.GetMessage())
}

func (a *A) GetName() string {
	return "A"
}

func (a *A) GetMessage() string {
	return "CallSecond"
}

type B struct {
	A
}

func (b *B) CallFirst() {
	fmt.Println("B CallFirst")
}

func (b *B) GetName() string {
	return "B"
}

func (b *B) CallSecond() {
	b.callSecond(b)
}

type C struct {
	A
}

func (c *C) GetName() string {
	return "C"
}

func (c *C) CallSecond() {
	c.callSecond(c)
}

func main() {
	a := new(A)
	a.CallFirst()
	b := new(B)
	b.CallFirst()

	a.CallSecond()
	b.CallSecond()

	c := new(C)

	DoSomething(a)
	DoSomething(b)
	DoSomething(c)
}

func DoSomething(s SuperType) {
	s.CallSecond()
}
