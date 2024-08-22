package bst

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"strconv"
)

type BSTExists struct {
	types.BaseCommand
}

func NewBSTExists(params *types.CommandParams) *BSTExists {
	bst := &BSTExists{}

	bst.DoExecuteFunc = bst.doExecute
	bst.Params = params

	return bst
}

func (b *BSTExists) doExecute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	exists := t.Exists(b.Params.Args[0])

	return strconv.FormatBool(exists), nil
}
