#!/bin/sh

GOOS=js GOARCH=wasm go build -o web/build/main.wasm ./src/web_integration/