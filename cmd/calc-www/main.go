package main

import (
	"net/http"
	"os"

	"github.com/camembertaulaitcrew/camembert-au-lait-crew"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/crew"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/log"
	"github.com/camembertaulaitcrew/moi-j-aime-generator"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "calc-www"
	app.Usage = "Camembert au lait crew's web server"
	app.Version = calc.VERSION

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetDebug(true)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "Start a calc-www server",
			Action: server,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bind-address, b",
					Usage: "TCP port address",
					Value: ":9000",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}

func server(c *cli.Context) error {
	r := gin.Default()
	pong := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "pong",
		})
	}
	r.GET("/api/ping", pong)
	r.POST("/api/ping", pong)
	r.PUT("/api/ping", pong)
	r.PATCH("/api/ping", pong)
	r.DELETE("/api/ping", pong)

	r.GET("/api/crew", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": calccrew.CALC,
		})
	})

	r.GET("/api/moijaime", func(c *gin.Context) {
		phrases := []string{}
		for i := 0; i < 20; i++ {
			phrases = append(phrases, moijaime.Generate())
		}
		c.JSON(http.StatusOK, gin.H{
			"result": phrases,
		})
	})

	// FIXME: handle socket.io
	http.Handle("/", r)
	log.Debugf("Listening and serving HTTP on %s", c.String("bind-address"))
	return http.ListenAndServe(c.String("bind-address"), nil)
}
