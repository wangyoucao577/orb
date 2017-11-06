package planar

import (
	"math"
	"testing"

	"github.com/paulmach/orb/geo"
)

var epsilon = 1e-6

func TestDistanceFrom_MultiPoint(t *testing.T) {
	mp := geo.MultiPoint{{0.0}, {1, 1}, {2, 2}}

	fromPoint := geo.Point{3, 2}
	if distance := DistanceFrom(mp, fromPoint); distance != 1 {
		t.Errorf("distance incorrect: %v != %v", distance, 1)
	}
}

func TestDistanceFrom_LineString(t *testing.T) {
	ls := geo.LineString{{0, 0}, {0, 3}, {4, 3}, {4, 0}}

	cases := []struct {
		name   string
		point  geo.Point
		result float64
	}{
		{
			point:  geo.NewPoint(4.5, 1.5),
			result: 0.5,
		},
		{
			point:  geo.NewPoint(0.4, 1.5),
			result: 0.4,
		},
		{
			point:  geo.NewPoint(-0.3, 1.5),
			result: 0.3,
		},
		{
			point:  geo.NewPoint(0.3, 2.8),
			result: 0.2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := DistanceFrom(ls, tc.point)
			if math.Abs(d-tc.result) > epsilon {
				t.Errorf("incorrect distance: %v != %v", d, tc.result)
			}
		})
	}
}

func TestDistanceFrom_Polygon(t *testing.T) {
	r1 := geo.Ring{{0, 0}, {3, 0}, {3, 3}, {0, 3}, {0, 0}}
	r2 := geo.Ring{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}

	poly := append(geo.NewPolygon(), r1, r2)

	cases := []struct {
		name   string
		point  geo.Point
		result float64
	}{
		{
			name:   "outside",
			point:  geo.NewPoint(-1, 2),
			result: 1,
		},
		{
			name:   "inside",
			point:  geo.NewPoint(0.4, 2),
			result: 0.4,
		},
		{
			name:   "in hole",
			point:  geo.NewPoint(1.3, 1.4),
			result: 0.3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if d := DistanceFrom(poly, tc.point); math.Abs(d-tc.result) > epsilon {
				t.Errorf("incorrect distance: %v != %v", d, tc.result)
			}
		})
	}
}
