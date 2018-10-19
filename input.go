package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func mouseCallback(w *glfw.Window, xpos, ypos float64) {
	if FIRSTMOUSE {
		GLOBALCAMERA.LastPos = [2]float64{xpos, ypos}
		FIRSTMOUSE = false
	}

	xoffset, yoffset := xpos-GLOBALCAMERA.LastPos[0], GLOBALCAMERA.LastPos[1]-ypos
	GLOBALCAMERA.LastPos[0], GLOBALCAMERA.LastPos[1] = xpos, ypos

	sensitivity := 0.1
	xoffset *= sensitivity
	yoffset *= sensitivity

	GLOBALCAMERA.Yaw += xoffset
	GLOBALCAMERA.Yaw = math.Mod(GLOBALCAMERA.Yaw, 360)
	if GLOBALCAMERA.Pitch += yoffset; GLOBALCAMERA.Pitch > 89.0 {
		GLOBALCAMERA.Pitch = 89.0
	} else if GLOBALCAMERA.Pitch < -89.0 {
		GLOBALCAMERA.Pitch = -89.0
	}

	GLOBALCAMERA.CamFront = mgl32.Vec3{
		float32(math.Cos(mgl64.DegToRad(GLOBALCAMERA.Pitch)) * math.Cos(mgl64.DegToRad(GLOBALCAMERA.Yaw))),
		float32(math.Sin(mgl64.DegToRad(GLOBALCAMERA.Pitch))),
		float32(math.Cos(mgl64.DegToRad(GLOBALCAMERA.Pitch)) * math.Sin(mgl64.DegToRad(GLOBALCAMERA.Yaw))),
	}.Normalize()
}

func processInput(window *glfw.Window) {
	camSpeed := float32(20 * DELTATIME)
	if window.GetKey(glfw.KeyW) == glfw.Press {
		GLOBALCAMERA.CamPosition = GLOBALCAMERA.CamPosition.Add(GLOBALCAMERA.CamFront.Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		GLOBALCAMERA.CamPosition = GLOBALCAMERA.CamPosition.Sub(GLOBALCAMERA.CamFront.Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		GLOBALCAMERA.CamPosition = GLOBALCAMERA.CamPosition.Add(mgl32.Vec3.Normalize(GLOBALCAMERA.CamFront.Cross(mgl32.Vec3{0, 1, 0})).Mul(camSpeed))
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		GLOBALCAMERA.CamPosition = GLOBALCAMERA.CamPosition.Sub(mgl32.Vec3.Normalize(GLOBALCAMERA.CamFront.Cross(mgl32.Vec3{0, 1, 0})).Mul(camSpeed))
	}
	if window.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press && MOUSESTATE == 0 {
		MOUSESTATE++
		CHUNKRENDERER.DestroyBlock(&MOUSESTATE)
	}
	if window.GetMouseButton(glfw.MouseButtonRight) == glfw.Press && MOUSESTATE == 0 {
		MOUSESTATE++
		CHUNKRENDERER.CreateBlock(&MOUSESTATE)
	}
}
