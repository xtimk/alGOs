package algos

func Zfunc(text string) []int {
	currentZIndexIntent := 1
	boxLeftBound := 0
	boxRightBound := 0

	zarray := make([]int, len(text))

	zarray[0] = 0

	for currentZIndexIntent < len(text) {

		if currentZIndexIntent > boxRightBound {
			// calc z of currentZIndexIntent following the naive mode
			zarray[currentZIndexIntent] = aux_findZ(text, currentZIndexIntent, 0)

			// eventually update the box
			candidateBoxLeft := zarray[currentZIndexIntent]
			candidateBoxRight := candidateBoxLeft + zarray[candidateBoxLeft] - 1

			if candidateBoxRight > boxRightBound {
				boxLeftBound = candidateBoxLeft
				boxRightBound = candidateBoxRight
			}

		} else {
			if currentZIndexIntent+zarray[currentZIndexIntent-boxLeftBound]-1 < boxRightBound {
				zarray[currentZIndexIntent] = zarray[currentZIndexIntent-boxLeftBound]

				// no need to check to update the box.

			} else {
				zarray[currentZIndexIntent] = aux_findZ(text, currentZIndexIntent+boxRightBound, (boxRightBound - currentZIndexIntent))

				// update the box (this will happen for sure..)
				candidateBoxLeft := zarray[currentZIndexIntent]
				candidateBoxRight := candidateBoxLeft + zarray[candidateBoxLeft] - 1

				if candidateBoxRight > boxRightBound {
					boxLeftBound = candidateBoxLeft
					boxRightBound = candidateBoxRight
				}

			}
		}
		currentZIndexIntent++
	}
	return zarray
}

// params
// text: the string for which we are calc z
// Zi: the index
// startingPoint: where to start cfr: the classic value it's 0 (the beginning of the text) or can be customized with another value
func aux_findZ(text string, Zi int, startingPoint int) int {
	pointerFromStart := startingPoint
	pointerFromZIndex := Zi

	for pointerFromZIndex < len(text) && pointerFromStart < len(text) && text[pointerFromStart] == text[pointerFromZIndex] {
		pointerFromStart++
		pointerFromZIndex++
	}

	return pointerFromStart
}
