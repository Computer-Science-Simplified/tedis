package types

import "fmt"

type BaseCommand struct {
	Params *CommandParams

	DoExecuteFunc ExecuteFunc
}

func (b BaseCommand) Execute(shouldReport bool) (string, error) {
	res, err := b.DoExecuteFunc(shouldReport)

	fmt.Println("template method")

	if err != nil {
		return "", err
	}

	return res, nil
}

func (b BaseCommand) GetParams() *CommandParams {
	return b.Params
}

func (b BaseCommand) String() string {
	return fmt.Sprintf("[%s] %s %s %v", b.Params.Type, b.Params.Name, b.Params.Key, b.Params.Args)
}
