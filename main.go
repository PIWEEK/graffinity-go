// graffinity project main.go
package main

import "runtime"
import "fmt"
import "time"

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

func (g Graffinity) calculate() map[string]map[string]map[string]float64 {
	runtime.GOMAXPROCS(len(g.funcs))
	//runtime.GOMAXPROCS(1)

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

	t1 := time.Now()

	var calculateFuncs = make(map[string]map[string]map[string]float64)
	for _, n1 := range nodenames {
		calculateFuncs[n1] = make(map[string]map[string]float64)
		for _, n2 := range nodenames {
			calculateFuncs[n1][n2] = make(map[string]float64)
			for namefunc, _ := range funcs {
				calculateFuncs[n1][n2][namefunc] = 0.0
			}
		}
	}

	t2 := time.Now()
	fmt.Println("ElapsedTime in seconds:", t2.Sub(t1))

	var channels = make(map[string]chan int, len(funcs))

	//Create a different routine for each function
	for namefunc, datafunc := range f {
		ch := make(chan int)
		channels[namefunc] = ch
		funcdef := funcs[namefunc]
		go calculateisolatedfunc(namefunc, datafunc, funcdef, &calculateFuncs, ch)
		fmt.Println("Launching", namefunc)
	}

	//Wait all the routines to finish
	for funcName, channel := range channels {
		r := <-channel
		fmt.Println("Finishing", funcName, r)
	}

	t3 := time.Now()
	fmt.Println("ElapsedTime in seconds:", t3.Sub(t2))

	for _, n1 := range nodenames {
		for _, n2 := range nodenames {
			calculateFuncs[n1][n2]["total"] = affinityFunc(calculateFuncs[n1][n2])
		}
	}

	t4 := time.Now()
	fmt.Println("ElapsedTime in seconds:", t4.Sub(t3))
	return calculateFuncs
}

func calculateisolatedfunc(namefunc string, datafunc []NodeAndData, funcdef graffinityfunc, calculatedIsalotatedFuncsRef *map[string]map[string]map[string]float64, ch chan int) {
	calculateFuncs := *calculatedIsalotatedFuncsRef
	for i := 0; i < len(datafunc); i++ {
		for j := i; j < len(datafunc); j++ {
			n1 := datafunc[i]
			n2 := datafunc[j]
			val := funcdef(append(n1.data, n2.data...))
			calculateFuncs[n1.name][n2.name][namefunc] = val
			calculateFuncs[n2.name][n1.name][namefunc] = val
		}
	}
	ch <- 1
}
