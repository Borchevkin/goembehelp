package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FirmwareInit doing the following:
// - Create folders according to internal standart
// - Create gitignore
// - Create clang-format
// - Create README.md
// - create CHANGELOG.md
// - create doxygen docs
func FirmwareInit(c *cli.Context) error {
	fmt.Println("Scaffold project's dirs according to the HiQo IoT Embedded Project Structure v1.5")

	var err error = nil
	var dirPath string = ""
	var rootPath string = ""

	if c.Args().First() != "" {
		rootPath = c.Args().First()
	} else {
		rootPath = "./"
	}

	dirPath = filepath.Join(rootPath, "sources", "application")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sources", "middleware")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sources", "hardware", "bsp")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sources", "hardware", "drivers")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sources", "hardware", "mcu_hal")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sources", "hardware", "system")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "vendor_sdk")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "project", "crossworks")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "sdk_utils")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	dirPath = filepath.Join(rootPath, "docs")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	var staticData []byte
	staticData, err = Asset("static/README.md")
	err = ioutil.WriteFile(filepath.Join(rootPath, "README.md"), staticData, 0644)
	if err != nil {
		return err
	}

	staticData, err = Asset("static/gitignore")
	err = ioutil.WriteFile(filepath.Join(rootPath, ".gitignore"), staticData, 0644)
	if err != nil {
		return err
	}

	staticData, err = Asset("static/clang-format")
	err = ioutil.WriteFile(filepath.Join(rootPath, ".clang-format"), staticData, 0644)
	if err != nil {
		return err
	}

	staticData, err = Asset("static/CHANGELOG.md")
	err = ioutil.WriteFile(filepath.Join(rootPath, "CHANGELOG.md"), staticData, 0644)
	if err != nil {
		return err
	}

	staticData, err = Asset("static/doxygen_doc")
	err = ioutil.WriteFile(filepath.Join(rootPath, "docs", "doxygen_doc"), staticData, 0644)
	if err != nil {
		return err
	}

	staticData, err = Asset("static/index.md")
	err = ioutil.WriteFile(filepath.Join(rootPath, "docs", "INDEX.md"), staticData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Scaffold project's dirs and files was created")
	return err
}
