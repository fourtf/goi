package goi

import "github.com/go-gl/gl/v4.1-compatibility/gl"

type Context struct {
	size Size
}

func newContext() *Context {
	return &Context{}
}

func (ctx *Context) FillRect(r Rectangle, c *Color) {
	gl.Begin(gl.TRIANGLES)

	gl.Color3ub(c.R, c.G, c.B)

	gl.Vertex2d(r.Pos.X, r.Pos.Y)
	gl.Vertex2d(r.Pos.X+r.Size.Width, r.Pos.Y)
	gl.Vertex2d(r.Pos.X+r.Size.Width, r.Pos.Y+r.Size.Height)

	gl.Vertex2d(r.Pos.X, r.Pos.Y)
	gl.Vertex2d(r.Pos.X, r.Pos.Y+r.Size.Height)
	gl.Vertex2d(r.Pos.X+r.Size.Width, r.Pos.Y+r.Size.Height)

	gl.End()
}

func (ctx *Context) BorderRect(r Rectangle, c *Color, width float64) {
	// gl.Enable(gl.LINE_SMOOTH)
	gl.LineWidth(float32(width))

	gl.Color3ub(c.R, c.G, c.B)

	gl.Begin(gl.LINE_LOOP)

	gl.Vertex2d(r.Pos.X, r.Pos.Y)
	gl.Vertex2d(r.Pos.X+r.Size.Width, r.Pos.Y)
	gl.Vertex2d(r.Pos.X+r.Size.Width, r.Pos.Y+r.Size.Height)
	gl.Vertex2d(r.Pos.X, r.Pos.Y+r.Size.Height)

	gl.End()
}
