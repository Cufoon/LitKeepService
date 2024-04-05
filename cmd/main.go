package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"cufoon.litkeep.service/internal"
)

func main() {
	println("<-----LitKeep for Ms. Tang!----->")
	app := &cli.App{
		Name:  "LitKeep Backend Service",
		Usage: "Just record your life bill!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./dev.yaml",
				Usage:   "config file to run the backend service",
				EnvVars: []string{"LITKEEP_CONFIG"},
			},
		},
		Action: func(cCtx *cli.Context) error {
			cf := cCtx.String("config")
			internal.StartAPP(cf)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
