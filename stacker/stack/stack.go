package stack

import (
	"errors"
)

//Stack is a slice of any type.
type Stack []interface{}

func (stack Stack) len() int {
	return len(stack)
}

func (stack Stack) cap() int {
	return cap(stack)
}

/*IsEmpty method takes a stack,
* and returns either true if the stack is empty,
* and false if the stack isn't empty.
* @param: Stack
 */
func (stack Stack) IsEmpty() bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

/*Push method takes a pointer to a Stack value,
* and adds an item into the stack.
* @param: Stack
 */
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

/*Top method takes a Stack,
* and returns the item at the top of stack if it exists.
* @param: Stack
* @output: (interface{}, error)
 */
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

/*Pop method takes a pointer to a Stack value,
* and deletes the last item in the stack if it exists.
* @param: *Stack
* @output: (interface{}, error)
 */
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}
