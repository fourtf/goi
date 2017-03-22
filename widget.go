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

	tmp := *s1
	c := &tmp

	if s2.BackgroundColor != nil {
		c.BackgroundColor = s2.BackgroundColor
	}
	if s2.BorderColor != nil {
		c.BorderColor = s2.BorderColor
	}
	if s2.BorderWidth >= 0 {
		c.BorderWidth = s2.BorderWidth
	}

	w.currentStyle = c
}

func (w *Widget) Draw(c *Context) {
	if w.currentStyle.BackgroundColor != nil {
		c.FillRect(w.rect, w.currentStyle.BackgroundColor)
		c.BorderRect(w.rect, w.currentStyle.BorderColor, w.currentStyle.BorderWidth)
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

func (w *Widget) handleMenter() {
	fmt.Println("enter")

	w.mover = true
	w.updateStyle()
}

func (w *Widget) handleMleave() {
	fmt.Println("leave")

	w.mover = false
	w.updateStyle()
}

func (w *Widget) handleMdown(x float64, y float64) {

}

func (w *Widget) handleMup(x float64, y float64) {

}
