package listeners

import (
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/rdb"
	"mmartinjoo/trees/store"
)

func LogWriteCommand(data map[string]any) {
	command, _ := data["command"].(string)
	key, _ := data["key"].(string)
	args, _ := data["args"].([]int64)

	aol.Write(command, key, args)

	if store.ShouldPersist() {
		rdb.Persist()
	}
}