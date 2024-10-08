package stack

import "fmt"

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0)}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item := s.items[s.Size()-1] // 获取栈顶元素
	s.items = s.items[:s.Size()-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[s.Size()-1]
}

func (s *Stack) Print() {
	fmt.Printf("栈顶<-栈底 [%d] \n", s.Size())
	for i:=s.Size()-1; i>0; i-- {
		fmt.Printf("[%d]:%v\n", i, s.items[i])
	}
}
