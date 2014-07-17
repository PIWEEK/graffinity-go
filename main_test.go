// tests
package main

import "testing"
import "math"

func TestBasicCalculation(t *testing.T) {
	data := map[string]map[string][]float64{
		"n1": {
			"gender":    []float64{0},
			"age":       []float64{36},
			"languages": []float64{2, 5, 6},
		},
		"n2": {
			"gender":    []float64{1},
			"age":       []float64{33},
			"languages": []float64{2, 6},
		},
		"n3": {
			"gender":    []float64{1},
			"age":       []float64{25},
			"languages": []float64{1, 6, 9, 10},
		},
		"n4": {
			"gender":    []float64{1},
			"age":       []float64{28},
			"languages": []float64{1},
		},
	}

	genderFunc := func(x []float64) float64 { return math.Abs(Mean(x)-Stdev(x)) / Mean(x) }
	ageFunc := func(x []float64) float64 { return math.Abs(Mean(x)-Stdev(x)) / Mean(x) }
	languagesFunc := func(x []float64) float64 { return 5 * float64(len(x)-len(RemoveDuplicates(x))) }

	funcs := map[string]func([]float64) float64{
		"gender":    genderFunc,
		"age":       ageFunc,
		"languages": languagesFunc,
	}

	affinityFunc := func(dictValues map[string]float64) float64 {
		return dictValues["gender"] + 3.5*dictValues["age"] + 0.1*dictValues["languages"]
	}

	g := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc}
	results := g.calculate()

	if results["n1"]["n2"] != 4.3478260869565215 {
		t.Error("n1 n2 Expected 4.3478260869565215, got ", results["n1"]["n2"])
	}
	if results["n2"]["n1"] != 4.3478260869565215 {
		t.Error("n2 n1 Expected 4.3478260869565215, got ", results["n2"]["n1"])
	}
	if results["n2"]["n3"] != 4.517241379310345 {
		t.Error("n2 n3 Expected 4.517241379310345, got ", results["n2"]["n3"])
	}
}
