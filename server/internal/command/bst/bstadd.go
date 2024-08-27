package bst

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTAdd struct {
	types.BaseCommand
}

func NewBSTAdd(params *types.CommandParams) *BSTAdd {
	bst := &BSTAdd{}

	bst.DoExecuteFunc = bst.doExecute
	bst.CommandParams = params
	bst.AccessType = enum.WRITE

	return bst
}

func (b *BSTAdd) doExecute() (string, error) {
	t, err := tree.Create(b.CommandParams.Key, b.CommandParams.Type)

	if err != nil {
		return "", err
	}

	t.Add(b.CommandParams.Args[0])

	return "ok", nil
}
