package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/camembertaulaitcrew/camembert-au-lait-crew"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/crew"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/log"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/random"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/soundcloud"
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
				cli.StringFlag{
					Name:   "soundcloud-client-id",
					Value:  "<configure-me>",
					Usage:  "SoundCloud CLIENT_ID",
					EnvVar: "SOUNDCLOUD_CLIENT_ID",
				},
				cli.IntFlag{
					Name:   "soundcloud-user-id",
					Value:  96137699,
					Usage:  "SoundCloud USER_ID",
					EnvVar: "SOUNDCLOUD_USER_ID",
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

	r.GET("/api/random/wotd", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": calcrand.WOTD(),
		})
	})

	soundcloud := calcsoundcloud.New(c.String("soundcloud-client-id"), uint64(c.Int("soundcloud-user-id")))
	r.GET("/api/soundcloud/me", func(c *gin.Context) {
		me, err := soundcloud.Me()
		if err != nil {
			log.Warnf("failed to get /api/soundcloud/me: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": me,
			})
		}
	})
	r.GET("/api/soundcloud/playlists", func(c *gin.Context) {
		playlists, err := soundcloud.Playlists()
		if err != nil {
			log.Warnf("failed to get /api/soundcloud/playlists: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": playlists,
			})
		}
	})
	r.GET("/api/soundcloud/playlists/:id", func(c *gin.Context) {
		playlistID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Invalid playlist id: %v", c.Param("id")),
			})
			return
		}

		playlist, err := soundcloud.Playlist(playlistID)
		if err != nil {
			log.Warnf("failed to get /api/soundcloud/playlists/%d: %v", playlistID, err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": playlist,
			})
		}
	})
	r.GET("/api/soundcloud/tracks/:id", func(c *gin.Context) {
		trackID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Invalid track id: %v", c.Param("id")),
			})
			return
		}

		track, err := soundcloud.Track(trackID)
		if err != nil {
			log.Warnf("failed to get /api/soundcloud/tracks/%d: %v", trackID, err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": track,
			})
		}
	})
	r.GET("/api/soundcloud/tracks", func(c *gin.Context) {
		tracks, err := soundcloud.Tracks()
		if err != nil {
			log.Warnf("failed to get /api/soundcloud/tracks: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": tracks,
			})
		}
	})

	// FIXME: handle socket.io
	http.Handle("/", r)
	log.Debugf("Listening and serving HTTP on %s", c.String("bind-address"))
	return http.ListenAndServe(c.String("bind-address"), nil)
}
