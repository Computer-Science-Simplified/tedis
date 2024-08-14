package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/rdb"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func AppendToAol(data map[string]any) error {
	command, _ := data["command"].(string)
	key, _ := data["key"].(string)
	args, _ := data["args"].([]int64)

	err := aol.Write(command, key, args)

	if err != nil {
		// Add the command to a dead letter queue and retry it later
		return err
	}

	if store.ShouldPersist() {
		err := rdb.Persist()
		if err != nil {
			return err
		}
	}

	return nil
}
