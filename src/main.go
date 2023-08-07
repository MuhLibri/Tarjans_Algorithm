package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Masukkan nama file: ")
	fileName := "test1.txt"
	//fmt.Scanln(&fileName)
	t := readFile("../tests/" + fileName)
	fmt.Println()
	fmt.Println("Berikut adalah adjacency list dari input: ")
	for _, elm := range t {
		elm.Print()
	}
	graphName := "original-graph"
	t.drawGraph(graphName)

	fmt.Println()
	fmt.Println()
	var scc Scc
	scc.adjacencyList = t
	scc.findScc()
	fmt.Println(scc.sccList)
	fmt.Println(scc.low)

	for i, a := range scc.sccList {
		name := "scc" + strconv.Itoa(i+1)
		a.drawGraph(name)
	}
}
