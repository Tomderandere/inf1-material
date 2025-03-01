package aufgabe6

// import "slices"

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {
	result := []int{}
	for i, el := range list {
		isdouble := false
		for _, v := range list[i+1:] {
			if el == v {
				isdouble = true
			}
		}
		if isdouble {
			isresult := false
			for _, res := range result {
				if el == res {
					isresult = true
				}
			}
			if !isresult {
				result = append(result, el)
			}
		}
	}
	return result
}
