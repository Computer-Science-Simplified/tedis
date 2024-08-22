package bst

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/Computer-Science-Simplified/tedis/server/internal/types"
	"github.com/gookit/event"
)

type BSTAdd struct {
	types.BaseCommand
}

func NewBSTAdd(params *types.CommandParams) *BSTAdd {
	fmt.Println("1")
	fmt.Println(params)
	bst := &BSTAdd{}

	bst.DoExecuteFunc = bst.doExecute
	bst.Params = params

	fmt.Println("2")
	fmt.Println(bst.Params)

	return bst
}

func (b *BSTAdd) doExecute(shouldReport bool) (string, error) {
	fmt.Println("3")
	fmt.Println(b)
	fmt.Println("--------")
	t, err := tree.Create(b.Params.Key, b.Params.Type)

	if err != nil {
		return "", err
	}

	t.Add(b.Params.Args[0])

	if shouldReport {
		event.MustFire(enum.WriteCommandExecuted, event.M{
			"command": b,
		})
	}

	return "ok", nil
}
