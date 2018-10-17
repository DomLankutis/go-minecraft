package main

import (
	"./utils"
)

func loadFiles() error {
	var err error
	vertexShader, err = utils.ReadFile("./shaders/vertexShader.glsl")
	if err != nil {
		return err
	}

	fragmentShader, err = utils.ReadFile("./shaders/fragmentShader.glsl")
	if err != nil {
		return err
	}

	testImg, err = utils.LoadImage("./textures/grass.png")
	if err != nil {
		return err
	}

	return nil
}
