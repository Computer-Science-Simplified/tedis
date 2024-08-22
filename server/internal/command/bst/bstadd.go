package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"github.com/gookit/event"
)

type BSTAdd struct {
	Params *types.CommandParams
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

func (b BSTAdd) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}

func (b BSTAdd) GetParams() *types.CommandParams {
	return b.Params
}
