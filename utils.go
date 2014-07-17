// utils
package main

import "math"
import "math/rand"
import "time"

func RemoveDuplicates(a []float64) []float64 {
	result := []float64{}
	seen := map[float64]float64{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func newRsdv() func(float64) float64 {
	var n, a, q float64
	return func(x float64) float64 {
		n++
		a1 := a + (x-a)/n
		q, a = q+(x-a)*(x-a1), a1
		return math.Sqrt(q / n)
	}
}

func Stdev(v []float64) (m float64) {
	r := newRsdv()
	var result float64
	for _, x := range []float64{2, 4, 4, 4, 5, 5, 7, 9} {
		result = r(x)
	}
	return result
}

func Mean(v []float64) (m float64) {
	// an algorithm that attempts to retain accuracy
	// with widely different values.
	var parts []float64
	for _, x := range v {
		var i int
		for _, p := range parts {
			sum := p + x
			var err float64
			switch ax, ap := math.Abs(x), math.Abs(p); {
			case ax < ap:
				err = x - (sum - p)
			case ap < ax:
				err = p - (sum - x)
			}
			if err != 0 {
				parts[i] = err
				i++
			}
			x = sum
		}
		parts = append(parts[:i], x)
	}
	var sum float64
	for _, x := range parts {
		sum += x
	}
	return sum / float64(len(v))
}

func Range(args ...int) []float64 {
	nargs := len(args)
	start, stop, stride := 0, 0, 0
	switch nargs {
	case 1:
		start = 0
		stop = args[0]
		stride = 1
	case 2:
		start = args[0]
		stop = args[1]
		stride = 1
	case 3:
		start = args[0]
		stop = args[1]
		stride = args[2]
	default:
		panic("boo")
	}
	out := []float64{}
	for i := start; i < stop; i += stride {
		out = append(out, float64(i))
	}
	return out
}

func Choice(arr []float64) float64 {
	rand.Seed(time.Now().Unix())
	pos := rand.Intn(len(arr))
	return arr[pos]
}

func Sample(arr []float64, quantity int) []float64 {
	rand.Seed(time.Now().Unix())
	ret := make([]float64, quantity)
	for i := 0; i < quantity; i++ {
		ret[i] = arr[rand.Intn(len(arr))]
	}
	return ret
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
