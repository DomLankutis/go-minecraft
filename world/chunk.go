package world

import (
	"../glhf"
	"../utils"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ojrac/opensimplex-go"
)

func left(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X() - 1, vec3.Y(), vec3.Z()}
}
func right(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X() + 1, vec3.Y(), vec3.Z()}
}
func top(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X(), vec3.Y() + 1, vec3.Z()}
}
func bottom(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X(), vec3.Y() - 1, vec3.Z()}
}
func back(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X(), vec3.Y(), vec3.Z() - 1}
}
func front(vec3 mgl32.Vec3) mgl32.Vec3{
	return mgl32.Vec3{vec3.X(), vec3.Y(), vec3.Z() + 1}
}

type Chunk struct {
	cube      *Cube
	blocks    map[mgl32.Vec3]*Cube
	drawable  map[mgl32.Vec3]*Cube
	Size      float32
	IsChanged bool
}

func NewChunk(cube *Cube) Chunk{
	return Chunk{
		cube,
		map[mgl32.Vec3]*Cube{},
		map[mgl32.Vec3]*Cube{},
		16,
		true,
	}
}

func (c *Chunk) UpdateMesh() {
	if c.IsChanged {
		for pos := range c.blocks {
			faces := [6]bool{false, false, false, false, false, false}
			add := false
			if c.blocks[front(pos)] == nil {
				faces[0] = true
				add = true
			}
			if c.blocks[back(pos)] == nil {
				faces[1] = true
				add = true
			}
			if c.blocks[left(pos)] == nil {
				faces[2] = true
				add = true
			}
			if c.blocks[right(pos)] == nil {
				faces[3] = true
				add = true
			}
			if c.blocks[top(pos)] == nil {
				faces[4] = true
				add = true
			}
			if c.blocks[bottom(pos)] == nil {
				faces[5] = true
				add = true
			}
			if add {
				c.drawable[pos] = NewCube(c.cube.shader, c.cube.texture, mgl32.Vec2{0, 0}, faces)
			}
		}
	}
	c.IsChanged = false
}

func GenerateChunkAt(seed opensimplex.Noise32, vec2 mgl32.Vec2, cube *Cube, shader *glhf.Shader, texture *glhf.Texture) *Chunk {
	c := NewChunk(cube)

	div := float32(16 * 4)
	startX := vec2.X() * c.Size
	startZ := vec2.Y() * c.Size
	endX := startX + c.Size
	endZ := startZ + c.Size
	var pos mgl32.Vec3
	for x := startX; x <= endX; x++ {
	for z := startZ; z <= endZ; z++ {
		y := float32(int(seed.Eval2(x/div, z/div) * c.Size))
		for yy := float32(-16); yy <= y; yy++ {
			pos = mgl32.Vec3{x, yy, z}
			c.blocks[pos] = c.cube
		}
	}
	}

	return &c
}

func (c *Chunk) Draw(camera utils.Camera, shader *glhf.Shader) {
	c.UpdateMesh()
	MV := camera.GetProjection().Mul4(camera.GetView())

	for pos := range c.drawable {
		MVP := MV.Mul4(mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()))
		shader.SetUniformAttr(0, MVP)
		c.drawable[pos].Draw()
	}
}