package main

import (
	"log"
	"os"

	peekcli "github.com/jonomuller/peek/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "peek",
		Usage:  "peek <filename>",
		Action: peekcli.Peek,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "max",
				Usage:   "max number of lines to peek",
				Aliases: []string{"n"},
				Value:   10,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
