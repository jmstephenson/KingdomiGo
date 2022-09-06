package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	width := 1280
	height := 720

	window := js.Global()
	ctx, canvasEl := initScreen(window, width, height)
	ctx.SetGlobalCompositeOperation("destination-over")

	mouseDown := false

	mouseDownEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		evt := args[0]
		if evt.Get("target") != canvasEl {
			return nil
		}
		mx := evt.Get("clientX").Float()
		my := evt.Get("clientY").Float()

		return nil
	})
	defer mouseDownEvt.Release()

	mouseUpEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouseDown = false
		return nil
	})
	defer mouseUpEvt.Release()

	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !mouseDown {
			return nil
		}
		evt := args[0]
		if evt.Get("target") != canvasEl {
			return nil
		}
		mx := evt.Get("clientX").Float() * worldScale
		my := evt.Get("clientY").Float() * worldScale
		thing.AddCircle(mx, my)
		return nil
	})
	defer mouseMoveEvt.Release()

	registerEventListeners()

	var gameLoop js.Func
	gameLoop = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			now := args[0].Float()

			game.Ship.Human = human

			if game.State == "menu" {
				game.Update(1.0/60.0, humanButtons)
				timeLeft = TIME_TO_PLAY
			} else if human {
				game.Update(1.0/60.0, humanButtons)
				timeLeft -= dt

			} else {
				botButtons, err := bot.GetOutputs(bots[botN], game)
				if err != nil {
					fmt.Println("bot outputs error", err)
					return
				}
				game.Update(1.0/60.0, botButtons)
				timeLeft -= dt

			}

			if game.Dead || game.Win {
				for _, b := range bots {
					b.Flush()
				}

				if human {
					humanscore = append(humanscore, game.Score)
				} else {
					botscore = append(botscore, game.Score)
				}
				human = !human
				if human {
					game.Seed = int64(len(humanscore)) + 21
					game.MenuText = fmt.Sprintf("PRESS ENTER: YOUR TURN")
				} else {
					botN = (botN + 1) % 5
					game.Seed = int64(len(botscore)) + 21
					game.MenuText = fmt.Sprintf("PRESS ENTER: BOTS TURN")
				}

				game.Dead = false
				game.Win = false
			}

			// Clear Screen
			ctx.ClearRect(0, 0, width, height) // clear canvas

			// Render
			render.RenderGame(ctx, game, human)
			render.RenderScorePlayer(ctx, game, humanscore, botscore, human, botN, timeLeft)
			diagnostics.Render(ctx)

			// Interactions phase
			window.Call("requestAnimationFrame", gameLoop)
		}()
		return nil
	})

	// start the game render
	window.Call("requestAnimationFrame", gameLoop)

	// Never return
	done := make(chan struct{}, 0)
	<-done
}

func getGenome(genomeStr string) *network.Network {
	genome, _ := genetics.ReadGenome(strings.NewReader(genomeStr), 1)

	net, _ := genome.Genesis(1)

	return net
}
