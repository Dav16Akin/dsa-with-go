package designpatterns

import "fmt"

type IDrawShape interface {
	drawShape(x[5] float32, y[5] float32)
}

type DrawShape struct {}

func (drawshape DrawShape) drawShape(x[5] float32, y[5] float32) {
	fmt.Println("Drawing shape")
}

type IContour interface {
	drawContour(x[5] float32 ,y[5] float32)
	resizeFactor(factor int)
}

type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

func (contour DrawContour) drawContour(x[5] float32 ,y[5] float32) {
	fmt.Println("Drawing contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour DrawContour) resizeFactor(factor int) {
	contour.factor = factor
}


// MAIN FUNCTION

func RunBridgePattren () {
	fmt.Println("--> Bridge Pattern")
	var x = [5]float32{1,2,3,4,5}
	var y = [5]float32{1,2,3,4,5}

	var contour IContour = DrawContour{x,y,DrawShape{},2}
	contour.drawContour(x,y)
	contour.resizeFactor(2)
}