package main

import (
	"log"
	"os"
	"runshell/lib/action"

	"github.com/urfave/cli/v2"
)

func main() {
	var file string = ""

	app := cli.App{
		Name:  "runshell",
		Usage: "Run commands",
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
			action.Run(file)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
