#include <cmath>
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

double Point::distance(Point *p, Point *q)
{
    double dx = p->x() - q->x();
    double dy = p->y() - q->y();

    return sqrt(dx * dx + dy * dy);
}
