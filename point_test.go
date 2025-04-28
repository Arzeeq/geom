package geom

import (
	"testing"

	"github.com/Arzeeq/geom/pkg/gen"
	"github.com/stretchr/testify/require"
)

func TestNewPoint2D(t *testing.T) {
	x, y := gen.RandFloat64n(10), gen.RandFloat64n(10)

	p := NewPoint2D(x, y)

	require.Equal(t, x, p.X())
	require.Equal(t, y, p.Y())
}

func TestIsPointInRect(t *testing.T) {
	type testCase struct {
		points           []Point2D
		vertex1, vertex2 Point2D
	}

	var testCases = map[string]testCase{
		"points inside rect": {
			points: []Point2D{
				NewPoint2D(1, 1),
				NewPoint2D(1, 9),
				NewPoint2D(9, 1),
				NewPoint2D(9, 9),
				NewPoint2D(5, 5),
			},
			vertex1: NewPoint2D(0, 0),
			vertex2: NewPoint2D(10, 10),
		},
		"another pair of oposite vertexies": {
			points: []Point2D{
				NewPoint2D(1, 1),
				NewPoint2D(1, 9),
				NewPoint2D(9, 1),
				NewPoint2D(9, 9),
				NewPoint2D(5, 5),
			},
			vertex1: NewPoint2D(0, 10),
			vertex2: NewPoint2D(10, 0),
		},
		"negative coordinates": {
			points: []Point2D{
				NewPoint2D(-15, -15),
			},
			vertex1: NewPoint2D(-10, -10),
			vertex2: NewPoint2D(-20, -20),
		},
		"points lie on borders": {
			points: []Point2D{
				NewPoint2D(0, 2),
				NewPoint2D(2, 0),
				NewPoint2D(5, 2),
				NewPoint2D(2, 5),
			},
			vertex1: NewPoint2D(0, 0),
			vertex2: NewPoint2D(5, 5),
		},
		"points lie on vertexies": {
			points: []Point2D{
				NewPoint2D(0, 0),
				NewPoint2D(0, 5),
				NewPoint2D(5, 0),
				NewPoint2D(5, 5),
			},
			vertex1: NewPoint2D(0, 0),
			vertex2: NewPoint2D(5, 5),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, p := range test.points {
				require.True(t, IsPointInRect(p, test.vertex1, test.vertex2))
				require.True(t, IsPointInRect(p, test.vertex2, test.vertex1))
			}
		})
	}
}

func TestIsNotInRect(t *testing.T) {
	vertex1 := NewPoint2D(-10, -10)
	vertex2 := NewPoint2D(10, 10)
	points := []Point2D{
		NewPoint2D(-1e9, 0),
		NewPoint2D(1e9, 0),
		NewPoint2D(0, -1e9),
		NewPoint2D(0, 1e9),
		NewPoint2D(15, 15),
		NewPoint2D(-15, -15),
		NewPoint2D(-15, 15),
		NewPoint2D(15, -15),
	}
	for _, p := range points {
		require.False(t, IsPointInRect(p, vertex1, vertex2))
		require.False(t, IsPointInRect(p, vertex2, vertex1))
	}
}

func TestIsPointInRectWithRandomPoints(t *testing.T) {
	n, min, max := 10000, -1e9, 1e9
	vertex1 := NewPoint2D(min, min)
	vertex2 := NewPoint2D(max, max)

	for range n {
		p := NewPoint2D(gen.RandFloat64Between(min, max), gen.RandFloat64Between(min, max))
		require.True(t, IsPointInRect(p, vertex1, vertex2))
		require.True(t, IsPointInRect(p, vertex2, vertex1))
	}
}

func TestIsEqualPoints(t *testing.T) {
	p1, p2 := NewPoint2D(5, 5), NewPoint2D(6, 6)
	require.True(t, IsEqualPoints(p1, p1))
	require.False(t, IsEqualPoints(p1, p2))
}
