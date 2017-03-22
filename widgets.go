package goi

func NewButton() *Widget {
	w := NewWidget()

	w.style.BackgroundColor = ParseColor("#F00")
	w.moverStyle.BackgroundColor = ParseColor("#0F0")
	w.mdownStyle.BackgroundColor = ParseColor("#00F")

	w.updateStyle()

	return w
}
