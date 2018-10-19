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

func NewChunkRender(cube *world.Cube, seed opensimplex.Noise32, camera *utils.Camera, shader *glhf.Shader, texture *glhf.Texture) ChunkRender {
	return ChunkRender{
		chunks:  map[mgl32.Vec2]*world.Chunk{},
		visible: []*world.Chunk{},
		cube:    cube,
		seed:    seed,
		camera:  camera,
		shader:  shader,
		texture: texture,
	}
}

func (c *ChunkRender) castRay() (bool, mgl32.Vec3, mgl32.Vec2) {
	var ray mgl32.Vec3
	ray = c.camera.CamPosition
	for x := 0; x <= 3; x++ {
		chunkPos := c.GetChunkPos(mgl32.Vec2{ray.X(), ray.Z()})
		if c.chunks[chunkPos] != nil {
			blockPos := mgl32.Vec3{float32(math.Round(float64(ray.X()))), float32(math.Round(float64(ray.Y()))), float32(math.Round(float64(ray.Z())))}
			if c.chunks[chunkPos].BlockExistsAt(blockPos) {
				return true, blockPos, chunkPos
			}
		}
		ray = ray.Sub(c.camera.GetView().Row(2).Vec3())
	}
	return false, mgl32.Vec3{}, mgl32.Vec2{}
}

func (c *ChunkRender) DestroyBlock(state *int) {
	if exist, pos, cPos := c.castRay(); exist {
		c.chunks[cPos].DestroyBlockAt(pos)
	}
	*state = 0
}

func (c *ChunkRender) CreateBlock(state *int) {
	if exists, pos, cPos := c.castRay(); exists {
		pos = pos.Add(c.camera.GetView().Row(2).Vec3())
		pos = mgl32.Vec3{float32(math.Round(float64(pos.X()))), float32(math.Round(float64(pos.Y()))), float32(math.Round(float64(pos.Z())))}
		c.chunks[cPos].CreateBlockAt(pos)
	}
	*state = 0
}

func (c *ChunkRender) GetChunkPos(pos mgl32.Vec2) mgl32.Vec2 {
	return mgl32.Vec2{float32(math.Round(float64(pos.X() / 16))), float32(math.Round(float64(pos.Y() / 16)))}
}

func (c *ChunkRender) GenerateChunkAt(pos mgl32.Vec2) {
	if c.chunks[pos] == nil {
		log.Println("made at ", pos)
		c.chunks[pos] = world.GenerateChunkAt(c.seed, pos, c.cube)
	}
}

func (c *ChunkRender) GetVisibleChunks(distance float32) {
	c.visible = []*world.Chunk{}

	view := c.GetChunkPos(mgl32.Vec2{c.camera.CamPosition.X(), c.camera.CamPosition.Z()})
	c.GenerateChunkAt(view)
	c.visible = append(c.visible, c.chunks[view])
	for x := view.X() - 2; x <= view.X()+2; x++ {
		for z := view.Y() - 2; z <= view.Y()+2; z++ {
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
