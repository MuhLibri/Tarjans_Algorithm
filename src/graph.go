package main

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"os"
)

type Node struct {
	name       string
	neighbours []string
}

func (n *Node) AddEdge(to string) {
	n.neighbours = append(n.neighbours, to)
}

type AdjacencyList []Node

func (a AdjacencyList) FindNeighboursOf(nodeName string) []string {
	for _, val := range a {
		if val.name == nodeName {
			return val.neighbours
		}
	}
	return nil
}

func (a AdjacencyList) GetIndex(nodeName string) int {
	for i, v := range a {
		if v.name == nodeName {
			return i
		}
	}
	return -1
}

func (n *Node) Print() {
	fmt.Printf("Node %s:", n.name)
	for _, element := range n.neighbours {
		fmt.Printf(" %s ", element)
	}
	fmt.Println()
}

func (a *AdjacencyList) drawGraph(graphName string) {
	// Make and draw graph
	g := graph.New(graph.StringHash, graph.Directed())

	for _, node := range *a {
		_ = g.AddVertex(node.name)
	}

	for _, node := range *a {
		for _, edge := range node.neighbours {
			_ = g.AddEdge(node.name, edge)
		}
	}
	file, _ := os.Create("../graph/" + graphName)
	_ = draw.DOT(g, file)
}
