package listeners

import (
	"mmartinjoo/trees/internal/persistence/aol"
	"mmartinjoo/trees/internal/persistence/rdb"
	"mmartinjoo/trees/internal/store"
)

func AppenToAol(data map[string]any) {
	command, _ := data["command"].(string)
	key, _ := data["key"].(string)
	args, _ := data["args"].([]int64)

	aol.Write(command, key, args)

	if store.ShouldPersist() {
		rdb.Persist()
	}
}
