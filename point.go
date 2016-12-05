package point

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lstdc++

#include "point.h"
*/
import "C"

import "fmt"

type Point struct {
  pt C.CPoint
}

func Init(x C.double, y C.double) *Point {
  p := new(Point)
  p.pt = C.point_new(x, y)
  return p
}

func (p *Point) GetX() C.double {
  return C.point_get_x(p.pt)
}

func (p *Point) GetY() C.double {
  return C.point_get_y(p.pt)
}

func (p *Point) Free() {
	C.point_free(p.pt)
}

func (p *Point) ToString() string {
  return fmt.Sprintf("(%f, %f)", p.GetX(), p.GetY())
}
