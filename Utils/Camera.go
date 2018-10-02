package Utils

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	shader      *glhf.Shader
	projection  mgl32.Mat4
	view        mgl32.Mat4
	CamFront    mgl32.Vec3
	CamPosition mgl32.Vec3
	LastPos		[2]float64
	Yaw			float64
	Pitch		float64
}

func InitCamera(width, height int, shader *glhf.Shader) Camera {
	return Camera{
		shader:     shader,
		projection: mgl32.Perspective(45, float32(width / height), 0.1, 1000),
		CamFront:   mgl32.Vec3{0.8, -0.1, .4},
		LastPos: 	[2]float64{float64(width / 2), float64(height / 2)},
		CamPosition: mgl32.Vec3{0, 0 , 0},
	}
}

func (c *Camera) Update() {
	c.view = mgl32.LookAtV(c.CamPosition, c.CamPosition.Add(c.CamFront), mgl32.Vec3{0, 1, 0})
}

func (c *Camera) GetProjection() mgl32.Mat4 {
	return c.projection
}

func (c *Camera) GetView() mgl32.Mat4 {
	return c.view
}