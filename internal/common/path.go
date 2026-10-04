package common

import (
	"strings"
)

func SafePath(path string) string {
	st := []string{}

	splittedPath := strings.SplitSeq(path, "/")
	for el := range splittedPath {
		switch el {
		case ".":
			continue
		case "..":
			stLen := len(st)
			if stLen > 0 {
				st = st[:stLen-1]
			}
		default:
			st = append(st, el)
		}
	}

	return strings.Join(st, "/")
}
