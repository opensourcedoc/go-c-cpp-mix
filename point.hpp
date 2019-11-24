#pragma once

#ifdef __cplusplus
class Point {
public:
    Point(double, double);
    double x();
    double y();
private:
    double _x;
    double _y;
};
#endif
