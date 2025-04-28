package geom

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func isEqual(a, b AffineTransform2D) bool {
	for i := range 3 {
		for j := range 3 {
			if math.Abs(a[i][j]-b[i][j]) > EPS {
				return false
			}
		}
	}
	return true
}

func TestInverse_Success(t *testing.T) {
	testcases := []AffineTransform2D{
		{
			{1, 2, 3},
			{0, 1, 4},
			{5, 6, 0},
		},
		{
			{1, 2, 3},
			{0, 4, 2},
			{5, 2, 1},
		},
		{
			{1, -2, 3},
			{0, 4, -1},
			{5, 0, 0},
		},
	}

	for i, tt := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inTr1, err := tt.Inverse()
			require.NoError(t, err)

			inTr2, err := inTr1.Inverse()
			require.NoError(t, err)

			require.True(t, isEqual(tt, inTr2))
		})
	}
}

func TestInverse_Error(t *testing.T) {
	testcases := []AffineTransform2D{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{2, -1, 3},
			{2, -1, 3},
			{5, 2, 1},
		},
		{
			{0, 0, 0},
			{0, 4, -1},
			{5, 0, 0},
		},
	}

	for i, tt := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := tt.Inverse()
			require.Error(t, err)
		})
	}
}
