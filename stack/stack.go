package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

/**
	返回顶部元素
 */
func (stack Stack) Top() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

/**
	弹出顶部元素
 */
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	newStack := *stack
	value := newStack[newStack.Len()-1]
	*stack = newStack[:newStack.Len()-1]
	return value, nil
}
