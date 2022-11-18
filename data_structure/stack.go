package data_structure

type node[T any] struct {
	next *node[T]
	val  T
}

type Stack[T any] struct {
	top *node[T]
	len int
}

func (s *Stack[T]) Len() int {
	return s.len
}

func (s *Stack[T]) String() string {
	cur := s.top
	res := ""
	for cur != nil {
		cur = cur.next
	}
	return res
}

func (s *Stack[T]) Peek() T {
	if s.top == nil {
		panic(any("end of stack"))
	}
	return s.top.val
}

func (s *Stack[T]) Pop() T {
	if s.top == nil {
		panic(any("end of stack"))
	}
	val := s.top.val
	s.top = s.top.next
	s.len -= 1
	return val
}

func (s *Stack[T]) Push(value T) {
	n := &node[T]{
		val:  value,
		next: s.top,
	}
	s.len += 1
	s.top = n
}
