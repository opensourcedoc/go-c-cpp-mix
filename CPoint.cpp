#include "point.hpp"
#include "point.h"

CPoint pointNew(double x, double y)
{
  return (CPoint) new Point(x, y);
}

double pointGetX(CPoint p)
{
  return ((Point*) p)->getX();
}

double pointGetY(CPoint p)
{
  return ((Point*) p)->getY();
}
