/*
	From https://www.codewars.com/kata/564057bc348c7200bd0000ff/go
*/

package kata

func Thirt(n int) int {
	// your code

	currentSum := CalculateRemainders(n)
	prevSum := 0

	for currentSum != prevSum {
		prevSum = currentSum
		currentSum = CalculateRemainders(prevSum)
	}
	return currentSum
}

func CalculateRemainders(n int) int {
	generalPattern := []int{1, 10, 9, 12, 3, 4}
	itr := 0
	sum := 0

	for n > 0 {
		multiplier := (n % 10)
		sum = sum + (generalPattern[itr%len(generalPattern)] * multiplier)
		n = n / 10
		itr++
	}

	return sum
}
