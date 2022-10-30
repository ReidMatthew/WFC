package main

import (
	"fmt"
)

type Vector struct {
	x int
	y int
}
type Tile struct {
	border [][]int // side patern clockwise N:0 E:1 S:2 W:3
	img    string
}

type Cell struct {
	border    [][]int // side patern clockwise N:0 E:1 S:2 W:3
	tileSet   []Tile  // all possible tiles
	pix       int     // pixilation of edge
	collapsed bool    // solved
	pos       Vector  // position on grid
}

func gridInit(n int) string {
	return "a"
}

func main() {

	fmt.Println(gridInit(3))
}
