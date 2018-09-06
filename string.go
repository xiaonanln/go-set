package go_set

type StringSet map[string]struct{}

func (s StringSet) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSet) Remove(v string) {
	delete(s, v)
}

func (s StringSet) Contains(v string) bool {
	_, ok := s[v]
	return ok
}

func (s StringSet) Elems() []string {
	elems := make([]string, 0, len(s))
	for v := range s {
		elems = append(elems, v)
	}
	return elems
}
