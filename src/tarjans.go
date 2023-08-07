package main

type Tuple struct {
	nodeName string
	val      int
}

type Tarjan struct {
	adjacencyList AdjacencyList
	n, time       int
	disc, low     []Tuple
	stack         Stack
	sccList       []AdjacencyList
	bridgeList    []AdjacencyList
}

func (t *Tarjan) findScc() {
	t.n = len(t.adjacencyList)
	t.low = make([]Tuple, t.n)
	t.disc = make([]Tuple, t.n)

	for i := 0; i < t.n; i++ {
		t.disc[i] = Tuple{
			t.adjacencyList[i].name,
			-1,
		}
		t.low[i] = Tuple{
			t.adjacencyList[i].name,
			-1,
		}
	}

	t.time = 0
	for i := 0; i < t.n; i++ {
		if t.disc[i].val == -1 {
			t.DFS(i)
		}
	}
}

func (t *Tarjan) DFS(i int) {
	t.disc[i].val = t.time
	t.low[i].val = t.time
	t.time = t.time + 1
	t.stack.Push(t.disc[i].nodeName)

	for _, v := range t.adjacencyList.FindNeighboursOf(t.disc[i].nodeName) {
		idx := t.adjacencyList.GetIndex(v)
		if t.disc[idx].val == -1 {
			t.DFS(idx)
			t.low[i].val = min(t.low[i].val, t.low[idx].val)
		} else if t.stack.Contain(v) {
			t.low[i].val = min(t.low[i].val, t.disc[idx].val)
		}

		if t.low[i].val == t.disc[i].val {
			var tempList AdjacencyList
			count := 1
			for t.stack.Peek() != t.disc[i].nodeName {
				nodeName := t.stack.Pop()
				idx := t.adjacencyList.GetIndex(nodeName)
				neighbours := t.findSccNeighbour(idx, t.low[idx].val)
				node := Node{nodeName, neighbours}
				tempList = append(tempList, node)
				count++
			}
			nodeName := t.stack.Pop()
			idx := t.adjacencyList.GetIndex(nodeName)
			neighbours := t.findSccNeighbour(idx, t.low[idx].val)
			node := Node{nodeName, neighbours}
			tempList = append(tempList, node)
			t.sccList = append(t.sccList, tempList)
			if count == 2 {
				t.bridgeList = append(t.bridgeList, tempList)
			}
		}
	}
}

func min(x int, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

func (t *Tarjan) findSccNeighbour(nodeIndex int, lowVal int) []string {
	neighbours := t.adjacencyList[nodeIndex].neighbours
	var tempList []string
	for _, v := range neighbours {
		idx := t.adjacencyList.GetIndex(v)
		if t.low[idx].val == lowVal {
			tempList = append(tempList, v)
		} else {
			t.bridgeList = append(t.bridgeList, []Node{{name: t.adjacencyList[nodeIndex].name, neighbours: []string{v}}})
		}
	}
	return tempList
}
