package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
	"strings"
)

func ChangelogExtractVersion(c *cli.Context) error {
	var filepath string = ""

	if c.Args().First() != "" {
		filepath = c.Args().First()
	} else {
		filepath = "./CHANGELOG.md"
	}

	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("cannot open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()

	startIdx := strings.Index(firstLine, "[")
	if startIdx == -1 {
		return errors.New("no left <[> in first string of the changelog")
	}
	startIdx++

	endIdx := strings.Index(firstLine, "]")
	if endIdx == -1 {
		return errors.New("no right <]> in first string of the changelog")
	}

	version := firstLine[startIdx:endIdx]

	fmt.Println(version)

	return nil
}
