package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "autocv",
		Usage: "this is the usage",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "data",
				Aliases:  []string{"d"},
				Usage:    "Usage of data",
				Required: true,
			},
			&cli.PathFlag{
				Name:     "template",
				Aliases:  []string{"t"},
				Usage:    "Usage of template",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			var dataPath string
			if cCtx.Path("data") != "" {
				dataPath = cCtx.Path("data")
			}
			fmt.Println("dataPath:", dataPath)

			var templatePath string
			if cCtx.Path("template") != "" {
				templatePath = cCtx.Path("template")
			}
			fmt.Println("templatePath:", templatePath)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
