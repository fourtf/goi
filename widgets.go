package goi

func NewWidget() *Widget {
	w := &Widget{
		rect:          Rect(0, 0, 60, 20),
		style:         &Style{},
		moverStyle:    &Style{},
		mdownStyle:    &Style{},
		selectedStyle: &Style{},
	}

	return w
}

func NewButton() *Widget {
	w := NewWidget()

	w.style.BackgroundColor = ParseColor("#F00")

	return w
}
