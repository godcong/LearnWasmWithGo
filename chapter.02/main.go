package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	cb := js.NewCallback(func(args []js.Value) {
		move := js.Global.Get("document").Call("getElementById", "myText").Get("value").Int()
		fmt.Println(move)
	})
	js.Global.Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
	// The goal of the channel is to wait indefinitly
	// Otherwise, the main function ends and the wasm modules stops
	<-c
}
