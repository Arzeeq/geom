package geom

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math"
	"os"

	"github.com/fogleman/gg"
)

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func SetPixel(x, y int, canvas *gg.Context) {
	tx, ty := canvas.TransformPoint(float64(x), float64(y))
	canvas.SetPixel(int(tx), int(ty))
}

func SetPixelTr(x, y int, canvas *gg.Context) {
	tr := NewTransferTransform(10, 10)
	sh := NewShearTransform(1.5)
	ro := NewRotateTransform(math.Pi / 4)
	h := NewHomothetyTransform(0.5, 0.5)
	a := Composition(tr, sh, ro, h)
	p := a.Transform(NewPoint2D(float64(x), float64(y)))
	tx, ty := canvas.TransformPoint(p.X(), p.Y())

	canvas.SetPixel(int(tx), int(ty))
}

func SimpleSegmentDrawer(p1, p2 Point2D, canvas *gg.Context) {
	p1x, p1y, p2x, p2y := p1.X(), p1.Y(), p2.X(), p2.Y()
	SetPixel(int(math.Round(p1x)), int(math.Round(p1y)), canvas)
	SetPixel(int(math.Round(p2x)), int(math.Round(p2y)), canvas)
	inverseMode := math.Abs(p1x-p2x) < math.Abs(p1y-p2y)
	if inverseMode {
		p1x, p1y = p1y, p1x
		p2x, p2y = p2y, p2x
	}
	if p1x > p2x {
		p1x, p2x = p2x, p1x
		p1y, p2y = p2y, p1y
	}
	k := (p2y - p1y) / (p2x - p1x)
	b := p1y - k*p1x
	for i := range int(p2x-p1x) + 1 {
		curX := p1x + float64(i)
		curY := curX*k + b
		curY = math.Round(curY)

		if inverseMode {
			curX, curY = curY, curX
		}
		SetPixel(int(curX), int(curY), canvas)
	}
}

func DrawSegmentBresenham(p1, p2 Point2D, canvas *gg.Context) {
	p1x, p1y, p2x, p2y := int(p1.X()), int(p1.Y()), int(p2.X()), int(p2.Y())
	SetPixel(p1x, p1y, canvas)
	SetPixel(p2x, p2y, canvas)
	inverseMode := Abs(p1x-p2x) < Abs(p1y-p2y)
	if inverseMode {
		p1x, p1y = p1y, p1x
		p2x, p2y = p2y, p2x
	}
	if p1x > p2x {
		p1x, p2x = p2x, p1x
		p1y, p2y = p2y, p1y
	}
	m_new := 2 * (p2y - p1y)
	slope_error_new := m_new - (p2x - p1x)

	for x, y := p1x, p1y; x <= p2x; x++ {
		if inverseMode {
			SetPixel(y, x, canvas)
		} else {
			SetPixel(x, y, canvas)
		}

		slope_error_new += m_new

		if slope_error_new >= 0 {
			y++
			slope_error_new -= 2 * (p2x - p1x)
		}
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func c(n, k int) int {
	return fact(n) / (fact(k) * fact(n-k))
}

func BezierCurve(p []Point2D, n int, canvas *gg.Context) {
	if len(p) == 0 {
		return
	}

	for i := range n + 1 {
		t := float64(i) / float64(n)

		t1 := make([]float64, len(p))
		t2 := make([]float64, len(p))
		t1[0], t2[0] = 1, 1
		for j := 1; j < len(p); j++ {
			t1[j] = t1[j-1] * t
			t2[j] = t2[j-1] * (1 - t)
		}

		var ansX, ansY float64
		for j := range len(p) {
			// todo: add factorialss
			ansX += float64(c(len(p)-1, j)) * p[j].X() * t1[j] * t2[len(p)-j-1]
			ansY += float64(c(len(p)-1, j)) * p[j].Y() * t1[j] * t2[len(p)-j-1]
		}
		SetPixel(int(math.Round(ansX)), int(math.Round(ansY)), canvas)
	}
}

func DrawRadarGif() {
	images := make([]image.Image, 0)
	n := 100
	t1 := 0.0
	t2 := 2 * math.Pi
	for i := range n {
		canvas := NewCanvas(200, 200, true)
		canvas.SetColor(color.RGBA{255, 0, 0, 255})
		t := t1 + (t2-t1)*(float64(i)/float64(n))
		p1 := NewPoint2D(0, 0)
		p2 := NewPoint2D(70*math.Cos(t), 70*math.Sin(t))
		DrawSegmentBresenham(p1, p2, canvas)
		images = append(images, canvas.Image())
	}
	CreateGif(images)
}

func CreateGif(images []image.Image) {

	imgs := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	for _, img := range images {
		bounds := img.Bounds()
		dst := image.NewPaletted(bounds, palette.WebSafe)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
		imgs = append(imgs, dst)
		delays = append(delays, 5)
	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: imgs,
		Delay: delays,
	})
}
