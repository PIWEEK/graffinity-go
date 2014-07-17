// graffinity project main.go
package main

type graffinityfunc func(x []float64) float64

type NodeAndData struct {
	name string
	data []float64
}

type Graffinity struct {
	data              map[string]map[string][]float64
	funcs             map[string]func([]float64) float64
	affinityFunc      func(map[string]float64) float64
	groupaffinityFunc string
}

func (g Graffinity) calculate() map[string]map[string]float64 {

	var data = g.data
	var funcs = g.funcs
	var affinityFunc = g.affinityFunc

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

	calculatedIsalotatedFuncs := make(map[string]map[string]map[string]float64)
	for namefunc, datafunc := range f {
		funcdef := funcs[namefunc]
		valIsolatedFund := calculateisolatedfunc(namefunc, datafunc, funcdef)
		calculatedIsalotatedFuncs[namefunc] = valIsolatedFund
	}

	var ret = make(map[string]map[string]float64)
	for i := 0; i < len(nodenames); i++ {
		nodename := nodenames[i]
		n := map[string]float64{
			nodename: 0.0,
		}
		ret[nodename] = n
		for i := 0; i < len(nodenames); i++ {
			othernodename := nodenames[i]
			funcValues := make(map[string]float64)
			for funcName, _ := range funcs {
				funcValues[funcName] = calculatedIsalotatedFuncs[funcName][nodename][othernodename]
			}
			ret[nodename][othernodename] = affinityFunc(funcValues)
		}
	}

	return ret
}

func calculateisolatedfunc(namefunc string, datafunc []NodeAndData, funcdef graffinityfunc) map[string]map[string]float64 {
	var nodenames = make([]string, len(datafunc))
	var nodedata = make([]float64, len(datafunc))

	for _, data := range datafunc {
		nodenames = append(nodenames, data.name)
		nodedata = append(nodedata, data.data...)
	}

	ret := make(map[string]map[string]float64)
	for _, n1 := range datafunc {
		ret[n1.name] = make(map[string]float64)
		for _, n2 := range datafunc {
			values := append(n1.data, n2.data...)
			ret[n1.name][n2.name] = funcdef(values)
		}

	}
	return ret
}
