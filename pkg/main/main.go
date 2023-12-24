package main

import (
	"learn/pkg/display"
	"learn/pkg/msg"
)


func main() {
	msg.Hi()
	display.Display("Hello from Display")
	msg.Exciting("an exciting message")
}