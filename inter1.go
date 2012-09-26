package main

import "fmt"

type S struct {i int}

func (p *S) Get() int {return p.i}

func (p *S) Put(v int) {p.i=v}

type R struct {i int}

func (p *R) Get() int {return p.i}

func (p *R) Put(v int) {p.i=v}

type Interface interface {

	Get() int
	Put(int)

}

func f(p Interface) {

	fmt.Println(p.Get())
	p.Put(1)

}

func main() {

	s := &S{i: 5}
	f(s)
	r := &R{i: 4}
	f(r)

}

