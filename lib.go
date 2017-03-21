package goi

import "github.com/go-gl/glfw/v3.2/glfw"

func Init() error {
	return glfw.Init()
}

func Terminate() {
	glfw.Terminate()
}

func Run() {
}
