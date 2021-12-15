package main

import "fmt"
import "errors"

type Stack struct {
	underlying []int
}

func (s *Stack) push(val int) {
	s.underlying = append(s.underlying, val)
}

func (s *Stack) peek() (int, error) {
	if len(s.underlying) <= 0 {
		return -1, errors.New("peek called on empty Stack")
	}
	return s.underlying[len(s.underlying)-1], nil
}

func (s *Stack) pop() (int, error) {
	if len(s.underlying) <= 0 {
		return -1, errors.New("pop called on empty Stack")
	}
	res := s.underlying[len(s.underlying)-1]
	s.underlying = s.underlying[:len(s.underlying)-1]
	return res, nil
}

type MaxStack struct {
	underlying *Stack
	maxHistory *Stack
}

func (s *MaxStack) push(val int) {
	s.underlying.push(val)
	currentMax, err := s.maxHistory.peek()
	if err != nil || val > currentMax {
		s.maxHistory.push(val)
	}
}

func (s *MaxStack) peek() (int, error) {
	return s.underlying.peek()
}

func (s *MaxStack) pop() (int, error) {
	res, err := s.underlying.pop()
	if err != nil {
		return -1, errors.New("pop called on empty MaxStack")
	}
	currentMax, err := s.maxHistory.peek()
	if err != nil {
		return -1, errors.New("BAD! maxHistory empty while underlying nonempty!")
	}
	if res >= currentMax {
		s.maxHistory.pop()
	}
	return res, nil
}

func (s *MaxStack) max() (int, error) {
	_, err := s.underlying.peek()
	if err != nil {
		return -1, errors.New("max called on empty MaxStack")
	}
	return s.maxHistory.peek()
}

func main() {
	st := &MaxStack{new(Stack), new(Stack)}
	st.push(3)
	st.push(2)
	st.push(1)
	st.push(0)
	st.push(-1)
	st.push(4)

	for i := 0; i < 6; i++ {
		m, _ := st.max()
		p, _ := st.pop()
		fmt.Printf("POP: %v | MAX: %v\n", p, m)
	}

	_, err := st.pop()
	if err == nil {
		fmt.Println("Stack should have been empty & returned error on pop()!")
	}

	_, err = st.maxHistory.pop()
	if err == nil {
		fmt.Println("maxHistory wasn't empty!")
	}
}
