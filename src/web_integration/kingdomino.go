// Wasming
// compile: GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
package main

import (
	"syscall/js"
)

var (
	width      float64
	height     float64
	mousePos   [2]float64
	ctx        Context2D
	lineDistSq float64 = 100 * 100
)

type Position struct {
	width  int
	height int
	posX   float64
	posY   float64
}

func (p *Position) Collision(posX, posY float64) bool {
	if posX < p.posX || posX > p.posX+width || posY < p.posY || posY > p.posX+height {
		return false
	}
}

func EventLoop() {

	// Init Canvas stuff
	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		mousePos[0] = e.Get("clientX").Float()
		mousePos[1] = e.Get("clientY").Float()
		return nil
	})
	defer mouseMoveEvt.Release()

	var loop js.Func
	loop = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx.ClearRect()
	})

	var renderFrame js.Func
	var tmark float64
	var markCount = 0
	var tdiffSum float64

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		now := args[0].Float()
		tdiff := now - tmark
		tdiffSum += now - tmark
		markCount++
		if markCount > 10 {
			tdiffSum, markCount = 0, 0
		}
		tmark = now

		// Pull window size to handle resize
		curBodyW := doc.Get("body").Get("clientWidth").Float()
		curBodyH := doc.Get("body").Get("clientHeight").Float()
		if curBodyW != width || curBodyH != height {
			width, height = curBodyW, curBodyH
			canvasEl.Set("width", width)
			canvasEl.Set("height", height)
		}
		Update(tdiff / 1000)

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done

}
