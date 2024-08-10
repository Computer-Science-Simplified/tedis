package aolreplayer

import (
	"mmartinjoo/trees/aol"
	"mmartinjoo/trees/factory"
)

func Replay() {
	commands := aol.Read()

	for _, command := range commands {
		tree := factory.Create(command.Key, "binary_tree")

		if command.Name == "BTADD" {
			tree.Insert(command.Args[0], tree.Root)
		}
	}
}