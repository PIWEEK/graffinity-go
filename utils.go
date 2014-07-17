// utils
package main

import "math"
import "math/rand"
import "time"

func RemoveDuplicates(a []float32) []float32 {
	result := []float32{}
	seen := map[float32]float32{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func newRsdv() func(float32) float32 {
	var n, a, q float32
	return func(x float32) float32 {
		n++
		a1 := a + (x-a)/n
		q, a = q+(x-a)*(x-a1), a1
		return float32(math.Sqrt(float64(q) / float64(n)))
	}
}

func Stdev(v []float32) (m float32) {
	r := newRsdv()
	var result float32
	for _, x := range v {
		result = r(x)
	}
	return result
}

func Mean(v []float32) (m float32) {
	// an algorithm that attempts to retain accuracy
	// with widely different values.
	var parts []float32
	for _, x := range v {
		var i int
		for _, p := range parts {
			sum := p + x
			var err float32
			switch ax, ap := float32(math.Abs(float64(x))), float32(math.Abs(float64(p))); {
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
	var sum float32
	for _, x := range parts {
		sum += x
	}
	return sum / float32(len(v))
}

func Range(args ...int) []float32 {
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
	out := []float32{}
	for i := start; i < stop; i += stride {
		out = append(out, float32(i))
	}
	return out
}

func Choice(arr []float32) float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	pos := rand.Intn(len(arr))
	return arr[pos]
}

func Sample(arr []float32, quantity int) []float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	ret := make([]float32, quantity)
	for i := 0; i < quantity; i++ {
		ret[i] = arr[rand.Intn(len(arr))]
	}
	return ret
}

func RandInt(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}
