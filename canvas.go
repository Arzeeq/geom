package geom

import (
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type Canvas struct {
	gg.Context
	w, h  int     // pixels
	cage  int     // pixels
	scale float64 // scale of 1 cage
}

func NewCanvas(w, h, cage int, scale float64) *Canvas {
	// create canvas
	canvas := &Canvas{
		Context: *gg.NewContext(w, h),
		w:       w,
		h:       h,
		cage:    cage,
		scale:   scale,
	}

	// fill canvas with white color
	canvas.SetRGB(1, 1, 1)
	canvas.Clear()

	// translate coordinates into center and reverse y
	canvas.Translate(float64(w)/2, float64(h)/2)
	canvas.Scale(1, -1)

	// drawing minor cages
	for x := 0; x < w/2; x += cage {
		fx := float64(x) + 0.5
		fy := float64(h / 2)
		canvas.DrawLine(fx, -fy, fx, fy)
		canvas.DrawLine(-fx, -fy, -fx, fy)
	}
	for y := 0; y < h/2; y += cage {
		fy := float64(y) + 0.5
		canvas.DrawLine(-float64(w)/2, fy, float64(w)/2, fy)
		canvas.DrawLine(-float64(w)/2, -fy, float64(w)/2, -fy)
	}
	canvas.SetLineWidth(1)
	canvas.SetRGBA(0, 0, 0, 0.25)
	canvas.Stroke()

	// drawing major cages
	for x := 0; x < w/2; x += 5 * cage {
		fx := float64(x) + 0.5
		fy := float64(h / 2)
		canvas.DrawLine(fx, -fy, fx, fy)
		canvas.DrawLine(-fx, -fy, -fx, fy)
	}
	for y := 0; y < h/2; y += 5 * cage {
		fx := float64(w) / 2
		fy := float64(y) + 0.5
		canvas.DrawLine(-fx, fy, fx, fy)
		canvas.DrawLine(-fx, -fy, fx, -fy)
	}
	canvas.SetLineWidth(1)
	canvas.SetRGBA(0, 0, 0, 0.5)
	canvas.Stroke()

	// draw axis X
	canvas.DrawLine(-float64(w)/2, 0, float64(w)/2, 0)
	canvas.DrawLine(float64(w)/2, 0, float64(w)/2-float64(cage), -float64(cage))
	canvas.DrawLine(float64(w)/2, 0, float64(w)/2-float64(cage), float64(cage))
	canvas.SetLineWidth(2)
	canvas.SetRGBA(0, 0, 0, 0.75)
	canvas.Stroke()

	// draw axis Y
	canvas.DrawLine(0, -float64(h)/2, 0, float64(h)/2)
	canvas.DrawLine(0, float64(h)/2, -float64(cage), float64(h)/2-float64(cage))
	canvas.DrawLine(0, float64(h)/2, float64(cage), float64(h)/2-float64(cage))
	canvas.SetLineWidth(2)
	canvas.SetRGBA(0, 0, 0, 0.75)
	canvas.Stroke()

	// sign axis with letter
	canvas.Scale(1, -1)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(font, &truetype.Options{Size: 20})
	canvas.SetFontFace(face)
	canvas.SetRGBA(0, 0, 0, 1)
	canvas.DrawStringAnchored("X", float64(w)/2-1.5*float64(cage), float64(cage), 0.5, 0.5)
	canvas.DrawStringAnchored("Y", -float64(cage), -(float64(h)/2 - 1.5*float64(cage)), 0.5, 0.5)
	// sign scale X and Y
	face = truetype.NewFace(font, &truetype.Options{Size: 14})
	canvas.SetFontFace(face)
	canvas.DrawStringAnchored(strconv.FormatFloat(scale, 'f', -1, 64), float64(cage), float64(cage)/2, 0.5, 0.5)
	canvas.DrawStringAnchored(strconv.FormatFloat(scale, 'f', -1, 64), -float64(cage)/2, -float64(cage), 0.5, 0.5)
	canvas.Scale(1, -1)

	// scale coordinates
	canvas.Scale(float64(cage)/scale, float64(cage)/scale)

	return canvas
}

func (c *Canvas) SetPixel(x, y int) {
	tx := c.w/2 + x
	ty := c.h/2 - y
	c.Context.SetPixel(tx, ty)
}
