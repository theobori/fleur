package gophermap

import "strings"

func RenderMenu(items ...*Item) string {
	ans := []string{}

	for _, item := range items {
		ans = append(ans, item.String())
	}

	return strings.Join(ans, "\n")
}
