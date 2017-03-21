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
