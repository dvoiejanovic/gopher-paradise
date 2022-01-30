package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{2, 3}
	v.Scale(10)
	ScaleFunc(&v, 10)

	p := Vertex{4, 3}
	p.Scale(5)
	ScaleFunc(&p, 10)

	fmt.Println(v, p)
}
