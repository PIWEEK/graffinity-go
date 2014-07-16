// graffinity project main.go
package main

import "math"

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

type Graffinity struct {
	data         map[string]map[string][]float64
	funcs        map[string]func([]float64) float64
	affinityFunc func(x []float64) float64
}

func (g Graffinity) calculate() map[string]map[string]float64 {
	ret := map[string]map[string]float64{
		"n1": {
			"n1": 33.0,
			"n2": 33.0,
		},
		"n2": {
			"n1": 33.0,
			"n2": 33.0,
			"n3": 33.0,
		},
	}
	return ret
}