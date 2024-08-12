package command

import (
	"fmt"
	"mmartinjoo/trees/commands"
	"mmartinjoo/trees/factory"
	"strconv"
)

type Command struct {
	Name string
	Key  string
	Args []int64
	Type string
}

func (c *Command) Execute() string {
	tree := factory.Create(c.Key, c.Type)

	switch c.Name {
	case commands.BTADD:
		tree.Insert(c.Args[0], tree.Root, true)
		return "ok"

	case commands.BTEXISTS:
		exists := tree.Exists(c.Args[0])
		return strconv.FormatBool(exists)

	case commands.BTGETALL:
		values := tree.ToArray()
		return fmt.Sprintf("%v", values)

	case commands.BTREM:
		tree.Remove(c.Args[0])
		return "ok"
	default:
		return "ok"
	}
}
