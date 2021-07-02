package alive

import "time"

var cache = make(map[string]Alive)

type Alive struct {
	LastReport time.Time
}

func IsAlive(id string) bool {
	a, ok := cache[id]
	if !ok {
		return false
	}
	return time.Now().UTC().Sub(a.LastReport) < time.Duration(60)*time.Second
}

func Refresh(id string) {
	cache[id] = Alive{LastReport: time.Now().UTC()}
}
