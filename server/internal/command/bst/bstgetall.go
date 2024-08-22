package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTGetAll struct {
	types.BaseCommand
}

func NewBSTGetAll(params types.CommandParams) *BSTGetAll {
	return &BSTGetAll{BaseCommand: types.BaseCommand{Params: params}}
}

func (b BSTGetAll) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	values := t.GetAll()

	return fmt.Sprintf("%v", values), nil
}
