package command

import (
	"fmt"
	"mmartinjoo/trees/internal/commands"
	"mmartinjoo/trees/internal/tree"
	"strconv"
)

type Command struct {
	Name string
	Key  string
	Args []int64
	Type string
}

func (c *Command) Execute() (string, error) {
	t, err := tree.Create(c.Key, c.Type)

	if err != nil {
		return "", err
	}

	switch c.Name {
	case commands.BSTADD:
		t.Add(c.Args[0], true)
		return "ok", nil

	case commands.BSTEXISTS:
		exists := t.Exists(c.Args[0])
		return strconv.FormatBool(exists), nil

	case commands.BSTGETALL:
		values := t.GetAll()
		return fmt.Sprintf("%v", values), nil

	case commands.BSTREM:
		t.Remove(c.Args[0], true)
		return "ok", nil
	default:
		return "", fmt.Errorf("command not found: %s", c.Name)
	}
}
