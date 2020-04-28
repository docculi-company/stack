package stack

//
type Stack struct {
	body []byte
}

//
func New() *Stack {
	return &Stack{}
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len((*s).body) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(bt byte) {
	(*s).body = append((*s).body, bt) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return 0
	}

	index := len((*s).body) - 1   // Get the index of the top most element.
	element := (*s).body[index]   // Index into the slice and obtain the element.
	(*s).body = (*s).body[:index] // Remove it from the stack by slicing it off.
	return element
}

// View the top element.
func (s *Stack) Peek() byte {
	if s.IsEmpty() {
		return 0
	}

	return (*s).body[len((*s).body)-1]
}
