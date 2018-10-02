package main

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
	"time"

	"./shape"
	"./Utils"
	"image"
)

var (
	WIDTH = 1920
	HEIGHT = 1080

	frameRate = time.Tick(time.Second/144)
	frames = 0
	deltaTime = 0.0
	lastFrame = 0.0

	vertexFormat = glhf.AttrFormat{
		{"position", glhf.Vec3},
		{"texture", glhf.Vec2},
	}

	testImg *image.NRGBA

	firstMouse = true

	globalCamera Utils.Camera
	oldChunkPos = mgl32.Vec2{0, 0}
	chunkPos = mgl32.Vec2{}
	cube shape.Cube

	vertexShader string
	fragmentShader string
)

