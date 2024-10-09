#ifndef FLOAT_VECTOR_VTABLE_H
    #define FLOAT_VECTOR_VTABLE_H

    struct FloatVector_s;
    typedef struct FloatVector_s FloatVector;

    typedef struct FloatVectorVtable_s {
        int (*add)(FloatVector *self, FloatVector *other);
        FloatVector *(*add_ret)(FloatVector *self, FloatVector *other);
        int (*sub)(FloatVector *self, FloatVector *other);
        FloatVector *(*sub_ret)(FloatVector *self, FloatVector *other);
        int (*multiply_scalar)(FloatVector *self, float n);
        FloatVector *(*multiply_scalar_ret)(FloatVector *self, float n);
        float (*dot_product)(FloatVector *self, FloatVector *other);
        float (*norm)(FloatVector *self);
        void (*normalize)(FloatVector *self);
        FloatVector *(*make_copy)(FloatVector *self);
    } FloatVectorVTable;

#endif