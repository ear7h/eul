package emath

import (
	"math"
	"testing"
)

func Apprx(y0, y1, eps float64) bool {
	return math.Abs(y1-y0) <= eps
}

func TestIntegrate(t *testing.T) {
	type tcase struct {
		fn     RealFunc
		d      float64
		x      float64
		eps    float64
		result float64
	}

	fn := func(tc tcase, t *testing.T) {
		result := Integrate(tc.fn, tc.d)(tc.x)
		if !Apprx(result, tc.result, tc.eps) {
			t.Fatalf("incorect value %f not %f +- %f", result, tc.result, tc.eps)
		}
	}

	testcases := map[string]tcase{
		"line": {
			fn:     func(x float64) float64 { return x },
			d:      0.0001,
			x:      1,
			eps:    0.001,
			result: 0.5,
		},
		"sin": {
			fn:     func(x float64) float64 { return math.Sin(x) },
			d:      0.0001,
			x:      2 * math.Pi,
			eps:    0.001,
			result: 0.0,
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}

func TestDefIntegrate(t *testing.T) {
	type tcase struct {
		a      float64
		b      float64
		fn     RealFunc
		d      float64
		eps    float64
		result float64
	}

	fn := func(tc tcase, t *testing.T) {
		result := DefIntegrate(tc.a, tc.b, tc.fn, tc.d)
		if !Apprx(result, tc.result, tc.eps) {
			t.Fatalf("incorect value %e not %e +- %e", result, tc.result, tc.eps)
		}
	}

	testcases := map[string]tcase{
		"line": {
			a:      0,
			b:      1,
			d:      0.0001,
			fn:     func(x float64) float64 { return x },
			eps:    0.001,
			result: 0.5,
		},
		"line 1": {
			a:      1,
			b:      0,
			d:      0.0001,
			fn:     func(x float64) float64 { return x },
			eps:    0.001,
			result: -0.5,
		},
		"sin": {
			a:      0,
			b:      math.Pi,
			fn:     math.Sin,
			d:      0.001,
			eps:    0.001,
			result: -math.Cos(math.Pi) - -math.Cos(0),
		},
		"sin 1": {
			a:      math.Pi,
			b:      0,
			fn:     math.Sin,
			d:      0.001,
			eps:    0.001,
			result: -math.Cos(0) - -math.Cos(math.Pi),
		},
		"exp": {
			a:      1,
			b:      10,
			fn:     math.Exp,
			d:      0.00001,
			eps:    1.0,
			result: math.Exp(10) - math.Exp(1),
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
