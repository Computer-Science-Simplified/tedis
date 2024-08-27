package command

import "github.com/Computer-Science-Simplified/tedis/server/internal/types"

type Command interface {
	Execute(shouldFireEvent bool) (string, error)
	GetParams() *types.CommandParams
	String() string
}
