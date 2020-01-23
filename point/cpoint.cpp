#include <cassert>
#include <cstdlib>
#include "point.h"
#include "point.hpp"

struct point_t {
    Point *obj;
};

point_t * point_new(double x, double y)
{
    point_t *pt = (point_t *) malloc(sizeof(point_t));
    if (!pt)
        return NULL;

    pt->obj = new Point(x, y);

    return pt;
}

void point_delete(void *self)
{
    assert(self);

    Point *point = ((point_t *) self)->obj;
    delete point;

    free(self);
}

double point_x(point_t *self)
{
    assert(self);

    return self->obj->x();
}

double point_y(point_t *self)
{
    assert(self);

    return self->obj->y();
}

double point_distance(point_t *p, point_t *q)
{
    assert(p);
    assert(q);

    return Point::distance(p->obj, q->obj);
}
