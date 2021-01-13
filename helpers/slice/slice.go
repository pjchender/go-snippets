// Package slice provide some utility methods
package slice

// Contains will check whether searchElement contains the target
func Contains(searchElement []*int, target int) bool {
	for _, element := range searchElement {
		if *element == target {
			return true
		}
	}
	return false
}

// Filter will check whether searchElement fulfill the handler condition
func Filter(searchElement []*int, handler func(*int) bool) []*int {
	n := 0
	for _, element := range searchElement {
		if handler(element) {
			searchElement[n] = element
			n++
		}
	}
	return searchElement[:n]
}
