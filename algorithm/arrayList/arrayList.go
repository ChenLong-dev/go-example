package arrayList

import "fmt"

type ArrayList struct {
	data []interface{}
	size int
}

func (l *ArrayList) Add(element interface{}) {
	l.data = append(l.data, element)
	l.size++
}

func (l *ArrayList) RemoveByIndex(index int) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}
	element := l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.size--
	return element
}

func (l *ArrayList) RemoveByValue(element interface{}) bool {
	for i, v := range l.data {
		if v == element {
			if l.RemoveByIndex(i) == nil {
				return false
			}
			return true
		}
	}
	return false
}

func (l *ArrayList) Size() int {
	return l.size
}

// Reverse 反转ArrayList
func (l *ArrayList) Reverse() {
	for i, n := 0, l.size; i < n/2; i++ {
		l.data[i], l.data[n-1-i] = l.data[n-1-i], l.data[i]
	}
}

// Print 打印ArrayList
func (l *ArrayList) Print() {
	fmt.Printf("-> Size: %d\n", l.size)
	for i, v := range l.data {
		fmt.Printf("[%d-%d] element: %v\n", l.size, i+1, v)
	}
}


