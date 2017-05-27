package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "doktor"
	app.Version = "0.0.1"
	app.Usage = "take two and call me in the morning (⌐■_■)"
	app.Description = "malware scanning and removal tool"

	app.Commands = []cli.Command{
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "clean all infections doktor finds",
			Action:  Clean,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all the infections doktor about",
			Action:  List,
		},
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "doktor will scan the host for known malware signatures",
			Action:  Scan,
		},
	}

	app.Run(os.Args)
}
