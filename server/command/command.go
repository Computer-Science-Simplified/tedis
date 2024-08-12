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
	case commands.BSTADD:
		tree.Add(c.Args[0], tree.Root, true)
		return "ok"

	case commands.BSTEXISTS:
		exists := tree.Exists(c.Args[0])
		return strconv.FormatBool(exists)

	case commands.BSTGETALL:
		values := tree.ToArray()
		return fmt.Sprintf("%v", values)

	case commands.BSTREM:
		tree.Remove(c.Args[0], true)
		return "ok"
	default:
		return "ok"
	}
}
