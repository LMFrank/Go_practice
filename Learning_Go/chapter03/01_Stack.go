package stack

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.i+1] = k
	s.i++
}

func (s *Stack) Pop() (res int) {
	res = s.data[s.i]
	s.i--
	return
}
