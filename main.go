// graffinity project main.go
package main

import "fmt"

type Graffinity struct {
	data              map[string]map[string][]float64
	funcs             map[string]func([]float64) float64
	affinityFunc      string
	groupaffinityFunc string
}

func (g Graffinity) calculate() map[string]map[string]float64 {

	type NodeAndData struct {
		name string
		data []float64
	}

	var data = g.data
	var funcs = g.funcs
	var affinityFunc = g.affinityFunc
	fmt.Println(affinityFunc)

	var f = make(map[string][]NodeAndData)

	for namefunc, _ := range funcs {
		for nodename, nodedata := range data {
			var nad = NodeAndData{nodename, nodedata[namefunc]}
			f[namefunc] = append(f[namefunc], nad)
		}
	}

	var nodenames []string
	for n, _ := range data {
		nodenames = append(nodenames, n)
	}

	var matrix = make(map[string]map[string]float64)
	for i := 0; i < len(nodenames); i++ {
		nodename := nodenames[i]
		n := map[string]float64{
			nodename: 0.0,
		}
		matrix[nodename] = n
		for i := 0; i < len(nodenames); i++ {
			othernodename := nodenames[i]
			matrix[nodename][othernodename] = 0.0
		}
	}

	fmt.Println(len(f))
	fmt.Println(len(matrix))

	// end of equivalent of python's constructor

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
