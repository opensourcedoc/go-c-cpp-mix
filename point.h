#ifndef POINT_H
#define POINT_H

#include "point.hpp"

typedef void* CPoint;

#ifdef __cplusplus
extern "C" {
#endif
  CPoint pointNew(double, double);
  double pointGetX(CPoint p);
  double pointGetY(CPoint p);
#ifdef __cplusplus
}
#endif

#endif
