package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/rdb"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
)

func AppendToAol(cmd *command.Command) error {
	err := aol.Append(cmd.Name, cmd.Key, cmd.Args)

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
