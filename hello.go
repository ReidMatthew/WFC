package main

import "fmt"

type tile struct {
	border [][]int // side patern clockwise N:0 E:1 S:2 W:3
}

type cell struct {
	border    [][]int // side patern clockwise N:0 E:1 S:2 W:3
	pix       int     // pixilation of edge
	collapsed bool    // solved
}

// func newPerson(name string) *person {

// 	p := person{name: name}
// 	p.age = 42
// 	return &p
// }

func main() {
	fmt.Println("a")
}
