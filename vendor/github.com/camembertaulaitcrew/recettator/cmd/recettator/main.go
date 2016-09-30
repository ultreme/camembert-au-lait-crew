package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/camembertaulaitcrew/recettator"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "recettator"
	app.Usage = "Generate CALC recipes"
	app.Version = "master"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.IntFlag{
			Name:  "seed, s",
			Usage: "Set seed value",
		},
		cli.IntFlag{
			Name:  "main-ingredients",
			Usage: "Amount of main-ingredients",
		},
		cli.IntFlag{
			Name:  "secondary-ingredients",
			Usage: "Amount of secondary-ingredients",
		},
		cli.IntFlag{
			Name:  "steps",
			Usage: "Amount of steps",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Use JSON output",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		//panic(err)
		logrus.Fatalf("%v", err)
	}
}

func run(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	seed := int64(c.Int("seed"))
	if seed == 0 {
		seed = time.Now().UTC().UnixNano()
	}
	rctt := recettator.New(seed)

	rctt.SetSettings(recettator.Settings{
		MainIngredients:      uint64(c.Int("main-ingredients")),
		SecondaryIngredients: uint64(c.Int("secondary-ingredients")),
		Steps:                uint64(c.Int("steps")),
	})

	var output string
	var err error

	if c.Bool("json") {
		output = rctt.JSON()
	} else {
		output, err = rctt.Markdown()
		if err != nil {
			return err
			//panic(err)
		}
	}

	fmt.Println(output)

	return nil
}
