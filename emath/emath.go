package emath

import "fmt"

type RealFunc func(float64) float64

func DefIntegrate(a, b float64,
	fn RealFunc,
	d float64) float64 {

	if d == 0.0 {
		panic("infitesimal is 0")
	}

	sum := float64(0.0)

	// orientation
	o := float64(1.0)
	if a > b {
		a, b = b, a
		o *= -1
	}

	for i := a; i < b; i += d {
		sum += fn(i)
	}

	return o * sum * d
}

/* bin algo
..........
|......... summable
|....|.... summable
|.|..|....
|.|..|.|.. summable
|||..|.|..
||||.|.|..
||||.|||..
||||.||||. summable
|||||||||.
|||||||||| summable
*/

// BinaryIntegrateCb uses a progressive binary partitioning
// algorithm for calculating the integral of a function R1->R1.
// It takes a callback which is called with progressively
// more accurate approximations of the integral
func BinaryIntegrateCb(a, b float64,
	fn RealFunc,
	depth int,
	cb func(float64)) float64 {

	roi := (b - a) // range of integration
	parts := float64(1.0)

	sum := fn(a)

	// placeholder for division
	delta := roi
	// same loop but without calls to cb
	if cb == nil {
		for i := 1; i <= depth; i++ {
			parts *= 2.0
			delta = roi / parts
			for j := float64(1.0); j < parts; j += 2.0 {
				sum += fn(a + j*delta)
			}
		}
	} else {
		cb(sum * roi)
		for i := 1; i <= depth; i++ {
			parts *= 2.0
			delta = roi / parts
			for j := float64(1.0); j < parts; j += 2.0 {
				sum += fn(a + j*delta)
			}
			cb(sum * delta)
		}
	}

	return (sum * delta)
}

// Integrate is a second order function which returns
// a function approximating indefinite integral of the
// argument fn. d is the infitesimal change in inputs to
// fn. Note: this is imply an integral from 0 to x, WITHOUT
// the added C term
func Integrate(fn RealFunc, d float64) RealFunc {

	/*
		note:

		if:
			Intgr:f(x)dx = DefIntgr:0, x, f(y) dy
		=>	F(x) + c = F(x) - F(0)
		=>	c = F(0)
	*/

	return func(x float64) float64 {
		sum := float64(0.0)
		for i := float64(0.0); i < x; i += d {
			sum += fn(i)
		}
		return sum * d
	}
}

// Derive is a second order function which returns
// a function approximating derivative of the
// argument fn. d is the infitesimal change in inputs to
// fn.
func Derive(fn RealFunc, d float64) RealFunc {
	return func(x float64) float64 {
		return (fn(x) - fn(x+d)) / d
	}
}

type RealNFunc struct {
	Fn       func(args []float64) (vec []float64)
	Dmn, Rng int
}

func IteratedIntegral(limits [][2]float64, fn RealNFunc, d []float64) []float64 {
	// check dims
	dmn, rng := fn.Dmn, fn.Rng
	if len(limits) != dmn || len(d) != dmn {
		panic(fmt.Sprintf("incorrect dimensions"))
	}

	cursor := make([]float64, dmn)
	for k, v := range limits {
		cursor[k] = v[0]
	}

	ret := make([]float64, rng)

	firstB := limits[0][1]

	last := len(cursor) - 1
	lastB := limits[last][1]
	lastD := d[last]

	fnfn := fn.Fn

	for cursor[0] < firstB {
		// loop through the innermost integral
		// fmt.Println(cursor)
		for ; cursor[last] < lastB; cursor[last] += lastD {
			add(ret, fnfn(cursor))
		}

		// do the appropriate carries
		for i := last; i > 0; i-- {
			//fmt.Println(cursor[i], limits[i][1], cursor[i] > limits[i][1])
			if cursor[i] > limits[i][1] {
				cursor[i] = limits[i][0]
				cursor[i-1] += d[i-1]
			} else {
				break
			}
		}
	}

	jac := d[0]
	for i := 1; i < len(d); i++ {
		jac *= d[i]
	}

	return muls(ret, jac)
}

// mul muls src to dst in place, with no length checks. returns dst.
func muls(dst []float64, s float64) []float64 {
	for k := range dst {
		dst[k] *= s
	}

	return dst
}

// add adds src to dst in place, with no length checks
func add(dst, src []float64) []float64 {
	for k, v := range src {
		dst[k] += v
	}

	return dst
}
