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
)

const fileName = "resources/rdb.bin"

func Persist() error {
	fmt.Println("RDB persisting to disk")

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	for _, key := range store.Keys() {
		tree, ok := store.Get(key)

		if !ok {
			continue
		}

		if tree.GetType() == enum.BinarySearchTree {
			persistTree(tree.GetKey(), tree, file)
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func Reload() (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, errors.New("couldn't read RDB")
	}

	numberOfItems := 0

	for {
		var keyLength byte

		err = binary.Read(file, binary.LittleEndian, &keyLength)

		if err != nil {
			if err == io.EOF {
				break
			}

			return numberOfItems, errors.New("rdb cannot be loaded. skipping reload")
		}

		var keyName string
		for i := 0; i < int(keyLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return numberOfItems, errors.New("rdb cannot be loaded. skipping reload")
			}

			keyName += string(c)
		}

		var treeTypeLength byte

		err = binary.Read(file, binary.LittleEndian, &treeTypeLength)

		if err != nil {
			return numberOfItems, errors.New("rdb cannot be loaded. skipping reload")
		}

		var treeType string
		for i := 0; i < int(treeTypeLength); i++ {
			var c byte
			err = binary.Read(file, binary.LittleEndian, &c)

			if err != nil {
				return numberOfItems, errors.New("rdb cannot be loaded. skipping reload")
			}

			treeType += string(c)
		}

		var valuesLength int64

		err = binary.Read(file, binary.LittleEndian, &valuesLength)

		if err != nil {
			return numberOfItems, errors.New("rdb cannot be loaded. skipping reload")
		}

		var values []int64

		for i := 0; i < int(valuesLength); i++ {
			var value int64
			err = binary.Read(file, binary.LittleEndian, &value)

			if err != nil {
				return numberOfItems, errors.New("rdb cannot be loaded. skipping replay")
			}

			values = append(values, value)
		}

		for _, value := range values {
			cmd := command.Command{
				Key:  keyName,
				Name: enum.BSTADD,
				Args: []int64{value},
				Type: enum.BinarySearchTree,
			}

			_, err := cmd.Execute()

			if err != nil {
				return numberOfItems, err
			}

			numberOfItems++
		}
	}

	fmt.Printf("Reloaded %d values\n", numberOfItems)

	err = file.Close()
	if err != nil {
		return numberOfItems, err
	}

	return numberOfItems, nil
}

func persistTree(key string, tree model.Tree, file *os.File) {
	values := tree.GetAll()

	length := len(values)

	writer := bufio.NewWriter(file)

	err := binary.Write(writer, binary.LittleEndian, byte(len(key)))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(key))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, byte(len(tree.GetType())))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, []byte(tree.GetType()))
	if err != nil {
		panic(err)
	}

	err = binary.Write(writer, binary.LittleEndian, int64(length))
	if err != nil {
		panic(err)
	}

	for _, value := range values {
		err = binary.Write(writer, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
	}

	err = writer.Flush()

	if err != nil {
		panic(err)
	}
}
