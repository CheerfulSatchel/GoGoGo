/*
	From https://www.codewars.com/kata/5541f58a944b85ce6d00006a
*/

package kata

func ProductFib(prod uint64) [3]uint64 {
	// your code
	finalFib := [3]uint64{0, 0, 0}

	var prevValue uint64 = 0
	var currentValue uint64 = 1
	currentProduct := uint64(prevValue * currentValue)

	for currentProduct < prod {
		oldCurrentValue := currentValue
		currentValue = prevValue + currentValue
		prevValue = oldCurrentValue

		currentProduct = uint64(currentValue * prevValue)
	}

	finalFib[0] = prevValue
	finalFib[1] = currentValue

	if currentProduct == prod {
		finalFib[2] = 1
	} else {
		finalFib[2] = 0
	}

	return finalFib
}
