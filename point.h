#ifndef POINT_H
#define POINT_H

#include "point.hpp"

typedef void* CPoint;

#ifdef __cplusplus
extern "C" {
#endif

CPoint point_new(double, double);
double point_x(CPoint p);
double point_y(CPoint p);
void point_delete(CPoint p);

#ifdef __cplusplus
}
#endif

#endif
