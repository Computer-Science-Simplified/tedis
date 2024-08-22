package listeners

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/persistence/aol"
)

func AppendToAol(cmd command.Command) error {
	err := aol.Append(cmd)

	if err != nil {
		// Add the command to a dead letter queue and retry it later
		return err
	}

	return nil
}
