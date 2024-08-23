package types

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/gookit/event"
)

type BaseCommand struct {
	Params        *CommandParams
	AccessType    string
	DoExecuteFunc ExecuteFunc
}

func (b BaseCommand) Execute(shouldFireEvent bool) (string, error) {
	res, err := b.DoExecuteFunc()

	if err != nil {
		return "", err
	}

	if b.AccessType == enum.WRITE && shouldFireEvent {
		event.MustFire(enum.WriteCommandExecuted, event.M{
			"command": b,
		})
	}

	event.MustFire(enum.CommandExecuted, event.M{
		"command": b,
	})

	return res, nil
}

func (b BaseCommand) GetParams() *CommandParams {
	return b.Params
}

func (b BaseCommand) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}
