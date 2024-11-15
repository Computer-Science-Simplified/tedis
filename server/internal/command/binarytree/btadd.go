package binarytree

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BTAdd struct {
	types.BaseCommand
}

func NewBTAdd(params *types.CommandParams) *BTAdd {
	cmd := &BTAdd{}

	cmd.DoExecuteFunc = cmd.doExecute
	cmd.CommandParams = params
	cmd.AccessType = enum.WRITE

	return cmd
}

func (b *BTAdd) doExecute() (string, error) {
	t, err := tree.Create(b.CommandParams.Key, b.CommandParams.Type)

	if err != nil {
		return "", err
	}

	t.Add(b.CommandParams.Args[0])

	return "ok", nil
}
