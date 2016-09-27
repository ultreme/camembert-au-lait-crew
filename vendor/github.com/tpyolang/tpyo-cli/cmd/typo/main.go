package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/codegangsta/cli"
	"github.com/tpyolang/tpyo-cli"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// main is the entrypoint
func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Tpyo atuhros"
	app.Email = "https://github.com/tpyolang/tpyo-cli"
	app.Version = "1.0.0"
	app.Usage = "Mkae tpyos"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "symbol",
			Usage: "Changes characters by others",
		},
		cli.BoolFlag{
			Name:  "keyboard",
			Usage: "Mkae keraybod-louayt bsaed tyops",
		},
	}

	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	tpyo := tpyo.NewTpyo()
	tpyo.Smybol = c.BoolT("symbol")
	tpyo.Kraoybed = c.BoolT("keyboard")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(tpyo.Enocde(line))
	}
}
