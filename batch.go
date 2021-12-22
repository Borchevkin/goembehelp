package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func BatchSuffix(c *cli.Context) error {
	var dirPath string = ""
	var suffix string = ""
	var extension string = ""

	if c.Args().Len() != 3 {
		return errors.New("wrong count of arguments")
	}

	dirPath = c.Args().Get(0)
	if dirPath == "" {
		return errors.New("empty dir path")
	}

	suffix = c.Args().Get(1)
	if suffix == "" {
		return errors.New("empty suffix")
	}

	extension = c.Args().Get(2)
	if extension == "" {
		return errors.New("empty extension")
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == extension {
			oldName := f.Name()[:len(f.Name())-len(extension)] + extension
			newName := f.Name()[:len(f.Name())-len(extension)] + suffix + extension
			fmt.Println("Rename: " + oldName + " -> " + newName)

			err = os.Rename(filepath.Join(dirPath, oldName), filepath.Join(dirPath, newName))
			if nil != err {
				return err
			}
		}
	}

	return nil
}
