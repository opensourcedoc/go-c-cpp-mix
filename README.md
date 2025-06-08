# Using C++ Classes in Go

Go programs can interface with C libraries using **cgo**. However, cgo does not natively support C++ classes or functions. To work around this limitation, C++ code must be wrapped in a C-compatible interface.

This repository provides a minimal working example demonstrating how to:

* Wrap a C++ class with a plain C interface
* Compile and link the C++ code as a shared library
* Call C++ class methods from Go using cgo through the C wrapper

This technique is useful when integrating Go applications with existing C++ codebases or leveraging C++ libraries in a Go environment.

---

## License

Apache License 2.0
Copyright (c) 2019–2020, ByteBard
