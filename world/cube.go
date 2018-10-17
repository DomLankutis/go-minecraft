package world

import (
	"../glhf"
	"github.com/go-gl/mathgl/mgl32"
)

type Cube struct {
	texture    *glhf.Texture
	shader     *glhf.Shader
	vertexData *glhf.VertexSlice
}

func (c *Cube) Draw () {
	c.vertexData.Begin()
	defer c.vertexData.End()
	c.vertexData.Draw()
}

/*faces go in the order:
	0 - front
	1 - back
	2 - left
	3 - right
	4 - top
	5 - bottom
*/
func NewCube(shader *glhf.Shader, texture *glhf.Texture, texCoords mgl32.Vec2, faces [6]bool) *Cube{
	c := Cube{
		texture,
		shader,
		glhf.MakeVertexSlice(shader, 0, 36),
	}

	tS := 16 / float32(c.texture.Width())
	u0, v0 := texCoords.X(), texCoords.Y()
	u1, v1 := u0 + tS, v0 + tS

	vertexArr := []float32{}

	c.vertexData.Begin()
	defer c.vertexData.End()

	if faces[0] {
		vertexArr = append(vertexArr, []float32{
			// front //
			-0.5, -0.5, 0.5, 	u1 + tS, v1,
			0.5, -0.5, 0.5, 	u0 + tS, v1,
			0.5, 0.5, 0.5, 		u0 + tS, v0,
			0.5, 0.5, 0.5, 		u0 + tS, v0,
			-0.5, 0.5, 0.5, 	u1 + tS, v0,
			-0.5, -0.5, 0.5, 	u1 + tS, v1,
		}...)
	}
	if faces[1] {
		vertexArr = append(vertexArr, []float32{
			// back //
			-0.5, 0.5, -0.5,	u1 + tS, v0,
			0.5, 0.5, -0.5,		u0 + tS, v0,
			0.5, -0.5, -0.5,  	u0 + tS, v1,
			0.5, -0.5, -0.5,	u0 + tS, v1,
			-0.5, -0.5, -0.5,	u1 + tS, v1,
			-0.5, 0.5, -0.5,	u1 + tS, v0,
		}...)
	}
	if faces[2] {
		vertexArr = append(vertexArr, []float32{
			// left //
			-0.5, -0.5, -0.5,	u0 + tS, v1,
			-0.5, -0.5, 0.5,	u1 + tS, v1,
			-0.5, 0.5, 0.5,		u1 + tS, v0,
			-0.5, 0.5, 0.5, 	u1 + tS, v0,
			-0.5, 0.5, -0.5,	u0 + tS, v0,
			-0.5, -0.5, -0.5,	u0 + tS, v1,
		}...)
	}
	if faces[3] {
		vertexArr = append(vertexArr, []float32{
			// right //
			0.5, -0.5, 0.5,		u1 + tS, v1,
			0.5, -0.5, -0.5,	u0 + tS, v1,
			0.5, 0.5, -0.5, 	u0 + tS, v0,
			0.5, 0.5, -0.5,		u0 + tS, v0,
			0.5, 0.5, 0.5,		u1 + tS, v0,
			0.5, -0.5, 0.5,		u1 + tS, v1,
		}...)
	}
	if faces[4] {
		vertexArr = append(vertexArr, []float32{
			// top //
			-0.5, 0.5, 0.5, 	u0, v1,
			0.5, 0.5, 0.5,		u1, v1,
			0.5, 0.5, -0.5,		u1, v0,
			0.5, 0.5, -0.5,		u1, v0,
			-0.5, 0.5, -0.5,	u1, v1,
			-0.5, 0.5, 0.5,		u0, v1,
		}...)
	}
	if faces[5] {
		vertexArr = append(vertexArr, []float32{
			// bottom //
			-0.5, -0.5, -0.5, 	u0 + 2 * tS, v0,
			0.5, -0.5, -0.5,	u1 + 2 * tS, v0,
			0.5, -0.5, 0.5,		u1 + 2 * tS, v1,
			0.5, -0.5, 0.5,		u1 + 2 * tS, v1,
			-0.5, -0.5, 0.5,	u0 + 2 * tS, v1,
			-0.5, -0.5, -0.5,	u0 + 2 * tS, v0,
		}...)
	}
	c.vertexData.SetLen(len(vertexArr) / 5)
	c.vertexData.SetVertexData(vertexArr)

	return &c
}