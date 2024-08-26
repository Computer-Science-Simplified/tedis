package binarytree

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BTRem struct {
	types.BaseCommand
}

func NewBTRem(params *types.CommandParams) *BTRem {
	cmd := &BTRem{}

	cmd.DoExecuteFunc = cmd.doExecute
	cmd.Params = params
	cmd.AccessType = enum.WRITE

	return cmd
}

func (b *BTRem) doExecute() (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)
	if err != nil {
		return "", err
	}

	t.Remove(b.Params.Args[0])

	return "ok", nil
}
