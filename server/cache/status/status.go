package status

import (
	"errors"

	"github.com/kzw200015/ServerStatus/types"
)

var cache = make(map[string]types.Status)

func GetCache(id string) (types.Status, error) {
	s, ok := cache[id]
	if !ok {
		return types.Status{}, errors.New("node not online")
	}
	return s, nil
}

func SetCache(id string, status types.Status) {
	cache[id] = status
}
