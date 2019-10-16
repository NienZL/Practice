package model

type Item struct {
	Value string
	Next  *Item
}
type Stack struct {
	Top  *Item
	Size int
}

func (s *Stack) Len() int {
	return s.Size
}

func (s *Stack) Peek() string {
	if s.Len() > 0 {
		top := s.Top.Value
		return top
	} else {
		return ""
	}

}
func (s *Stack) Pop() string {
	if s.Len() > 0 {
		top := s.Peek()
		s.Top = s.Top.Next
		s.Size--
		return top
	} else {
		return ""
	}

}
func (s *Stack) Push(newInput string) {
	s.Top = &Item{
		Value: newInput,
		Next:  s.Top,
	}
	s.Size++
}

func (s *Stack) Remove() {
	if s.Len() > 0 {
		s.Top = s.Top.Next
		s.Size--
	}
}
