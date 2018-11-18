package emath

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

// Integrate is a second order function which returns
// a function approximating derivative of the
// argument fn. d is the infitesimal change in inputs to
// fn.
func Derive(fn RealFunc, d float64) RealFunc {
	return func(x float64) float64 {
		return (fn(x) - fn(x+d)) / d
	}
}
