#include "point.hpp"
#include "point.h"

CPoint point_new(double x, double y)
{
  return (CPoint) new Point(x, y);
}

double point_get_x(CPoint p)
{
  return ((Point*) p)->getX();
}

double point_get_y(CPoint p)
{
  return ((Point*) p)->getY();
}

void point_free(CPoint p)
{
  delete (Point*) p;
}
