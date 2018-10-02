package shape

import (
	"../Utils"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ojrac/opensimplex-go"
	"math"
)

type Chunk struct {
	Position       mgl32.Vec2
	Cube           *Cube
	blockPositions []mgl32.Vec3
	size           float32
}

func trunc(val float32) float32 {
	return float32(math.Round(float64(val)))
}

func GenerateChunkAroundPlayer(camera Utils.Camera, cube *Cube, pos mgl32.Vec2, seed opensimplex.Noise32) []Chunk {
	chunks := []Chunk{}

	a := float32(camera.Yaw - 90)
	for di := float32(0); di <= 4; di++ {
		d := mgl32.Rotate2D(mgl32.DegToRad(a)).Mul2x1(mgl32.Vec2{0, di})

		rX := mgl32.Rotate2D(mgl32.DegToRad(55)).Mul2x1(d).Add(pos)
		lX := mgl32.Rotate2D(mgl32.DegToRad(-55)).Mul2x1(d).Add(pos)

		chunks = append(chunks, GenerateChunk(pos, cube, seed))
		var mm mgl32.Vec2
		if rX.X() > lX.X() {
			mm = mgl32.Vec2{lX.X(), rX.X()}
		} else {
			mm = mgl32.Vec2{rX.X(), lX.X()}
		}
		var my mgl32.Vec2
		if rX.Y() > lX.Y() {
			my = mgl32.Vec2{lX.Y(), rX.Y()}
		} else {
			my = mgl32.Vec2{rX.Y(), lX.Y()}
		}
		for x := trunc(mm.X()); x <= trunc(mm.Y()); x++ {
			for y := trunc(my.X()); y <= trunc(my.Y()); y++ {
				chunks = append(chunks, GenerateChunk(mgl32.Vec2{x, y}, cube, seed))
			}
		}
	}

	return chunks
}

func GenerateChunk(pos mgl32.Vec2, cube *Cube, seed opensimplex.Noise32) Chunk {
	c := Chunk{
		pos,
		cube,
		nil,
		16,
	}

	div := float32(32)

	c.blockPositions = append(c.blockPositions, mgl32.Vec3{pos.X() * c.size, -1.0, pos.Y() * c.size})

	for x := pos.X() * c.size - c.size / 2; x < pos.X() * c.size + c.size / 2; x++ {
		for z := pos.Y() * c.size - c.size / 2; z < pos.Y() * c.size + c.size / 2; z++ {
			y := float32(math.Floor(float64(seed.Eval2(x / div, z / div) * 16)))
			c.blockPositions = append(c.blockPositions, mgl32.Vec3{x, y, z})
		}
	}

	return c
}

func (c *Chunk) Draw(camera Utils.Camera) {
	MV := camera.GetProjection().Mul4(camera.GetView())


	c.Cube.shader.Begin()
	defer c.Cube.shader.End()

	c.Cube.texture.Begin()
	defer c.Cube.texture.End()

	c.Cube.slice.Begin()
	defer c.Cube.slice.End()

	for _, tr := range c.blockPositions {
		MVP := MV.Mul4(mgl32.Translate3D(tr.X(), tr.Y(), tr.Z()).Mul4(c.Cube.model))
		c.Cube.shader.SetUniformAttr(1, MVP)
		c.Cube.slice.Draw()
	}

}

