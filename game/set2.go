package game

type Set2[T any] struct {
	isSet []bool
	items []T
}

func newSet2[T any]() *Set2[T] {
	return &Set2[T]{}
}

func (s *Set2[T]) add(v T) {
	for i, isSet := range s.isSet {
		if !isSet {
			s.isSet[i] = true
			s.items[i] = v
			return
		}
	}
	s.isSet = append(s.isSet, true)
	s.items = append(s.items, v)
}

func (s Set2[T]) remove(i int) {
	s.isSet[i] = false
}

func (s Set2[T]) len() int {
	res := 0
	for _, isSet := range s.isSet {
		if isSet {
			res++
		}
	}
	return res
}

func (s Set2[T]) iter() []bool {
	return s.isSet
}

func (s Set2[T]) ref(index int) *T {
	return &s.items[index]
}

func (s Set2[T]) empty() bool {
	for _, isSet := range s.isSet {
		if isSet {
			return false
		}
	}
	return true
}
