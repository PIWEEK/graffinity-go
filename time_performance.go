package main

import "fmt"
import "math"
import "time"

//import "runtime"

var genderchoices = []float32{1, 2}
var agechoices = Range(18, 70)
var languageslist = Range(1, 20)
var friendslist = Range(1, 10)
var knowngameslist = Range(1, 5000)
var toplaywishlistlist = knowngameslist
var favouritegameslist = knowngameslist
var preferenceslist = Range(1, 5)
var vetoeslist = Range(1, 5)
var gametypelist = Range(1, 64)
var skill01choices = Range(1, 100)
var skill02choices = skill01choices
var skill03choices = skill01choices
var skill04choices = skill01choices
var skill05choices = skill01choices
var skill06choices = skill01choices
var skill07choices = skill01choices
var skill08choices = skill01choices
var skill09choices = skill01choices
var skill10choices = skill01choices
var guildslist = Range(1, 200)

func datagenerator(n int) map[string]map[string][]float32 {
	data := map[string]map[string][]float32{}

	for i := 1; i < n; i++ {
		name := fmt.Sprintf("n%d", i)
		data[name] = map[string][]float32{
			"gender":         []float32{Choice(genderchoices)},
			"age":            []float32{Choice(agechoices)},
			"languages":      Sample(languageslist, 3),
			"friends":        Sample(friendslist, RandInt(1, 5)),
			"knowngames":     Sample(knowngameslist, RandInt(1, 20)),
			"toplaywishlist": Sample(toplaywishlistlist, RandInt(1, 20)),
			"favouritegames": Sample(favouritegameslist, RandInt(1, 20)),
			"preferences":    Sample(preferenceslist, RandInt(1, 3)),
			"vetoes":         Sample(vetoeslist, RandInt(1, 3)),
			"gametype":       Sample(gametypelist, RandInt(1, 5)),
			"skill01":        []float32{Choice(skill01choices)},
			"skill02":        []float32{Choice(skill02choices)},
			"skill03":        []float32{Choice(skill03choices)},
			"skill04":        []float32{Choice(skill04choices)},
			"skill05":        []float32{Choice(skill05choices)},
			"skill06":        []float32{Choice(skill06choices)},
			"skill07":        []float32{Choice(skill07choices)},
			"skill08":        []float32{Choice(skill08choices)},
			"skill09":        []float32{Choice(skill09choices)},
			"skill10":        []float32{Choice(skill10choices)},
			"guilds":         Sample(guildslist, RandInt(1, 5)),
		}
	}

	return data
}

