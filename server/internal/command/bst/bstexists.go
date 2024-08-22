package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"strconv"
)

type BSTExists struct {
	Params *types.CommandParams
}

func (b BSTExists) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	exists := t.Exists(b.Params.Args[0])

	return strconv.FormatBool(exists), nil
}

func (b BSTExists) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}

func (b BSTExists) GetParams() *types.CommandParams {
	return b.Params
}
