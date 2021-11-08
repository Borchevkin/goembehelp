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
		Copyright: "MIT-2 License",
		HelpName:  "goembehelp",
		Usage:     "Helper tools for embedded developer",
		UsageText: "type <goembehelp> and check available options",
		ArgsUsage: "[folder]",

		HideHelp:    false,
		HideVersion: false,

		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Command <%q> not found.\n", command)
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
			{
				Name:  "changelog",
				Usage: "CHANGELOG.md-related helpers",
				Subcommands: []*cli.Command{
					{
						Name:   "lastversion",
						Usage:  "Extract last version from CHANGELOG.md",
						Action: ChangelogExtractVersion,
					},
				},
			},
			{
				Name:  "cred",
				Usage: "Credentials-realated helpers",
				Subcommands: []*cli.Command{
					{
						Name:   "cstring",
						Usage:  "Extract a credential from file and print it as C-string",
						Action: CredGetCstring,
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
