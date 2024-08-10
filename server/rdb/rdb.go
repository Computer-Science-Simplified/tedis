package rdb

import (
	"fmt"
	"mmartinjoo/trees/factory"
)

func Persist() {
	for _, item := range factory.Store {
		if item.Type == "binary_tree" {
			values := item.Value.ToArray()

			fmt.Println(values)
		}
	}
}