package bst

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTRem struct {
	types.BaseCommand
}

func NewBSTRem(params *types.CommandParams) *BSTRem {
	bst := &BSTRem{}

	bst.DoExecuteFunc = bst.doExecute
	bst.Params = params
	bst.AccessType = enum.WRITE

	return bst
}

func (b *BSTRem) doExecute() (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	t.Remove(b.Params.Args[0])

	return "ok", nil
}
