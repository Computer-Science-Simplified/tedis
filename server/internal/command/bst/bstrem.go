package bst

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"github.com/gookit/event"
)

type BSTRem struct {
	types.BaseCommand
}

func NewBSTRem(params *types.CommandParams) *BSTRem {
	bst := &BSTRem{}

	bst.DoExecuteFunc = bst.doExecute
	bst.Params = params

	return bst
}

func (b *BSTRem) doExecute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	t.Remove(b.Params.Args[0])

	if shouldReport {
		event.MustFire(enum.WriteCommandExecuted, event.M{
			"command": b,
		})
	}

	return "ok", nil
}
