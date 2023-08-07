package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Masukkan nama file: ")
	var fileName string
	fmt.Scanln(&fileName)
	t := readFile("../tests/" + fileName)

	fmt.Println("\n\nBerikut adalah adjacency list dari input: ")
	for _, elm := range t {
		elm.Print()
	}
	graphName := "original-graph"
	t.DrawGraph(graphName)

	fmt.Println()
	fmt.Println()

	var scc Tarjan
	scc.adjacencyList = t
	scc.findScc()
	scc.findBridge()

	fmt.Println("SSC List:")
	fmt.Println(scc.sccList)
	fmt.Println("Bridge List:")
	fmt.Println(scc.bridgeList)
	runTime := 2.4
	fmt.Printf("Run time: %f \n", runTime)

	for i, a := range scc.sccList {
		name := "scc" + strconv.Itoa(i+1)
		a.DrawGraph(name)
	}

	for i, b := range scc.bridgeList {
		name := "bridge" + strconv.Itoa(i+1)
		b.DrawGraph(name)
	}
}
