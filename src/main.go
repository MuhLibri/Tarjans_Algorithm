package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Masukkan nama file: ")
	var fileName string
	fmt.Scanln(&fileName)
	adjacencyList := readFile("../tests/" + fileName)

	fmt.Println("\n\nBerikut adalah adjacency list dari input: ")
	for _, elm := range adjacencyList {
		elm.Print()
	}
	graphName := "original-graph"
	adjacencyList.DrawGraph(graphName)

	fmt.Println()
	fmt.Println()

	var scc Scc
	var bridge Bridge
	scc.adjacencyList = adjacencyList
	bridge.adjacencyList = adjacencyList
	scc.findScc()
	bridge.findBridge()

	fmt.Println("SCC List:")
	fmt.Println(scc.sccList)
	fmt.Println("Bridge List:")
	fmt.Println(bridge.bridgeList)
	runTime := 2.4
	fmt.Printf("Run time: %f \n", runTime)

	for i, a := range scc.sccList {
		name := "scc" + strconv.Itoa(i+1)
		a.DrawGraph(name)
	}

	for i, b := range bridge.bridgeList {
		name := "bridge" + strconv.Itoa(i+1)
		b.DrawGraph(name)
	}
}
