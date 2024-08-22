package command

import "github.com/Computer-Science-Simplified/tedis/server/internal/types"

type Command interface {
	Execute(shouldReport bool) (string, error)
	GetParams() *types.CommandParams
	String() string
}

//type CommandParams struct {
//	Name string
//	Key  string
//	Args []int64
//	Type string
//}

//func (c *CommandParams) Execute(shouldReport bool) (string, error) {
//	t, err := tree.Create(c.Key, c.Type)
//
//	if err != nil {
//		return "", err
//	}
//
//	var result string
//
//	switch c.Name {
//	case enum.BSTADD:
//		t.Add(c.Args[0])
//
//		if shouldReport {
//			event.MustFire(enum.WriteCommandExecuted, event.M{
//				"command": c,
//			})
//		}
//
//		result = "ok"
//
//	case enum.BSTEXISTS:
//		exists := t.Exists(c.Args[0])
//		result = strconv.FormatBool(exists)
//
//	case enum.BSTGETALL:
//		values := t.GetAll()
//		result = fmt.Sprintf("%v", values)
//
//	case enum.BSTREM:
//		t.Remove(c.Args[0])
//
//		if shouldReport {
//			event.MustFire(enum.WriteCommandExecuted, event.M{
//				"command": c,
//			})
//		}
//
//		result = "ok"
//	default:
//		result = ""
//		err = fmt.Errorf("command not found: %s", c.Name)
//	}
//
//	event.MustFire(enum.CommandExecuted, event.M{
//		"command": c,
//	})
//
//	return result, err
//}
//
//func (c *CommandParams) String() string {
//	return fmt.Sprintf("[%s] %s %s %v", c.Type, c.Name, c.Key, c.Args)
//}
