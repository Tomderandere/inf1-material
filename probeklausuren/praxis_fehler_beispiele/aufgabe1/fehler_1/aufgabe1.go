package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	longest := 0
	longeststring := -1
	for i, el := range list {
		if len(el) >= 3 {
			if el[:3] == "abc" {
				if len(el) > longest {
					longest = len(el)
					longeststring = i
				}
			}
		}
	}
	if longeststring == -1 {
		return ""
	}
	return list[longeststring]
}
