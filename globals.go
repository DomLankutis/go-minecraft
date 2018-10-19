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
	WIDTH  = 1920
	HEIGHT = 1080

	FRAMERATE = time.Tick(time.Second / 144)
	FRAMES    = 0
	DELTATIME = 0.0
	LASTFRAME = 0.0

	MOUSESTATE = 0

	SEED opensimplex.Noise32

	VERTEXFORMAT = glhf.AttrFormat{
		{"position", glhf.Vec3},
		{"texture", glhf.Vec2},
	}
	SHADER  *glhf.Shader
	texture *glhf.Texture

	TEXTUREATLAS *image.NRGBA

	FIRSTMOUSE = true

	GLOBALCAMERA utils.Camera

	CHUNKRENDERER ChunkRender
	CHUNK         world.Chunk

	VERTEXSHADER   string
	FRAGMENTSHADER string
)
