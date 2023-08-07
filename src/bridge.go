package main

type Bridge struct {
	adjacencyList AdjacencyList
	time          int
	disc, low     []TupleNameValue
	parent        []TupleParent
	bridgeList    []AdjacencyList
}

func (b *Bridge) findBridge() {
	n := len(b.adjacencyList)
	b.low = make([]TupleNameValue, n)
	b.disc = make([]TupleNameValue, n)
	b.parent = make([]TupleParent, n)

	for i := 0; i < n; i++ {
		b.disc[i] = TupleNameValue{
			b.adjacencyList[i].name,
			-1,
		}
		b.low[i] = TupleNameValue{
			b.adjacencyList[i].name,
			-1,
		}
		b.parent[i] = TupleParent{
			b.adjacencyList[i].name,
			"",
		}
	}

	b.time = 0
	for i := 0; i < n; i++ {
		if b.disc[i].val == -1 {
			b.DFS(i)
		}
	}

}

func (b *Bridge) DFS(i int) {
	b.disc[i].val = b.time
	b.low[i].val = b.time
	b.time = b.time + 1

	for _, v := range b.adjacencyList.FindNeighboursOf(b.disc[i].nodeName) {
		idx := b.adjacencyList.GetIndex(v)
		if b.disc[idx].val == -1 {
			b.parent[idx].parentName = b.disc[i].nodeName
			b.DFS(idx)
			b.low[i].val = min(b.low[i].val, b.low[idx].val)

			if b.low[idx].val > b.disc[i].val {
				b.bridgeList = append(b.bridgeList, []Node{{b.disc[i].nodeName, []string{b.disc[idx].nodeName}}})
			}
		} else if v != b.parent[i].parentName {
			b.low[i].val = min(b.low[i].val, b.disc[idx].val)
		}

	}
}
