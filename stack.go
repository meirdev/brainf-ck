package main

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}

	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]

	return value, true
}
