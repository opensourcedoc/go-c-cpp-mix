# 在 Go 語言中使用 C++ class

**警告** 本程式以 Valgrind 實測可能有 memory leak，僅供參考。

Go 語言可以藉由 cgo 使用 C 語言的函式；然而，cgo 無法直接使用 C++ 程式碼，所以，要將 C++ 函式和物件轉為 C 的型別。

本範例簡單示範如何在 Go 語言中使用 C++ class。
