#include "point.hpp"
#include "point.h"

CPoint point_new(double x, double y)
{
  return (CPoint) new Point(x, y);
}

double point_x(CPoint p)
{
  return ((Point*) p)->x();
}

double point_y(CPoint p)
{
  return ((Point*) p)->y();
}

void point_delete(void *p)
{
  delete ((Point*) p);
}
