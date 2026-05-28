package evaluator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/theobori/fleur/gophermap"
)

func ExtendRandomNumberItem(_arguments ...any) (string, error) {
	t := time.Now().UnixNano()
	r := rand.New(rand.NewSource(t))

	number := r.Uint32()

	item := gophermap.Item{
		ItemType:    gophermap.ItemTypeInlineText,
		Description: fmt.Sprintf("%d", number),
		Selector:    "",
		Domain:      "",
		Port:        1,
	}

	ans := item.String()

	return ans, nil
}
