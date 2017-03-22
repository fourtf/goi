package goi

func NewButton() *Widget {
	w := NewWidget()

	w.style.BackgroundColor = ParseColor("#EEE")
	w.style.BorderColor = ParseColor("#666")
	w.style.BorderWidth = 1.0
	w.moverStyle.BackgroundColor = ParseColor("#DDD")
	w.moverStyle.BorderWidth = 3.0
	w.moverStyle.BorderColor = ParseColor("#55F")
	w.mdownStyle.BackgroundColor = ParseColor("#CCC")

	w.updateStyle()

	return w
}
