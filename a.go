package main

import "syscall/js"

func main() {
	a := []any{nil}
	b := []any{a, a}
	js.Global().Call("foo", js.ValueOf(b))
}
