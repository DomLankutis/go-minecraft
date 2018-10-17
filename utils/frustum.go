package utils

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Frustum struct {
	planes []mgl32.Vec4
}

func (f *Frustum) ContainsPoint(p mgl32.Vec3) bool {
	for _, plane := range f.planes {
		if plane.Dot(p.Vec4(1)) < 0 {
			return false
		}
	}
	return true
}

func (f *Frustum) ContainsChunk(pos mgl32.Vec3, size float32) bool {
	points := []mgl32.Vec3{
		mgl32.Vec3{pos.X(), 0, pos.Z()},
		mgl32.Vec3{pos.X() + size, 0, pos.Z() + size},
		mgl32.Vec3{pos.X() + size, 0, pos.Z()},
		mgl32.Vec3{pos.X(), 0, pos.Z() + size},
	}

	in := 0
	for _, point := range points {
		if f.ContainsPoint(point) {
			in++
		}
	}
	return in >= 1
}

func (f *Frustum) Update(mat4 mgl32.Mat4) {
	c1, c2, c3, c4 := mat4.Rows()
	f.planes = []mgl32.Vec4{
		c4.Add(c1),          // left
		c4.Sub(c1),          // right
		c4.Sub(c2),          // top
		c4.Add(c2),          // bottom
		c4.Mul(0.1).Add(c3), // front
		c4.Mul(100).Sub(c3), // back
	}
}

func GenFrustum(c *Camera) Frustum{
	c1, c2, c3, c4 := c.GetProjection().Mul4(c.GetView()).Rows()
	return Frustum{
		planes: []mgl32.Vec4{
			c4.Add(c1),          	// left
			c4.Sub(c1),          	// right
			c4.Sub(c2),          	// top
			c4.Add(c2),          	// bottom
			c4.Mul(0.1).Add(c3), 	// front
			c4.Mul(100).Sub(c3), 	// back
		},
	}
}
