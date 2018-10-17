package main

import (
	"./glhf"
	"./utils"
	"./world"

	"github.com/ojrac/opensimplex-go"
	"time"

	"image"
)

var (
	WIDTH = 1920
	HEIGHT = 1080

	frameRate = time.Tick(time.Second/144)
	frames = 0
	deltaTime = 0.0
	lastFrame = 0.0
	seed opensimplex.Noise32

	vertexFormat = glhf.AttrFormat{
		{"position", glhf.Vec3},
		{"texture", glhf.Vec2},
	}
	shader *glhf.Shader
	texture *glhf.Texture

	testImg *image.NRGBA

	firstMouse = true

	globalCamera utils.Camera

	chunkRender ChunkRender
	chunk world.Chunk

	vertexShader string
	fragmentShader string
)

