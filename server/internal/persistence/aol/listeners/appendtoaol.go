package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/rdb"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
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
