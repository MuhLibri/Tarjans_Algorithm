package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fileName string) AdjacencyList {
	var adjacencyList AdjacencyList
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		found := false
		for i := 0; i < len(adjacencyList); i++ {
			if adjacencyList[i].name == string(scanner.Text()[0]) {
				adjacencyList[i].AddEdge(string(scanner.Text()[2]))
				found = true
			}
		}
		if !found {
			adjacencyList = append(adjacencyList, Node{name: string(scanner.Text()[0]), neighbours: []string{string(scanner.Text()[2])}})
		}
	}
	return adjacencyList
}

func min(x int, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

type TupleNameValue struct {
	nodeName string
	val      int
}

type TupleParent struct {
	nodeName   string
	parentName string
}
