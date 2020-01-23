#ifndef POINT_HPP
#define POINT_HPP

class Point {
public:
    Point(double x, double y);
    double x();
    double y();
    static double distance(Point *p, Point *q);
private:
    double _x;
    double _y;
};

#endif  /* POINT_HPP */