package rpn

type stack []string

func (s *stack) push(val string) {
	*s = append(*s, val)
}

func (s *stack) pop() string {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *stack) top() string {
	return (*s)[len(*s)-1]
}
