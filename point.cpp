#include "point.hpp"

Point::Point(double x, double y)
{
  this->x = x;
  this->y = y;
}

double Point::getX()
{
  return this->x;
}

double Point::getY()
{
  return this->y;
}
