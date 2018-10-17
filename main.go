package main

import (
	"./glhf"
	"fmt"
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ojrac/opensimplex-go"
	"math/rand"
	"time"

	"./utils"
	"./world"
)

func run() {
	tick := time.Tick(time.Second)

	rand.Seed(time.Now().UnixNano())
	seed = opensimplex.New32(rand.Int63())

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
		win, err = glfw.CreateWindow(WIDTH, HEIGHT, "BtecCraft", nil, nil)
		if err != nil {
			panic(err)
		}

		win.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
		win.SetCursorPosCallback(mouseCallback)

		win.MakeContextCurrent()
		glhf.Init()
	})

	mainthread.Call(func() {
		var err error
		shader, err = glhf.NewShader(vertexFormat, glhf.AttrFormat{
			{"MVP", glhf.Mat4},
		}, vertexShader, fragmentShader)
		if err != nil {
			panic(err)
		}

		texture = glhf.NewTexture(testImg.Bounds().Dx(), testImg.Bounds().Dy(),
		false, testImg.Pix,)

		globalCamera = utils.InitCamera(WIDTH, HEIGHT, shader)

		grass := world.NewCube(shader, texture, mgl32.Vec2{0}, [6]bool{})

		chunkRender = NewChunkRender(grass, seed, &globalCamera, shader, texture)
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

			glhf.Clear(0.0, 0.4, 0.8, 1)


			processInput(win)
			globalCamera.Update()

			chunkRender.GetVisibleChunks(4)
			chunkRender.Draw()

			win.SwapBuffers()
			glfw.PollEvents()

			//<-frameRate
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
