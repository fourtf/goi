package main

import (
	"runtime"

	"github.com/fourtf/goi"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := goi.Init()

	defer goi.Terminate()

	w, err := goi.CreateWindow(640, 480, "Testing")
	if err != nil {
		panic(err)
	}

	btn := goi.NewButton()
	w.AddWidget(btn)

	w.Run()

	// select {}
}
