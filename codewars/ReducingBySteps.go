/*
	From https://www.codewars.com/kata/56efab15740d301ab40002ee
*/

package kata

import (
	"math"
)

func Gcdi(x, y int) int {
	// your code
	itr := int(math.Abs(float64(Mini(x, y))))

	for itr > 1 && !(x%itr == 0 && y%itr == 0) {
		itr--
	}

	return itr
}
func Som(x, y int) int {
	// your code
	return x + y
}
func Maxi(x, y int) int {
	// your code
	return int(math.Max(float64(x), float64(y)))
}
func Mini(x, y int) int {
	// your code
	return int(math.Min(float64(x), float64(y)))
}
func Lcmu(x, y int) int {
	// your code
	itr := int(math.Abs(float64(Maxi(x, y))))

	for !(itr%x == 0 && itr%y == 0) {
		itr++
	}

	return itr
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	// your code
	operatedArray := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			operatedArray[i] = f(init, arr[i])
		} else {
			operatedArray[i] = f(operatedArray[i-1], arr[i])
		}
	}

	return operatedArray
}
