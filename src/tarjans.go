package main

type Tuple struct {
	nodeName string
	val      int
}

type Scc struct {
	adjacencyList AdjacencyList
	n, time       int
	disc, low     []Tuple
	stack         Stack
	sccList       []AdjacencyList
}

func (s *Scc) findScc() {
	s.n = len(s.adjacencyList)
	s.low = make([]Tuple, s.n)
	s.disc = make([]Tuple, s.n)

	for i := 0; i < s.n; i++ {
		s.disc[i] = Tuple{
			s.adjacencyList[i].name,
			-1,
		}
		s.low[i] = Tuple{
			s.adjacencyList[i].name,
			-1,
		}
	}

	s.time = 0
	for i := 0; i < s.n; i++ {
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
			for s.stack.Peek() != s.disc[i].nodeName {
				nodeName := s.stack.Pop()
				idx := s.adjacencyList.GetIndex(nodeName)
				allNeighbours := s.adjacencyList[idx].neighbours
				neighbours := s.findSccNeighbour(allNeighbours, s.low[idx].val)
				node := Node{nodeName, neighbours}
				tempList = append(tempList, node)
			}
			nodeName := s.stack.Pop()
			idx := s.adjacencyList.GetIndex(nodeName)
			allNeighbours := s.adjacencyList[idx].neighbours
			neighbours := s.findSccNeighbour(allNeighbours, s.low[idx].val)
			node := Node{nodeName, neighbours}
			tempList = append(tempList, node)
			s.sccList = append(s.sccList, tempList)
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

func (s *Scc) findSccNeighbour(neighbours []string, lowVal int) []string {
	var tempList []string
	for _, v := range neighbours {
		idx := s.adjacencyList.GetIndex(v)
		if s.low[idx].val == lowVal {
			tempList = append(tempList, v)
		}
	}
	return tempList
}
