package main

type Scc struct {
	adjacencyList AdjacencyList
	time          int
	disc, low     []TupleNameValue
	stack         Stack
	sccList       []AdjacencyList
}

func (s *Scc) findScc() {
	n := len(s.adjacencyList)
	s.low = make([]TupleNameValue, n)
	s.disc = make([]TupleNameValue, n)

	for i := 0; i < n; i++ {
		s.disc[i] = TupleNameValue{
			s.adjacencyList[i].name,
			-1,
		}
		s.low[i] = TupleNameValue{
			s.adjacencyList[i].name,
			-1,
		}
	}

	s.time = 0
	for i := 0; i < n; i++ {
		if s.disc[i].val == -1 {
			s.DFS(i)
		}
	}
}

func (s *Scc) DFS(i int) {
	s.disc[i].val = s.time
	s.low[i].val = s.time
	s.time = s.time + 1
	s.stack.Push(s.disc[i].nodeName)

	for _, v := range s.adjacencyList.FindNeighboursOf(s.disc[i].nodeName) {
		idx := s.adjacencyList.GetIndex(v)
		if s.disc[idx].val == -1 {
			s.DFS(idx)
			s.low[i].val = min(s.low[i].val, s.low[idx].val)
		} else if s.stack.Contain(v) {
			s.low[i].val = min(s.low[i].val, s.disc[idx].val)
		}

		if s.low[i].val == s.disc[i].val {
			var tempList AdjacencyList
			count := 1
			if !s.stack.IsEmpty() {
				for s.stack.Peek() != s.disc[i].nodeName {
					nodeName := s.stack.Pop()
					idx := s.adjacencyList.GetIndex(nodeName)
					neighbours := s.findSccNeighbour(idx, s.low[idx].val)
					node := Node{nodeName, neighbours}
					tempList = append(tempList, node)
					count++
				}
				nodeName := s.stack.Pop()
				idx := s.adjacencyList.GetIndex(nodeName)
				neighbours := s.findSccNeighbour(idx, s.low[idx].val)
				node := Node{nodeName, neighbours}
				tempList = append(tempList, node)
				s.sccList = append(s.sccList, tempList)
			}
		}
	}
}

func (s *Scc) findSccNeighbour(nodeIndex int, lowVal int) []string {
	neighbours := s.adjacencyList[nodeIndex].neighbours
	var tempList []string
	for _, v := range neighbours {
		idx := s.adjacencyList.GetIndex(v)
		if s.low[idx].val == lowVal {
			tempList = append(tempList, v)
		}
	}
	return tempList
}
