package types

import "fmt"

type BaseCommand struct {
	Params CommandParams
}

func (b BaseCommand) GetParams() *CommandParams {
	return &b.Params
}

func (b BaseCommand) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}
