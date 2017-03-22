package goi

import (
	"time"

	"github.com/go-gl/gl/v4.1-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	rect          Rectangle
	title         string
	window        *glfw.Window
	viewInitiated bool
	widgets       []*Widget
	moverWidget   *Widget
}

func CreateWindow(width float64, height float64, title string) (*Window, error) {
	glW, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)

	if err != nil {
		return nil, err
	}

	glW.SetFramebufferSizeCallback(updateWindowView)

	glW.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	w := &Window{
		rect:   Rect(0, 0, float64(width), float64(height)),
		title:  title,
		window: glW,
	}

	glW.SetFramebufferSizeCallback(func(_ *glfw.Window, width int, height int) {
		if w.viewInitiated {
			// pop 2
			gl.PopMatrix()

			// pop 1
			gl.PopMatrix()
		}

		// push 1
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Viewport(0, 0, int32(width), int32(height))
		gl.PushMatrix()

		// push 2
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Ortho(0, float64(width), float64(height), 0, -1, 1)
		gl.PushMatrix()

		if !w.viewInitiated {
			draw(w)
		}

		w.viewInitiated = true
	})

	glW.SetCursorPosCallback(func(_ *glfw.Window, x float64, y float64) {
		var moverWidget *Widget

		for _, widget := range w.widgets {
			if widget.rect.Contains(x, y) {
				moverWidget = widget
			}
		}

		if moverWidget != w.moverWidget {
			if w.moverWidget != nil {
				w.moverWidget.handleMleave()
			}

			w.moverWidget = moverWidget
			if moverWidget != nil {
				moverWidget.handleMenter()
			}
		}

		if moverWidget != nil {
			widgetX, widgetY := moverWidget.GetPos()

			moverWidget.handleMover(x-widgetX, y-widgetY)
		}
	})

	return w, nil
}

func (w *Window) Run() {
	for !w.window.ShouldClose() {
		draw(w)

		glfw.PollEvents()

		time.Sleep(20 * time.Microsecond)
	}
}

func (win *Window) AddWidget(w *Widget) {
	win.widgets = append(win.widgets, w)
}

func draw(w *Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	width, height := w.window.GetFramebufferSize()

	if !w.viewInitiated {
		updateWindowView(w.window, width, height)
	}

	ctx := newContext()

	// drawing
	for _, widget := range w.widgets {
		widget.Draw(ctx)
	}

	// gl.Flush()

	w.window.SwapBuffers()
}

func updateWindowView(w *glfw.Window, width int, height int) {
}
