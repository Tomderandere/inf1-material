package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	shortest := 0
	longest := 0
	l1islongest := false

	result := []int{}
	if len(l1) >= len(l2) {
		shortest = len(l2)
		longest = len(l1)
		l1islongest = true
	}
	if len(l1) <= len(l2) {
		shortest = len(l1)
		longest = len(l2)

	}
	for i := 0; i < longest; i++ {
		if i < shortest {
			result = append(result, l1[i]*l2[i])
		}
		if shortest <= i {
			if l1islongest {
				result = append(result, l1[i])
			} else {
				result = append(result, l2[i])
			}
		}
	}

	return result
}
