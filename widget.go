package goi

type Widget struct {
	style         *Style
	moverStyle    *Style
	mdownStyle    *Style
	selectedStyle *Style
	widgets       []*Widget
	rect          Rectangle
}

func (w *Widget) Draw(c *Context) {
	c.FillRect(w.rect, w.style.BackgroundColor)
}
