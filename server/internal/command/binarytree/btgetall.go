package binarytree

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BTGetAll struct {
	types.BaseCommand
}

func NewBTGetAll(params *types.CommandParams) *BTGetAll {
	bst := &BTGetAll{}

	bst.DoExecuteFunc = bst.doExecute
	bst.Params = params
	bst.AccessType = enum.READ

	return bst
}

func (b *BTGetAll) doExecute() (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	values := t.GetAll()

	return fmt.Sprintf("%v", values), nil
}
