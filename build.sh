#!/bin/sh

GOOS=js GOARCH=wasm go build -o web/build/lib.wasm ./src/game/kingdomino.go