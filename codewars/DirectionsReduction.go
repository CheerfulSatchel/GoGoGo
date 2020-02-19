/*
	From https://www.codewars.com/kata/550f22f4d758534c1100025a
*/

package kata

import (
	"strings"
)

func DirReduc(arr []string) []string {
	// your code
	itr := 0
	reducedDirs := arr[:]
	for itr < len(reducedDirs)-1 && len(reducedDirs) >= 2 {
		if (strings.Compare(reducedDirs[itr], "NORTH") == 0 && strings.Compare(reducedDirs[itr+1], "SOUTH") == 0) ||
			(strings.Compare(reducedDirs[itr], "SOUTH") == 0 && strings.Compare(reducedDirs[itr+1], "NORTH") == 0) ||
			(strings.Compare(reducedDirs[itr], "EAST") == 0 && strings.Compare(reducedDirs[itr+1], "WEST") == 0) ||
			(strings.Compare(reducedDirs[itr], "WEST") == 0 && strings.Compare(reducedDirs[itr+1], "EAST") == 0) {
			reducedDirs = append(reducedDirs[:itr], reducedDirs[itr+2:]...)
			itr = 0
		} else {
			itr++
		}
	}

	return reducedDirs
}
