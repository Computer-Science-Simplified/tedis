package command

import "github.com/Computer-Science-Simplified/tedis/server/internal/types"

type Command interface {
	Execute(shouldReport bool) (string, error)
	GetParams() *types.CommandParams
	String() string
}
