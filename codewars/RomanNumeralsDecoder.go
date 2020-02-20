/*
  From https://www.codewars.com/kata/51b6249c4612257ac0000005
*/

package kata

var valueMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

type stackInterface interface {
	peek() int
	pop() int
	push(val int)
	isEmpty() bool
}

type stack []int

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func Decode(roman string) int {
	result := 0
	var myStack stack

	for i := 0; i < len(roman); i++ {
		currentChar := roman[i]
		currentValue := valueMap[currentChar]

		if !myStack.isEmpty() {
			if myStack.peek() < currentValue {
				oldValue := myStack.pop()
				newValue := currentValue - oldValue
				myStack.push(newValue)
			} else {
				myStack.push(currentValue)
			}
		} else {
			myStack.push(currentValue)
		}
	}

	for _, val := range myStack {
		result = result + val
	}

	return result
}
