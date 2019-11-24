# Using C++ class in Golang

**Warning** The sample program may have memory leak. Use it with caution.

Golang programs can utilize C functions through cgo. Nevertheless, cgo cannot access C++ code base. To utilize C++ code base in Golang, we have to wrap C++ code with C.

Here is a small sample to show such usage.

Copyright 2019, Michael Chen; licensed under Apache 2.0.
