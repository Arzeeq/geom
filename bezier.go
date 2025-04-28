package geom

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func c(n, k int) int {
	return fact(n) / (fact(k) * fact(n-k))
}

// p - control points, n - approximation degree
func (cvs *Canvas) BezierCurve(p []Point2D, n int) {
	if len(p) == 0 {
		return
	}

	for i := 0; i <= n; i++ {
		t := float64(i) / float64(n)

		t1 := make([]float64, len(p))
		t2 := make([]float64, len(p))
		t1[0], t2[0] = 1, 1
		for j := 1; j < len(p); j++ {
			t1[j] = t1[j-1] * t
			t2[j] = t2[j-1] * (1 - t)
		}

		var x, y float64
		for j := range len(p) {
			x += float64(c(len(p)-1, j)) * p[j].X() * t1[j] * t2[len(p)-j-1]
			y += float64(c(len(p)-1, j)) * p[j].Y() * t1[j] * t2[len(p)-j-1]
		}
		if i == 0 {
			cvs.MoveTo(x, y)
		} else {
			cvs.LineTo(x, y)
		}
	}
}
