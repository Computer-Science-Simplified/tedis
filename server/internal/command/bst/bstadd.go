package bst

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"github.com/gookit/event"
)

type BSTAdd struct {
	types.BaseCommand
}

func (b BSTAdd) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	t.Add(b.Params.Args[0])

	if shouldReport {
		event.MustFire(enum.WriteCommandExecuted, event.M{
			"command": b,
		})
	}

	return "ok", nil
}
