package go_set

type IntSet map[int]struct{}

func (s IntSet) Add(v int) {
	s[v] = struct{}{}
}

func (s IntSet) Remove(v int) {
	delete(s, v)
}

func (s IntSet) Contains(v int) bool {
	_, ok := s[v]
	return ok
}

func (s IntSet) Elems() []int {
	elems := make([]int, 0, len(s))
	for v := range s {
		elems = append(elems, v)
	}
	return elems
}
