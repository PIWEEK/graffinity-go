// tests
package main

import "testing"
import "math"
import "fmt"

func TestBasicCalculation(t *testing.T) {
	data := map[string]map[string][]float32{
		"n1": {
			"gender":    []float32{0},
			"age":       []float32{36},
			"languages": []float32{2, 5, 6},
		},
		"n2": {
			"gender":    []float32{1},
			"age":       []float32{33},
			"languages": []float32{2, 6},
		},
		"n3": {
			"gender":    []float32{1},
			"age":       []float32{25},
			"languages": []float32{1, 6, 9, 10},
		},
		"n4": {
			"gender":    []float32{1},
			"age":       []float32{28},
			"languages": []float32{1},
		},
	}

	genderFunc := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	ageFunc := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	languagesFunc := func(x []float32) float32 { return 5 * float32(len(x)-len(RemoveDuplicates(x))) }

	funcs := map[string]func([]float32) float32{
		"gender":    genderFunc,
		"age":       ageFunc,
		"languages": languagesFunc,
	}

	affinityFunc := func(dictValues map[string]float32) float32 {
		return dictValues["gender"] + 3.5*dictValues["age"] + 0.1*dictValues["languages"]
	}

	groupaffinityFunc := func(dictValues map[string]float32) float32 {
		return dictValues["gender"] + 3.5*dictValues["age"] + 0.1*dictValues["languages"]
	}

	g := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
	results := g.calculate()
	fmt.Println("results", results["n1"])

	if results["n1"]["n2"]["total"] != 4.347826 {
		t.Error("n1 n2 Expected 4.347826, got ", results["n1"]["n2"]["total"])
	}
	if results["n2"]["n1"]["total"] != 4.347826 {
		t.Error("n2 n1 Expected 4.347826, got ", results["n2"]["n1"]["total"])
	}
	if results["n2"]["n3"]["total"] != 4.5172415 {
		t.Error("n2 n3 Expected 4.5172415, got ", results["n2"]["n3"]["total"])
	}
	g2 := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
	results2 := g2.calculatefornode("n1")
	fmt.Println("results2", results2["n1"])

	if results2["n1"]["n2"]["total"] != 4.347826 {
		t.Error("n1 n2 Expected 4.347826, got ", results2["n1"]["n2"]["total"])
	}
	g3 := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
	results3 := g3.calculateforgroup([]string{"n1", "n2", "n3"})
	if results3 != 5.43242 {
		t.Error("n1 n2 b3 group affinity Expected 5.43242, got ", results3)
	}
	fmt.Println("group affinity", results3)
}