func main() {
	fmt.Println("Creating sample data")
	data := datagenerator(100001)

	//startTime := time.Now()

	genderFunc := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	ageFunc := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	languagesFunc := func(x []float32) float32 { return 5. * float32(len(x)-len(RemoveDuplicates(x))) }
	friendsFunc := func(x []float32) float32 { return 2.5 * float32(len(x)-len(RemoveDuplicates(x))) }
	knowngamesFunc := func(x []float32) float32 { return 2.5 * float32((len(x)-len(RemoveDuplicates(x)))/len(x)) }
	toplaywishlistFunc := func(x []float32) float32 { return 10. * float32((len(x)-len(RemoveDuplicates(x)))/len(x)) }
	favouritegamesFunc := func(x []float32) float32 { return 20. * float32((len(x)-len(RemoveDuplicates(x)))/len(x)) }
	preferencesFunc := func(x []float32) float32 { return 1. * float32(len(x)-len(RemoveDuplicates(x))) }
	vetoesFunc := func(x []float32) float32 { return 1. * float32(len(x)-len(RemoveDuplicates(x))) }
	gametypeFunc := func(x []float32) float32 { return 15. * float32((len(x)-len(RemoveDuplicates(x)))/len(x)) }
	skill01Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill02Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill03Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill04Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill05Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill06Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill07Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill08Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill09Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	skill10Func := func(x []float32) float32 { return float32(math.Abs(float64(Mean(x)-Stdev(x))) / float64(Mean(x))) }
	guildsFunc := func(x []float32) float32 { return 5. * float32((len(x)-len(RemoveDuplicates(x)))/len(x)) }

	funcs := map[string]func([]float32) float32{
		"gender":         genderFunc,
		"age":            ageFunc,
		"languages":      languagesFunc,
		"friends":        friendsFunc,
		"knowngames":     knowngamesFunc,
		"toplaywishlist": toplaywishlistFunc,
		"favouritegames": favouritegamesFunc,
		"preferences":    preferencesFunc,
		"vetoes":         vetoesFunc,
		"gametype":       gametypeFunc,
		"skill01":        skill01Func,
		"skill02":        skill02Func,
		"skill03":        skill03Func,
		"skill04":        skill04Func,
		"skill05":        skill05Func,
		"skill06":        skill06Func,
		"skill07":        skill07Func,
		"skill08":        skill08Func,
		"skill09":        skill09Func,
		"skill10":        skill10Func,
		"guilds":         guildsFunc,
	}

	affinityFunc := func(dictValues map[string]float32) float32 {
		return dictValues["gender"] + dictValues["age"] + dictValues["languages"] +
			dictValues["friends"] + dictValues["knowngames"] + dictValues["toplaywishlist"] +
			dictValues["favouritegames"] + dictValues["preferences"] + dictValues["vetoes"] +
			dictValues["gametype"] + dictValues["skill01"] + dictValues["skill02"] +
			dictValues["skill03"] + dictValues["skill04"] + dictValues["skill05"] +
			dictValues["skill06"] + dictValues["skill07"] + dictValues["skill08"] +
			dictValues["skill09"] + dictValues["skill10"] + dictValues["guilds"]
	}

	groupaffinityFunc := func(dictValues map[string]float32) float32 {
		return dictValues["gender"] + dictValues["age"] + dictValues["languages"] +
			dictValues["friends"] + dictValues["knowngames"] + dictValues["toplaywishlist"] +
			dictValues["favouritegames"] + dictValues["preferences"] + dictValues["vetoes"] +
			dictValues["gametype"] + dictValues["skill01"] + dictValues["skill02"] +
			dictValues["skill03"] + dictValues["skill04"] + dictValues["skill05"] +
			dictValues["skill06"] + dictValues["skill07"] + dictValues["skill08"] +
			dictValues["skill09"] + dictValues["skill10"] + dictValues["guilds"]
	}

	/*
		g := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
		results := g.calculate()
		fmt.Println("[n1][n18]", results["n1"]["n18"], "[n18][n1]", results["n18"]["n1"])

		memstats := new(runtime.MemStats)
		runtime.ReadMemStats(memstats)
		fmt.Println("memstats before GC: bytes =", memstats.HeapAlloc, "footprint =",
			memstats.Sys)

	*/
	endTime := time.Now()
	//fmt.Println("ElapsedTime in seconds:", endTime.Sub(startTime))

	g2 := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
	startTime2 := time.Now()
	fmt.Println("ElapsedTime in seconds:", startTime2.Sub(endTime))
	fmt.Println("---------------> Launching for node n1")
	r := g2.calculatefornode("n1")
	fmt.Println("n1 n1000", r.get("n1", "n1000"), r.get("n1000", "n1"))

	endTime2 := time.Now()
	fmt.Println("ElapsedTime in seconds:", endTime2.Sub(startTime2))

	/*
		g3 := Graffinity{data: data, funcs: funcs, affinityFunc: affinityFunc, groupaffinityFunc: groupaffinityFunc}
		startTime3 := time.Now()
		fmt.Println("ElapsedTime in seconds:", endTime2.Sub(endTime))
		fmt.Println("----------------> Launching for group n1 n2 n3")
		g3.calculateforgroup([]string{"n1", "n2", "n3"})

		endTime3 := time.Now()
		fmt.Println("ElapsedTime in seconds:", endTime3.Sub(startTime3))
	*/
}
