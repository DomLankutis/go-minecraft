package main

import (
	"./glhf"
	"./utils"
	"./world"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ojrac/opensimplex-go"
	"log"
	"math"
)

type ChunkRender struct {
	chunks  map[mgl32.Vec2]*world.Chunk
	visible []*world.Chunk
	cube    *world.Cube
	seed    opensimplex.Noise32
	camera  *utils.Camera
	shader  *glhf.Shader
	texture *glhf.Texture
}

func NewChunkRender(cube *world.Cube, seed opensimplex.Noise32, camera *utils.Camera, shader *glhf.Shader, texture *glhf.Texture) ChunkRender{
	return ChunkRender{
		chunks: map[mgl32.Vec2]*world.Chunk{},
		visible: []*world.Chunk{},
		cube: cube,
		seed: seed,
		camera: camera,
		shader: shader,
		texture: texture,
	}
}

func (c *ChunkRender) GetChunkPos(pos mgl32.Vec2) mgl32.Vec2 {
	return mgl32.Vec2{float32(math.Round(float64(pos.X() / 16))), float32(math.Round(float64(pos.Y() / 16)))}
}

func (c *ChunkRender) GenerateChunkAt(pos mgl32.Vec2) {
	if c.chunks[pos] == nil {
		log.Println("made at ", pos)
		c.chunks[pos] = world.GenerateChunkAt(c.seed, pos, c.cube, c.shader, c.texture)
	}
}

func (c *ChunkRender) GetVisibleChunks(distance float32) {
	c.visible = []*world.Chunk{}

	view := c.GetChunkPos(mgl32.Vec2{c.camera.CamPosition.X() ,c.camera.CamPosition.Z()})
	c.GenerateChunkAt(view)
	c.visible = append(c.visible, c.chunks[view])
	for x := view.X() - 2; x <= view.X() + 2; x++ {
	for z := view.Y() - 2; z <= view.Y() + 2; z++ {
		pos := mgl32.Vec3{x * 16, 0, z * 16}

		if c.camera.Frustum.ContainsChunk(pos, 16) {
			chunkPos := mgl32.Vec2{x, z}
			c.GenerateChunkAt(chunkPos)
			c.visible = append(c.visible, c.chunks[chunkPos])
		}
	}
	}
}

func (c *ChunkRender) Draw() {
	c.shader.Begin()
	defer c.shader.End()
	c.texture.Begin()
	defer c.texture.End()

	for _, chunk := range c.visible {
		chunk.Draw(*c.camera, c.shader)
	}
}
