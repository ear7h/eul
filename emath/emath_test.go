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

func TestIteratedIntegral(t *testing.T) {
	type tcase struct {
		limits [][2]float64
		fn     RealNFunc
		d      []float64
		eps    []float64
		result []float64
	}

	fn := func(tc tcase, t *testing.T) {
		result := IteratedIntegral(tc.limits, tc.fn, tc.d)
		for k, v := range tc.result {
			if !Apprx(result[k], v, tc.eps[k]) {
				t.Fatalf("incorect result[%d] %e not %e +- %e",
					k, result[k], v, tc.eps[k])
			}
		}
	}

	testcases := map[string]tcase{
		"unit cube": {
			limits: [][2]float64{{0, 1}, {0, 1}},
			fn: RealNFunc{
				Fn: func([]float64) []float64 {
					return []float64{1}
				},
				Dmn: 2, Rng: 1,
			},
			d:      []float64{0.001, 0.001},
			eps:    []float64{0.001},
			result: []float64{1.0},
		},
		"unit 3 sphere": {
			limits: [][2]float64{
				{0, 1},
				{0, 2 * math.Pi},
				{0, math.Pi}},
			fn: RealNFunc{
				// f(rho, theta, phi)
				Fn: func(args []float64) []float64 {
					return []float64{args[0] * args[0] * math.Sin(args[2])}
				},
				Dmn: 3, Rng: 1,
			},
			d:      []float64{0.01, 0.01, 0.01},
			eps:    []float64{0.1},
			result: []float64{(4.0 / 3.0) * math.Pi},
		},
	}

	for k, v := range testcases {
		t.Run(k, func(t *testing.T) {
			fn(v, t)
		})
	}
}
