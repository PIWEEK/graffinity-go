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

	affinityFunc := "gender_func(x) + 3.5*age_func(x) + 0.1*languages_func(x)"

	g := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc}
	results := g.calculate()
	if results["n1"]["n1"] != 0.0 {
		t.Error("Expected 0.0, got ", results["n1"]["n1"])
	}
	if results["n1"]["n2"] != 4.699007150707624 {
		t.Error("Expected 4.699007150707624, got ", results["n1"]["n2"])
	}
	if results["n2"]["n1"] != 4.699007150707624 {
		t.Error("Expected 4.699007150707624, got ", results["n2"]["n1"])
	}
	if results["n2"]["n2"] != 0.0 {
		t.Error("Expected 0.0, got ", results["n2"]["n2"])
	}
	if results["n2"]["n3"] != 4.317276211268162 {
		t.Error("Expected 4.317276211268162, got ", results["n2"]["n3"])
	}
}
