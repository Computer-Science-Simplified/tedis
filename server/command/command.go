package command

import (
	"mmartinjoo/trees/factory"
	"strconv"
)

type Command struct {
	Name string
	Key string
	Args []int64
	Type string
}

func (c *Command) Execute() string {
	tree := factory.Create(c.Key, c.Type)

	if c.Name == "BTADD" {
		tree.Insert(c.Args[0], tree.Root)

		return "ok"
	}

	if c.Name == "BTEXISTS" {
		exists := tree.Exists(c.Args[0])

		return strconv.FormatBool(exists)
	}

	return "ok"
}