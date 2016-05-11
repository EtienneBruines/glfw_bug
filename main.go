package main

import (
	"log"
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"
)

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(1920, 1200, "Hello Bug", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()

	log.Println("Sleeping to keep the invisible window open")
	time.Sleep(time.Second * 10)
}
