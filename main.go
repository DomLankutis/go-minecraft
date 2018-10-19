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
	SEED = opensimplex.New32(rand.Int63())

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
		SHADER, err = glhf.NewShader(VERTEXFORMAT, glhf.AttrFormat{
			{"MVP", glhf.Mat4},
		}, VERTEXSHADER, FRAGMENTSHADER)
		if err != nil {
			panic(err)
		}

		texture = glhf.NewTexture(TEXTUREATLAS.Bounds().Dx(), TEXTUREATLAS.Bounds().Dy(),
			false, TEXTUREATLAS.Pix)

		GLOBALCAMERA = utils.InitCamera(WIDTH, HEIGHT, SHADER)

		grass := world.NewCube(SHADER, texture, mgl32.Vec2{0}, [6]bool{})

		CHUNKRENDERER = NewChunkRender(grass, SEED, &GLOBALCAMERA, SHADER, texture)
	})

	shouldClose := false
	for !shouldClose {
		mainthread.Call(func() {
			if win.ShouldClose() {
				shouldClose = true
			}

			currentTime := glfw.GetTime()
			DELTATIME = currentTime - LASTFRAME
			LASTFRAME = currentTime

			glhf.Clear(0.0, 0.4, 0.8, 1)

			processInput(win)
			GLOBALCAMERA.Update()

			CHUNKRENDERER.GetVisibleChunks(4)
			CHUNKRENDERER.Draw()

			win.SwapBuffers()
			glfw.PollEvents()

			//<-FRAMERATE
			FRAMES++
			select {
			case <-tick:
				win.SetTitle(fmt.Sprintf("FPS: %v | DT: %v", FRAMES, DELTATIME))
				FRAMES = 0
			default:
			}
		})
	}
}

func main() {
	if err := loadFiles(); err != nil {
		panic(err)
	}
	mainthread.Run(run)
}
