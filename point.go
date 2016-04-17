package point

// #include "point.h"
import "C"

import "fmt"

type Point struct {
  pt C.CPoint
}

func Init(x C.double, y C.double) *Point {
  p := new(Point)
  p.pt = C.pointNew(x, y)
  return p
}

func (p *Point) GetX() C.double {
  return C.pointGetX(p.pt)
}

func (p *Point) GetY() C.double {
  return C.pointGetY(p.pt)
}

func (p *Point) ToString() string {
  return fmt.Sprintf("(%f, %f)", p.GetX(), p.GetY())
}
