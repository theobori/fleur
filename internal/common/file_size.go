package common

import "fmt"

func GetStringFromSize(size int64) string {
	if size == 1 {
		return "1 byte"
	}

	sizeUnits := []string{"bytes", "KB", "MB", "GB", "TB"}
	i := 0
	n := len(sizeUnits)

	for {
		nextSize := size / 1000
		if nextSize <= 0 || i >= n {
			break
		}

		size = nextSize
		i += 1
	}

	ans := fmt.Sprintf("%d %s", size, sizeUnits[i])

	return ans
}
