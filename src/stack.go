package main

import "fmt"

type Stack struct {
	elm []string
}

func (s *Stack) Push(v string) {
	s.elm = append(s.elm, v)
}

func (s *Stack) Pop() string {
	if len(s.elm) == 0 {
		fmt.Println("Stack is empty")
		return ""
	} else {
		l := len(s.elm)
		retVal := s.elm[l-1]
		s.elm = s.elm[:l-1]
		return retVal
	}
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return "Stack is Empty"
	} else {
		return s.elm[len(s.elm)-1]
	}
}

func (s *Stack) Contain(str string) bool {
	for _, val := range s.elm {
		if val == str {
			return true
		}
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	if len(s.elm) == 0 {
		return true
	} else {
		return false
	}
}
