package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTGetAll struct {
	Params *types.CommandParams
}

func (b BSTGetAll) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	values := t.GetAll()

	return fmt.Sprintf("%v", values), nil
}

func (b BSTGetAll) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}

func (b BSTGetAll) GetParams() *types.CommandParams {
	return b.Params
}
