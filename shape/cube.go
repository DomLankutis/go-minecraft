package shape

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
)

type Cube struct {
	shader       *glhf.Shader
	texture      *glhf.Texture
	slice        *glhf.VertexSlice
	model        mgl32.Mat4
}

func NewCube(shader *glhf.Shader, textures *glhf.Texture, tileSize float32, textureOffset mgl32.Vec2) Cube {
	tS := tileSize / float32(textures.Width())
	u0, v0 := textureOffset.X(), textureOffset.Y()
	u1, v1 := u0 + tS, v0 + tS


	slice := glhf.MakeVertexSlice(shader, 36, 36)
	slice.Begin()
	slice.SetVertexData([]float32{
		-1, -1, -1, u1 + tS, v1, // Side
		1, -1, -1, 	u0 + tS, v1,
		1, 1, -1, 	u0 + tS, v0,
		1, 1, -1, 	u0 + tS, v0,
		-1, 1, -1, 	u1 + tS, v0,
		-1, -1, -1, u1 + tS, v1,

		-1, -1, 1, 	u1 + tS, v1, // Side
		1, -1, 1, 	u0 + tS, v1,
		1, 1, 1, 	u0 + tS, v0,
		1, 1, 1, 	u0 + tS, v0,
		-1, 1, 1, 	u1 + tS, v0,
		-1, -1, 1, 	u1 + tS, v1,

		-1, 1, 1, 	v1 + tS, u0, // Side
		-1, 1, -1, 	v0 + tS, u0,
		-1, -1, -1, v0 + tS, u1,
		-1, -1, -1, v0 + tS, u1,
		-1, -1, 1, 	v1 + tS, u1,
		-1, 1, 1, 	v1 + tS, u0,

		1, 1, 1, 	v1 + tS, u0,  // Side
		1, 1, -1, 	v0 + tS, u0,
		1, -1, -1, 	v0 + tS, u1,
		1, -1, -1, 	v0 + tS, u1,
		1, -1, 1, 	v1 + tS, u1,
		1, 1, 1, 	v1 + tS, u0,

		-1, -1, -1, u0 + tS * 2, v1, // Bottom
		1, -1, -1, 	u1 + tS * 2, v1,
		1, -1, 1, 	u1 + tS * 2, v0,
		1, -1, 1, 	u1 + tS * 2, v0,
		-1, -1, 1, 	u0 + tS * 2, v0,
		-1, -1, -1, u0 + tS * 2, v1,

		-1, 1, -1, 	u0, v1, // Top
		1, 1, -1, 	u1, v1,
		1, 1, 1, 	u1, v0,
		1, 1, 1, 	u1, v0,
		-1, 1, 1, 	u0, v0,
		-1, 1, -1, 	u0, v1,
	})
	slice.End()

	c := Cube{
		shader,
		textures,
		slice,
		mgl32.Scale3D(0.5, 0.5, 0.5),
	}
	c.shader.Begin()
	c.shader.SetUniformAttr(0, mgl32.Vec4{1, 1, 1, 1})
	c.shader.End()
	return c
}

func (c *Cube) TransformModel(transformation mgl32.Mat4) {
	c.model = transformation
}