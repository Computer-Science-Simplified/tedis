package command

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/tree"
	"github.com/gookit/event"
	"strconv"
)

type Command struct {
	Name string
	Key  string
	Args []int64
	Type string
}

func (c *Command) Execute(shouldReport bool) (string, error) {
	t, err := tree.Create(c.Key, c.Type)

	if err != nil {
		return "", err
	}

	switch c.Name {
	case enum.BSTADD:
		t.Add(c.Args[0])

		if shouldReport {
			event.MustFire(enum.WriteCommandExecuted, event.M{
				"command": c,
			})
		}

		return "ok", nil

	case enum.BSTEXISTS:
		exists := t.Exists(c.Args[0])
		return strconv.FormatBool(exists), nil

	case enum.BSTGETALL:
		values := t.GetAll()
		return fmt.Sprintf("%v", values), nil

	case enum.BSTREM:
		t.Remove(c.Args[0])

		if shouldReport {
			event.MustFire(enum.WriteCommandExecuted, event.M{
				"command": c,
			})
		}

		return "ok", nil
	default:
		return "", fmt.Errorf("command not found: %s", c.Name)
	}
}

func (c *Command) String() string {
	return fmt.Sprintf("[%s] %s %s %v", c.Type, c.Name, c.Key, c.Args)
}
