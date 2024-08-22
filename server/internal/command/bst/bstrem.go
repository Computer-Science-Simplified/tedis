package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
)

type BSTRem struct {
	Params *types.CommandParams
}

func (b BSTRem) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	t.Remove(b.Params.Args[0])

	return "ok", nil
}

func (b BSTRem) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}

func (b BSTRem) GetParams() *types.CommandParams {
	return b.Params
}
