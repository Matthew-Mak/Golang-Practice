package main

import (
	"lets_go_again/lectures/demo/pkg/display"
	"lets_go_again/lectures/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("An exciting message")
}
