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
	cmd.CommandParams = params
	cmd.AccessType = enum.READ

	return cmd
}

func (b *BTExists) doExecute() (string, error) {
	t, err := tree.Create(b.CommandParams.Key, b.CommandParams.Type)

	if err != nil {
		return "", err
	}

	exists := t.Exists(b.CommandParams.Args[0])

	return strconv.FormatBool(exists), nil
}
