package main

import (
	"fmt"
	"strconv"
	"time"
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
	startTime := time.Now().UnixNano()
	scc.findScc()
	bridge.findBridge()
	endTime := time.Now().UnixNano()
	elapsedTime := endTime - startTime

	fmt.Println("SCC List:")
	fmt.Println(scc.sccList)
	fmt.Println("Bridge List:")
	fmt.Println(bridge.bridgeList)
	fmt.Printf("Run time: %v \n", elapsedTime)

	for i, a := range scc.sccList {
		name := "scc" + strconv.Itoa(i+1)
		a.DrawGraph(name)
	}

	for i, b := range bridge.bridgeList {
		name := "bridge" + strconv.Itoa(i+1)
		b.DrawGraph(name)
	}
}
