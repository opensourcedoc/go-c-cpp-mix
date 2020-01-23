#ifndef POINT_H
#define POINT_H

typedef struct point_t point_t;

#ifdef __cplusplus
extern "C" {
#endif

point_t * point_new(double x, double y);
void point_delete(void *self);
double point_x(point_t *self);
double point_y(point_t *self);
double point_distance(point_t *p, point_t *q);

#ifdef __cplusplus
}
#endif

#endif  /* POINT_H */
