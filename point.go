package point

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lstdc++

#include "point.h"
*/
import "C"

import (
	"fmt"
	_ "runtime"
)

type Point struct {
	pt C.CPoint
}

func NewPoint(x C.double, y C.double) *Point {
	p := new(Point)
	p.pt = C.point_new(x, y)
	//runtime.SetFinalizer(p, free)
	return p
}

func (p *Point) Delete() {
	C.point_delete(p.pt)
}

func (p *Point) X() C.double {
	return C.point_x(p.pt)
}

func (p *Point) Y() C.double {
	return C.point_y(p.pt)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.X(), p.Y())
}
