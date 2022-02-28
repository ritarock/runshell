package main

import (
	"log"
	"os"
	"ritarock/runshell/internal/action"

	"github.com/urfave/cli/v2"
)

const VERSION = "1.0"

func main() {
	var file string = ""

	app := cli.App{
		Name:    "runshell",
		Usage:   "Run commands",
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Usage:       "set file",
				Destination: &file,
				Aliases:     []string{"f"},
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			if err := action.Run(file); err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
