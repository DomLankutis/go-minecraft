package main

import (
	"./eUtils"
)

func loadFiles() error {
	var err error
	vertexShader, err = eUtils.ReadFile("./shaders/vertexShader.vertexshader")
	if err != nil {
		return err
	}

	fragmentShader, err = eUtils.ReadFile("./shaders/fragmentShader.fragmentshader")
	if err != nil {
		return err
	}

	testImg, err = eUtils.LoadImage("./textures/grass.png")
	if err != nil {
		return err
	}

	return nil
}
