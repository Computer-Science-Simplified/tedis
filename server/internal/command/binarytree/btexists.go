package binarytree

import (
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"strconv"
)

type BTExists struct {
	types.BaseCommand
}

func NewBTExists(params *types.CommandParams) *BTExists {
	cmd := &BTExists{}

	cmd.DoExecuteFunc = cmd.doExecute
	cmd.Params = params
	cmd.AccessType = enum.READ

	return cmd
}

func (b *BTExists) doExecute() (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	exists := t.Exists(b.Params.Args[0])

	return strconv.FormatBool(exists), nil
}
