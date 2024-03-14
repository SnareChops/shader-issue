package main

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed shader.kage
var shader []byte

var Shader *ebiten.Shader

func init() {
	var err error
	Shader, err = ebiten.NewShader(shader)
	if err != nil {
		panic(err)
	}
}
