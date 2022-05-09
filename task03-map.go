package homework
import "sort"
func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	myslice := make([]string, 0, len(input))

	for _, k := range keys {
		myslice = append(myslice, input[k])
	}

	return myslice
}
