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
	t.DrawGraph(graphName)

	fmt.Println()
	fmt.Println()
	var scc Tarjan
	scc.adjacencyList = t
	scc.findScc()
	fmt.Println(scc.low)
	fmt.Println(scc.sccList)
	fmt.Println(scc.bridgeList)

	for i, a := range scc.sccList {
		name := "scc" + strconv.Itoa(i+1)
		//fmt.Println(a)
		a.DrawGraph(name)
	}

	for i, b := range scc.bridgeList {
		name := "bridge" + strconv.Itoa(i+1)
		b.DrawGraph(name)
	}
}
