#include "point.hpp"

Point::Point(double x, double y)
{
  this->_x = x;
  this->_y = y;
}

double Point::x()
{
  return this->_x;
}

double Point::y()
{
  return this->_y;
}
