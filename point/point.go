package point

// #include "point.h"
import "C"
import "unsafe"

type Point struct {
	point *C.point_t
}

func NewPoint(x float64, y float64) *Point {
	pt := new(Point)
	pt.point = C.point_new(C.double(x), C.double(y))
	return pt
}

func (pt *Point) Delete() {
	C.point_delete(unsafe.Pointer(pt.point))
}

func (pt *Point) X() float64 {
	return float64(C.point_x(pt.point))
}

func (pt *Point) Y() float64 {
	return float64(C.point_y(pt.point))
}

func Distance(p *Point, q *Point) float64 {
	return float64(C.point_distance(p.point, q.point))
}
