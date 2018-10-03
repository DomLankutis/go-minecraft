package main

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"time"

	"./Utils"
	"./shape"
)

func run() {
	tick := time.Tick(time.Second)
	var win *glfw.Window

	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	mainthread.Call(func() {
		glfw.Init()

		glfw.WindowHint(glfw.ContextVersionMajor, 4)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.DoubleBuffer, 0)
		glfw.WindowHint(glfw.Resizable, glfw.False)

		var err error
		win, err = glfw.CreateWindow(WIDTH, HEIGHT, "test", nil, nil)
		if err != nil {
			panic(err)
		}

		win.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
		win.SetCursorPosCallback(mouseCallback)

		win.MakeContextCurrent()
		glhf.Init()
	})

	mainthread.Call(func() {
		shader, err := glhf.NewShader(vertexFormat, glhf.AttrFormat{
			{"colorHue", glhf.Vec4},
			{"MVP", glhf.Mat4},
		}, vertexShader, fragmentShader)
		if err != nil {
			panic(err)
		}

		texture := glhf.NewTexture(testImg.Bounds().Dx(), testImg.Bounds().Dy(),
			false, testImg.Pix,)

		cube = shape.NewCube(shader, texture, 16, mgl32.Vec2{0, 0})

		globalCamera = Utils.InitCamera(WIDTH, HEIGHT, shader)
	})

	shouldClose := false
	for !shouldClose {
		mainthread.Call(func() {
			if win.ShouldClose() {
				shouldClose = true
			}

			currentTime := glfw.GetTime()
			deltaTime = currentTime - lastFrame
			lastFrame = currentTime

			glhf.Clear(0.2, 0.4, 1, 1)

			processInput(win)

			go globalCamera.Update()
			chunks = shape.GenerateChunkAroundPlayer(globalCamera, &cube, chunkPos, seed)

			for _, chunk := range chunks {
				chunk.Draw(globalCamera)
			}
			win.SwapBuffers()
			glfw.PollEvents()

			<-frameRate
			frames++
			select {
			case <- tick:
				win.SetTitle(fmt.Sprintf("FPS: %v | DT: %v", frames, deltaTime))
				frames = 0
			default:
			}
		})
	}
}

func main () {
	if err := loadFiles(); err != nil {
		panic(err)
	}
	mainthread.Run(run)
}
