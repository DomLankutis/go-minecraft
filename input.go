package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func mouseCallback(w *glfw.Window, xpos, ypos float64) {
	if firstMouse {
		globalCamera.LastPos = [2]float64{xpos, ypos}
		firstMouse = false
	}

	xoffset, yoffset := xpos - globalCamera.LastPos[0], globalCamera.LastPos[1] - ypos
	globalCamera.LastPos[0], globalCamera.LastPos[1] = xpos, ypos

	sensitivity := 0.1
	xoffset *= sensitivity
	yoffset *= sensitivity

	globalCamera.Yaw += xoffset
	globalCamera.Yaw = math.Mod(globalCamera.Yaw, 360)
	if globalCamera.Pitch += yoffset; globalCamera.Pitch > 89.0 {
		globalCamera.Pitch = 89.0
	}else if globalCamera.Pitch < -89.0 {
		globalCamera.Pitch = -89.0
	}

	globalCamera.CamFront = mgl32.Vec3{
		float32(math.Cos(mgl64.DegToRad(globalCamera.Pitch)) * math.Cos(mgl64.DegToRad(globalCamera.Yaw))),
		float32(math.Sin(mgl64.DegToRad(globalCamera.Pitch))),
		float32(math.Cos(mgl64.DegToRad(globalCamera.Pitch)) * math.Sin(mgl64.DegToRad(globalCamera.Yaw))),
	}.Normalize()
}

func processInput(window *glfw.Window) {
	camSpeed := float32(20 * deltaTime)
	if window.GetKey(glfw.KeyW) == glfw.Press {
		globalCamera.CamPosition = globalCamera.CamPosition.Add(globalCamera.CamFront.Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		globalCamera.CamPosition = globalCamera.CamPosition.Sub(globalCamera.CamFront.Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		globalCamera.CamPosition = globalCamera.CamPosition.Add(mgl32.Vec3.Normalize(globalCamera.CamFront.Cross(mgl32.Vec3{0, 1, 0})).Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		globalCamera.CamPosition = globalCamera.CamPosition.Sub(mgl32.Vec3.Normalize(globalCamera.CamFront.Cross(mgl32.Vec3{0, 1, 0})).Mul(camSpeed))
	}

	chunkPos = mgl32.Vec2{float32(math.Round(float64(globalCamera.CamPosition.X()) / 16)), float32(math.Round(float64(globalCamera.CamPosition.Z() / 16)))}
}
