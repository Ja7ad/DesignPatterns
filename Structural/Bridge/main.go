package main

import "fmt"

type IDrawShape interface {
	drawshape(x[5] float32, y[5] float32)
}

type DrawShape struct {}

/*
The drawShape method draws the shape given the coordinates
 */
func (drawShape DrawShape) drawshape(x[5] float32, y[5] float32)  {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x[5] float32, y[5] float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

/*
The drawContour method of the DrawContour class calls the drawShape method on the
shape instance
 */
func (contour DrawContour) drawContour(x[5] float32, y[5] float32)  {
	fmt.Println("Drawing Contour")
	contour.shape.drawshape(contour.x, contour.y)
}

func (contour DrawContour) resizeByFactor(factor int)  {
	contour.factor = factor
}

func main() {
	var x = [5]float32{1,2,3,4,5}
	var y = [5]float32{1,2,3,4,5}

	var contour IContour = DrawContour{x,y,DrawShape{},2}
	contour.drawContour(x,y)
	contour.resizeByFactor(2)
}