package goi

import (
	"fmt"
)

type Widget struct {
	moverWidget   *Widget
	style         *Style
	moverStyle    *Style
	mdownStyle    *Style
	selectedStyle *Style
	currentStyle  *Style
	widgets       []*Widget
	rect          Rectangle
	mover         bool
	mdown         bool
}

func NewWidget() *Widget {
	w := &Widget{
		rect:          Rect(40, 40, 60, 20),
		style:         &Style{},
		moverStyle:    &Style{},
		mdownStyle:    &Style{},
		selectedStyle: &Style{},
		currentStyle:  &Style{},
	}

	return w
}

func (w *Widget) GetPos() (x float64, y float64) {
	return w.rect.Pos.X, w.rect.Pos.Y
}

func (w *Widget) updateStyle() {
	s1 := w.style
	var s2 *Style

	if w.mdown {
		s2 = w.mdownStyle
	} else if w.mover {
		s2 = w.mdownStyle
	} else {
		w.currentStyle = s1
		return
	}

	x := *s1
	c := &x

	if s2.BackgroundColor != nil {
		c.BackgroundColor = s2.BackgroundColor
	}

	w.currentStyle = c
}

func (w *Widget) Draw(c *Context) {
	if w.currentStyle.BackgroundColor != nil {
		c.FillRect(w.rect, w.currentStyle.BackgroundColor)
	}
}

func (w *Widget) handleMover(x float64, y float64) {
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
		moverWidget.handleMenter()
	}

	if moverWidget != nil {
		widgetX, widgetY := moverWidget.GetPos()

		moverWidget.handleMover(x-widgetX, y-widgetY)
	}
}

func (w *Widget) handleMleave() {
	fmt.Println("enter")
	w.updateStyle()
}

func (w *Widget) handleMenter() {
	fmt.Println("leave")
	w.updateStyle()
}

func (w *Widget) handleMdown(x float64, y float64) {

}

func (w *Widget) handleMup(x float64, y float64) {

}
