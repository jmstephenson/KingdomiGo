package main

import (
	gs "kingdomino/src/gamestate"
	"syscall/js"
)

// func mouseMoveEvt(this js.Value, args []js.Value) interface{} {
// 	e := args[0]
// 	mousePos[0] = e.Get("clientX").Float()
// 	mousePos[1] = e.Get("clientY").Float()

// 	return nil
// }

func registerEventListener(ev string, function func(js.Value, []js.Value) interface{}) {
	js.Global().Call(
		"addEventListener", ev, js.FuncOf(function),
	)
}

func RegisterEventListeners() {
	registerEventListener("keyup", keyUp)
	registerEventListener("keydown", keyDown)

	AddEventListner("draw_tile_1", clickDrawTile)
	AddEventListner("draw_tile_2", clickDrawTile)
	AddEventListner("draw_tile_3", clickDrawTile)
	AddEventListner("draw_tile_4", clickDrawTile)

	AddEventListner("revealed_tile_1", clickRevealedTile)
	AddEventListner("revealed_tile_2", clickRevealedTile)
	AddEventListner("revealed_tile_3", clickRevealedTile)
	AddEventListner("revealed_tile_4", clickRevealedTile)

	AddEventListner("active_player_space_1", clickPlayerTile)
	AddEventListner("active_player_space_2", clickPlayerTile)
	AddEventListner("active_player_space_3", clickPlayerTile)
	AddEventListner("active_player_space_4", clickPlayerTile)
	AddEventListner("active_player_space_5", clickPlayerTile)
	AddEventListner("active_player_space_6", clickPlayerTile)
	AddEventListner("active_player_space_7", clickPlayerTile)
	AddEventListner("active_player_space_8", clickPlayerTile)
	AddEventListner("active_player_space_9", clickPlayerTile)
	AddEventListner("active_player_space_10", clickPlayerTile)
	AddEventListner("active_player_space_11", clickPlayerTile)
	AddEventListner("active_player_space_12", clickPlayerTile)
	AddEventListner("active_player_space_13", clickPlayerTile)
	AddEventListner("active_player_space_14", clickPlayerTile)
	AddEventListner("active_player_space_15", clickPlayerTile)
	AddEventListner("active_player_space_16", clickPlayerTile)
	AddEventListner("active_player_space_17", clickPlayerTile)
	AddEventListner("active_player_space_18", clickPlayerTile)
	AddEventListner("active_player_space_19", clickPlayerTile)
	AddEventListner("active_player_space_20", clickPlayerTile)
	AddEventListner("active_player_space_21", clickPlayerTile)
	AddEventListner("active_player_space_22", clickPlayerTile)
	AddEventListner("active_player_space_23", clickPlayerTile)
	AddEventListner("active_player_space_24", clickPlayerTile)
	AddEventListner("active_player_space_25", clickPlayerTile)

}

func AddEventListner(eleId string, f func(js.Value, []js.Value) interface{}) {
	ele := js.Global().Get("document").Call("getElementById", eleId)

	ele.Set("onclick", js.FuncOf(f))
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

// dom element clicks
func clickDrawTile(this js.Value, args []js.Value) interface{} {
	println("clicked draw tile: ", this.Get("id").String(), args)

	return nil
}

func clickRevealedTile(this js.Value, args []js.Value) interface{} {
	println("clicked revealed tile: ", this.Get("id").String(), args)

	return nil
}

func clickPlayerTile(this js.Value, args []js.Value) interface{} {
	println("clicked player tile: ", this.Get("id").String(), args)

	return nil
}
