package medium

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func calculate(s string) int {
	val := Stack[int]{}
	op := Stack[byte]{}
	i := 0

	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}

		if isDigit(s[i]) {
			num := 0
			for i < len(s) && isDigit(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			val.Push(num)
			i--
		} else {
			curOp := s[i]
			for op.Len() > 0 && priority(op.Peek()) >= priority(curOp) {
				process(&val, op.Pop())
			}
			op.Push(curOp)
		}
		i++
	}

	for op.Len() > 0 {
		process(&val, op.Pop())
	}

	if val.Len() != 1 {
		panic("invalid expression")
	}

	return val.Peek()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func priority(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		panic("invalid operator")
	}
}

func process(val *Stack[int], op byte) {
	if val.Len() < 2 {
		panic("invalid expression")
	}

	r := val.Pop()
	l := val.Pop()
	switch op {
	case '+':
		val.Push(l + r)
	case '-':
		val.Push(l - r)
	case '*':
		val.Push(l * r)
	case '/':
		if r == 0 {
			panic("division by zero")
		}
		val.Push(l / r)
	default:
		panic("invalid operator")
	}
}
