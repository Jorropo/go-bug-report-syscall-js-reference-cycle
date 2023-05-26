package main

import "syscall/js"

func main() {
	a := []any{nil}
	a[0] = a
	js.Global().Get("console").Call("log", js.ValueOf(a))
}
