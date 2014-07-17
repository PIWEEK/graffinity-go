// graffinity project main.go
package main

import "runtime"
import "fmt"
import "time"

type graffinityfunc func(x []float32) float32

type NodeAndData struct {
	name string
	data []float32
}

type Graffinity struct {
	data              map[string]map[string][]float32
	funcs             map[string]func([]float32) float32
	affinityFunc      func(map[string]float32) float32
	groupaffinityFunc string
}

func (g Graffinity) calculate() map[string]map[string]map[string]float32 {

	nodenames, funcs, affinityFunc, f := g.init()

	t1 := time.Now()

	var calculateFuncs = make(map[string]map[string]map[string]float32)
	for _, n1 := range nodenames {
		calculateFuncs[n1] = make(map[string]map[string]float32)
		for _, n2 := range nodenames {
			calculateFuncs[n1][n2] = make(map[string]float32)
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

func (g Graffinity) calculateforgroup(nodegroup []string) map[string]map[string]map[string]float32 {

	nodenames, funcs, affinityFunc, f := g.init()

	t1 := time.Now()

	mygroup := nodegroup

	var calculateFuncs = make(map[string]map[string]map[string]float32)
	for _, n1 := range mygroup {
		calculateFuncs[n1] = make(map[string]map[string]float32)
		for _, n2 := range mygroup {
			calculateFuncs[n1][n2] = make(map[string]float32)
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

func calculateisolatedfunc(namefunc string, datafunc []NodeAndData, funcdef graffinityfunc, calculatedIsalotatedFuncsRef *map[string]map[string]map[string]float32, ch chan int) {
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

func (g Graffinity) calculatefornode(nodename string) map[string]map[string]map[string]float32 {

	nodenames, funcs, affinityFunc, f := g.init()

	t1 := time.Now()

	mynode := nodename

	var calculateFuncs = make(map[string]map[string]map[string]float32)
	calculateFuncs[mynode] = make(map[string]map[string]float32)
	for _, n2 := range nodenames {
		calculateFuncs[nodename][n2] = make(map[string]float32)
		for namefunc, _ := range funcs {
			calculateFuncs[nodename][n2][namefunc] = 0.0
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
		go calculateisolatedfuncfornode(namefunc, mynode, datafunc, funcdef, &calculateFuncs, ch)
		fmt.Println("Launching", namefunc)
	}

	//Wait all the routines to finish
	for funcName, channel := range channels {
		r := <-channel
		fmt.Println("Finishing", funcName, r)
	}

	t3 := time.Now()
	fmt.Println("ElapsedTime in seconds:", t3.Sub(t2))

	for _, n2 := range nodenames {
		calculateFuncs[mynode][n2]["total"] = affinityFunc(calculateFuncs[mynode][n2])

	}

	t4 := time.Now()
	fmt.Println("ElapsedTime in seconds:", t4.Sub(t3))
	return calculateFuncs
}

func calculateisolatedfuncfornode(namefunc string, anode string, datafunc []NodeAndData, funcdef graffinityfunc, calculatedIsalotatedFuncsRef *map[string]map[string]map[string]float32, ch chan int) {
	calculateFuncs := *calculatedIsalotatedFuncsRef

	anodeanddata := NodeAndData{"n1", []float32{33}}
	for i := 0; i < len(datafunc); i++ {
		n1 := anodeanddata
		n2 := datafunc[i]
		val := funcdef(append(n1.data, n2.data...))
		calculateFuncs[anode][n2.name][namefunc] = val
	}
	ch <- 1
}

func (g Graffinity) init() ([]string, map[string]func([]float32) float32, func(map[string]float32) float32, map[string][]NodeAndData) {

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

	return nodenames, funcs, affinityFunc, f

}
