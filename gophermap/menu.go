package gophermap

import "strings"

func RenderMenu(items ...*Item) string {
	n := len(items)
	arr := make([]string, n)

	for i := range n {
		arr[i] = items[i].String()
	}

	ans := strings.Join(arr, "\n")

	return ans
}
