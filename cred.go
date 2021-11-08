package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
)

func CredGetCstring(c *cli.Context) error {
	var filepath string = ""

	if c.Args().First() == "" {
		return errors.New("no file provided")
	}

	filepath = c.Args().First()

	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("cannot open provided file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var isFirstLine bool = true
	for scanner.Scan() {
		if isFirstLine {
			fmt.Println()
			isFirstLine = false
		} else {
			fmt.Print(" \\")
			fmt.Println()
		}

		line := scanner.Text()
		fmt.Print("\"", line, "\\n\"")
	}
	fmt.Print(";")
	fmt.Println()
	fmt.Println()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
