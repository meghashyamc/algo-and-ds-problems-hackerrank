package graphs

import (
	"errors"
)

// The Ruler of HackerLand believes that every citizen of the country should have access to a library. Unfortunately, HackerLand was hit by a tornado that destroyed all of its libraries and obstructed its roads! As you are the greatest programmer of HackerLand, the ruler wants your help to repair the roads and build some new libraries efficiently.

// HackerLand has n cities numbered from 1 to n. The cities are connected by m bidirectional roads. A citizen has access to a library if:

// Their city contains a library.
// They can travel by road from their city to a city containing a library.
// The following figure is a sample map of HackerLand where the dotted lines denote obstructed roads:
// https://s3.amazonaws.com/hr-challenge-images/0/1481983010-b779ad2b2b-torque1.png

// The cost of repairing any road is c-road dollars, and the cost to build a library in any city is c-lib dollars. If in the above example c-road = 2 and c-lib=3, we would build 5 roads at a cost of 5 x 2 = 10 and 2 libraries for a cost of 6. We don't need to rebuild one of the roads in the cycle 1->2->3->1.

// You are given q queries, where each query consists of a map of HackerLand and value of c-lib and c-road. For each query, find the minimum cost of making libraries accessible to all the citizens and print it on a new line.

// Function Description

// Complete the function roadsAndLibraries in the editor below. It must return the minimal cost of providing libraries to all, as an integer.

// roadsAndLibraries has the following parameters:

// n: integer, the number of cities
// c_lib: integer, the cost to build a library
// c_road: integer, the cost to repair a road
// cities: 2D array of integers where each cities[i] contains two integers that represent cities connected by an obstructed road
//more info: https://www.hackerrank.com/challenges/torque-and-development

//array to store values in a union-find type structure
var cityRoots []int32
var size []int32

func roadsAndLibraries(n int32, clib int32, croad int32, cities [][]int32) int64 {

	if n <= 0 {
		return 0
	}
	if len(cities) == 0 {
		return int64(clib) * int64(n)
	}
	if len(cities[0]) != 2 {
		panic(errors.New("parameters received in incorrect format, cities[i] must have length 2"))
	}

	//final cost is a x clib + b x croad such that a+b = n
	//if clib <= croad, the (a+b)  x clib will be <= a x clib + b x croad
	if clib <= croad {
		return int64(clib) * int64(n)
	}
	cityRoots = make([]int32, n+1)
	initialize(cityRoots)
	size = make([]int32, n+1)
	var roadsToRebuild int32

	citiesList := make(map[int32]bool)
	for _, road := range cities {
		citiesList[road[0]] = true
		citiesList[road[1]] = true
		if !areConnected(road[0], road[1]) {
			connect(road[0], road[1])
			roadsToRebuild++
		}
	}

	var costOfRoads int64 = int64(roadsToRebuild) * int64(croad)

	var costOfLibs int64 = int64(n-roadsToRebuild) * int64(clib)
	return costOfRoads + costOfLibs
}

func areConnected(p int32, q int32) bool {

	return root(p) == root(q)
}

func root(i int32) int32 {

	for {
		if cityRoots[i] == i {
			return int32(i)
		}
		cityRoots[i] = cityRoots[cityRoots[i]]
		i = cityRoots[i]
	}
}

func connect(p int32, q int32) {

	rootp := root(p)
	rootq := root(q)
	if rootp == rootq {
		return
	}

	if size[rootp] < size[rootq] {
		cityRoots[rootp] = rootq
		size[rootq] += size[rootp]
		return
	}
	cityRoots[rootq] = rootp
	size[rootp] += size[rootq]

}

func initialize(arr []int32) {

	for i := 0; i < len(arr); i++ {
		arr[i] = int32(i)
	}
}
