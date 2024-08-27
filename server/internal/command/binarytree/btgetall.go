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
	cmd := &BTGetAll{}

	cmd.DoExecuteFunc = cmd.doExecute
	cmd.CommandParams = params
	cmd.AccessType = enum.READ

	return cmd
}

func (b *BTGetAll) doExecute() (string, error) {
	t, err := tree.Create(b.CommandParams.Key, b.CommandParams.Type)

	if err != nil {
		return "", err
	}

	values := t.GetAll()

	return fmt.Sprintf("%v", values), nil
}
