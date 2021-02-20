# Birdge Design Pattern
Bridge decouples the implementation from the abstraction. The abstract base class can be
subclassed to provide different implementations and allow implementation details to be
modified easily. The interface, which is a bridge, helps in making the functionality of
concrete classes independent from the interface implementer classes. The bridge patterns
allow the implementation details to change at runtime.

The bridge pattern demonstrates the principle, preferring composition over inheritance. It
helps in situations where one should subclass multiple times orthogonal to each other.
Runtime binding of the application, mapping of orthogonal class hierarchies, and platform
independence implementation are the scenarios where the bridge pattern can be applied.

The bridge pattern components are abstraction, refined abstraction, implementer, and
concrete implementer. Abstraction is the interface implemented as an abstract class that
clients invoke with the method on the concrete implementer. Abstraction maintains a hasa relationship with the implementation, instead of an is-a relationship. The hasa relationship is maintained by composition. Abstraction has a reference of the
implementation. Refined abstraction provides more variations than abstraction.

Let's say IDrawShape is an interface with the drawShape method. DrawShape implements
the IDrawShape interface. We create an IContour bridge interface with
the drawContour method. The contour class implements the IContour interface. The
ellipse class will have a, b , r properties and drawShape (an instance of DrawShape). The
ellipse class implements the contour bridge interface to implement
the drawContour method. The drawContour method calls the drawShape method on
the drawShape instance.

```go
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
```

```
go run main.go

Drawing Contour
Drawing Shape
```