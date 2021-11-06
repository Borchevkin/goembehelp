package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:     "goembehelp",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Danil Borchevkin",
				Email: "danil.borchevkin@gmail.com",
			},
		},
		Copyright: "MIT-2 Licence",
		HelpName:  "goembehelp",
		Usage:     "Helper tools for embedded developer",
		UsageText: "type <goembehelp> and check available options",
		ArgsUsage: "[folder]",

		HideHelp:    false,
		HideVersion: false,

		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
		},

		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "firmware",
				Aliases: []string{"fw"},
				Usage:   "Firmware-related helpers",
				Subcommands: []*cli.Command{
					{
						Name:   "init",
						Usage:  "Init folder structure for a firmware project",
						Action: FirmwareInit,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
