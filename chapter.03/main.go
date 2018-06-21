package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	add := func(i []js.Value) {
		js.Global.Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	}
	js.Global.Set("add", js.NewCallback(add))
	<-c
}
