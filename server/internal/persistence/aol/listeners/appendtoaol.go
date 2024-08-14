package listeners

import (
	"mmartinjoo/tedis/internal/persistence/aol"
	"mmartinjoo/tedis/internal/persistence/rdb"
	"mmartinjoo/tedis/internal/store"
)

func AppendToAol(data map[string]any) {
	command, _ := data["command"].(string)
	key, _ := data["key"].(string)
	args, _ := data["args"].([]int64)

	aol.Write(command, key, args)

	if store.ShouldPersist() {
		rdb.Persist()
	}
}
