package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTGetAll struct {
	types.BaseCommand
}

func NewBSTGetAll(params *types.CommandParams) *BSTGetAll {
	bst := &BSTGetAll{}

	bst.DoExecuteFunc = bst.doExecute
	bst.CommandParams = params
	bst.AccessType = enum.READ

	return bst
}

func (b *BSTGetAll) doExecute() (string, error) {
	t, err := tree.Create(b.CommandParams.Key, b.CommandParams.Type)

	if err != nil {
		return "", err
	}

	values := t.GetAll()

	return fmt.Sprintf("%v", values), nil
}
