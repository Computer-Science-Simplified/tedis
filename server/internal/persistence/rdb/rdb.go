package rdb

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/command"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
	"github.com/Computer-Science-Simplified/tedis/server/internal/model"
	"github.com/Computer-Science-Simplified/tedis/server/internal/store"
	"io"
	"os"
	"path/filepath"
)

const fileName = "./resources/rdb.bin"

var CurrentUnsavedWriteCommands = 0
var MaxUnsavedWriteCommands = 3

func ShouldPersist() bool {
	return CurrentUnsavedWriteCommands%MaxUnsavedWriteCommands == 0
}

func Persist() error {
	fmt.Println("RDB persisting to disk")
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	for _, key := range store.Keys() {
		tree, ok := store.Get(key)

		if !ok {
			continue
		}

		err = persistTree(tree, file)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	CurrentUnsavedWriteCommands = 0

	return nil
}

func Reload() (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, errors.New("couldn't read RDB")
	}

	numberOfKeys := 0

	for {
		var keyLength byte

		err = binary.Read(file, binary.LittleEndian, &keyLength)

		if err != nil {
			if err == io.EOF {
				break
			}

			return numberOfKeys, errors.New("rdb cannot be loaded. skipping reload")
		}

		var keyName string
		for i := 0; i < int(keyLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return numberOfKeys, errors.New("rdb cannot be loaded. skipping reload")
			}

			keyName += string(c)
		}

		var treeTypeLength byte

		err = binary.Read(file, binary.LittleEndian, &treeTypeLength)

		if err != nil {
			return numberOfKeys, errors.New("rdb cannot be loaded. skipping reload")
		}

		var treeType string
		for i := 0; i < int(treeTypeLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return numberOfKeys, errors.New("rdb cannot be loaded. skipping reload")
			}

			treeType += string(c)
		}

		var valuesLength int64

		err = binary.Read(file, binary.LittleEndian, &valuesLength)

		if err != nil {
			return numberOfKeys, errors.New("rdb cannot be loaded. skipping reload")
		}

		var values []int64

		for i := 0; i < int(valuesLength); i++ {
			var value int64
			err = binary.Read(file, binary.LittleEndian, &value)

			if err != nil {
				return numberOfKeys, errors.New("rdb cannot be loaded. skipping replay")
			}

			values = append(values, value)
		}

		for _, value := range values {
			err = executeCommand(keyName, value, treeType)
			if err != nil {
				return numberOfKeys, err
			}
		}

		numberOfKeys++
	}

	err = file.Close()
	if err != nil {
		return numberOfKeys, err
	}

	return numberOfKeys, nil
}

func persistTree(tree model.Tree, file *os.File) error {
	values := tree.GetAll()

	length := len(values)

	writer := bufio.NewWriter(file)

	err := binary.Write(writer, binary.LittleEndian, byte(len(tree.GetKey())))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(tree.GetKey()))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(tree.GetType())))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(tree.GetType()))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, int64(length))
	if err != nil {
		return err
	}

	for _, value := range values {
		err = binary.Write(writer, binary.LittleEndian, value)
		if err != nil {
			return err
		}
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}

func executeCommand(keyName string, value int64, treeType string) error {
	var cmdName string

	switch treeType {
	case enum.BinarySearchTree:
		cmdName = enum.BSTADD
	case enum.BinaryTree:
		cmdName = enum.BTADD
	}

	cmd, err := command.Create(cmdName, keyName, []int64{value})
	if err != nil {
		return err
	}

	_, err = cmd.Execute(false)
	if err != nil {
		return err
	}

	return nil
}
