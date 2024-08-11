package listeners

import (
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/rdb"
)

func LogWriteCommand(data map[string]any, numberOfWriteCommands int) {
	command, _ := data["command"].(string)
	key, _ := data["key"].(string)
	args, _ := data["args"].([]int64)

	aol.Write(command, key, args)

	if numberOfWriteCommands%10 == 0 {
		rdb.Persist()
	}
}