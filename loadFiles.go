package main

import (
	"./utils"
)

func loadFiles() error {
	var err error
	VERTEXSHADER, err = utils.ReadFile("./shaders/vertexShader.glsl")
	if err != nil {
		return err
	}

	FRAGMENTSHADER, err = utils.ReadFile("./shaders/fragmentShader.glsl")
	if err != nil {
		return err
	}

	TEXTUREATLAS, err = utils.LoadImage("./textures/grass.png")
	if err != nil {
		return err
	}

	return nil
}
