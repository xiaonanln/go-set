package go_set

type InterfaceSet map[interface{}]struct{}

func (s InterfaceSet) Add(v interface{}) {
	s[v] = struct{}{}
}

func (s InterfaceSet) Remove(v interface{}) {
	delete(s, v)
}

func (s InterfaceSet) Contains(v interface{}) bool {
	_, ok := s[v]
	return ok
}

func (s InterfaceSet) Elems() []interface{} {
	elems := make([]interface{}, 0, len(s))
	for v := range s {
		elems = append(elems, v)
	}
	return elems
}
