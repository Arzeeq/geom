package geom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLineWithoutError(t *testing.T) {
	tests := []struct {
		p1, p2 Point2D
	}{
		{
			p1: NewPoint2D(0, 0),
			p2: NewPoint2D(1, 0),
		},
		{
			p1: NewPoint2D(0, 0),
			p2: NewPoint2D(0, 1),
		},
		{
			p1: NewPoint2D(-1e9, -1e9),
			p2: NewPoint2D(1e9, 1e9),
		},
		{
			p1: NewPoint2D(3, 4),
			p2: NewPoint2D(10, -9),
		},
		{
			p1: NewPoint2D(15, 3),
			p2: NewPoint2D(8, 4),
		},
		{
			p1: NewPoint2D(1, 2),
			p2: NewPoint2D(4, 3),
		},
	}
	for _, test := range tests {
		_, err1 := NewLine(test.p1, test.p2)
		_, err2 := NewLine(test.p2, test.p1)

		require.NoError(t, err1)
		require.NoError(t, err2)
	}
}

func TestNewLineWithError(t *testing.T) {
	tests := []struct {
		p1, p2 Point2D
	}{
		{
			p1: NewPoint2D(0, 0),
			p2: NewPoint2D(0, 0),
		},
		{
			p1: NewPoint2D(1, 1),
			p2: NewPoint2D(1, 1),
		},
		{
			p1: NewPoint2D(1e9, 1e9),
			p2: NewPoint2D(1e9, 1e9),
		},
		{
			p1: NewPoint2D(-1e9, -1e9),
			p2: NewPoint2D(-1e9, -1e9),
		},
	}
	for _, test := range tests {
		_, err1 := NewLine(test.p1, test.p2)
		_, err2 := NewLine(test.p2, test.p1)

		require.ErrorIs(t, err1, ErrLineCreation)
		require.ErrorIs(t, err2, ErrLineCreation)
	}
}

func TestIsParallel(t *testing.T) {
	tests := map[string][4]Point2D{
		"horizontal lines": {
			NewPoint2D(0, 0),
			NewPoint2D(5, 0),
			NewPoint2D(0, 5),
			NewPoint2D(5, 5),
		},
		"vertical lines": {
			NewPoint2D(0, 0),
			NewPoint2D(0, 5),
			NewPoint2D(5, 0),
			NewPoint2D(5, 5),
		},
		"random lines": {
			NewPoint2D(3, 4),
			NewPoint2D(9, 12),
			NewPoint2D(-2, -1),
			NewPoint2D(4, 7),
		},
		"huge numbers lines": {
			NewPoint2D(-1e9, -1e9),
			NewPoint2D(1e9, 1e9),
			NewPoint2D(-1e9+1, -1e9+1),
			NewPoint2D(1e9+1, 1e9+1),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l1, err1 := NewLine(test[0], test[1])
			l2, err2 := NewLine(test[2], test[3])
			require.NoError(t, err1)
			require.NoError(t, err2)
			require.True(t, IsParallel(l1, l2))
		})
	}
}

func TestIsNotParallel(t *testing.T) {
	p1 := NewPoint2D(0, 0)
	p2 := NewPoint2D(0, 10)
	p3 := NewPoint2D(0, 0)
	p4 := NewPoint2D(10, 0)
	l1, err1 := NewLine(p1, p2)
	l2, err2 := NewLine(p3, p4)

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.False(t, IsParallel(l1, l2))
}
