package aolreplayer

import (
	"fmt"
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/factory"
)

func Replay() {
	commands, err := aol.Read()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	for _, command := range commands {
		tree := factory.Create(command.Key, "binary_tree")

		if command.Name == "BTADD" {
			tree.Insert(command.Args[0], tree.Root)
		}
	}
}