package goi

type Rectangle struct {
	Pos  Point
	Size Size
}

func Rect(x float64, y float64, w float64, h float64) Rectangle {
	return Rectangle{
		Pos: Point{
			X: x,
			Y: y,
		},
		Size: Size{
			Width:  w,
			Height: h,
		},
	}
}

func (rect *Rectangle) Contains(x float64, y float64) bool {
	return rect.Pos.X <= x &&
		rect.Pos.Y <= y &&
		rect.Pos.X+rect.Size.Width >= x &&
		rect.Pos.Y+rect.Size.Height >= y
}

func (rect *Rectangle) ContainsPt(pt Point) bool {
	return rect.Contains(pt.X, pt.Y)
}
