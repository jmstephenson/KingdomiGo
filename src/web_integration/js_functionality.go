package main

import (
	gs "kingdomino/src/gamestate"
	"syscall/js"
)

func mouseMoveEvt(this js.Value, args []js.Value) interface{} {
	e := args[0]
	mousePos[0] = e.Get("clientX").Float()
	mousePos[1] = e.Get("clientY").Float()

	return nil
}

func initScreen(window js.Value, w, h int) (*Context2D, js.Value) {
	document := window.Get("document")
	canvasEle := document.Call("createElement", "canvas")

	canvasEle.Set("id", "canvas")
	canvasEle.Set("width", w)
	canvasEle.Set("height", h)

	document.Get("body").Call("appendChild", canvasEle)
	document.Get("body").Set("style", "margin: 0px; overflow: hidden")

	return NewContext2D(canvasEle), canvasEle
}

func registerEventListener(ev string, function func(js.Value, []js.Value) interface{}) {
	js.Global().Call(
		"addEventListener", ev, js.FuncOf(function),
	)
}

func registerEventListeners() {
	registerEventListener("keyup", keyUp)
	registerEventListener("keydown", keyDown)

}

// key presses
func keyUp(this js.Value, args []js.Value) interface{} {
	e := args[0]
	e.Call("preventDefault")
	gs.HumanButtons[e.Get("keyCode").Int()] = false
	return nil
}

func keyDown(this js.Value, args []js.Value) interface{} {
	e := args[0]
	e.Call("preventDefault")
	gs.HumanButtons[e.Get("keyCode").Int()] = true
	return nil
}

func gameLoop(this js.Value, args []js.Value) interface{} {

}
