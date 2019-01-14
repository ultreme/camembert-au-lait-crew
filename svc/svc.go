package svc // import "ultre.me/calcbiz/svc"

import (
	"context"
	"fmt"

	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/pkg/crew"
	"ultre.me/kryptos"
)

type Options struct {
	SoundcloudUserID   int
	SoundcloudClientID string
}

type svc struct {
	opts Options
	// soundcloud
	// dashboard
}

func New(opts Options) (api.ServerServer, error) {
	//soundcloud := calcsoundcloud.New(c.String("soundcloud-client-id"), uint64(c.Int("soundcloud-user-id")))
	//dashboard := calcdashboard.New()
	//dashboard.SetSoundCloud(&soundcloud)

	return &svc{opts: opts}, nil
}

func (svc *svc) Ping(_ context.Context, input *api.Void) (*api.Pong, error) {
	return &api.Pong{Pong: "pong"}, nil
}

func (svc *svc) KryptosEncrypt(_ context.Context, input *api.KryptosInput) (*api.KryptosOutput, error) {
	return &api.KryptosOutput{
		To: kryptos.Encrypt(input.From),
	}, nil
}

func (svc *svc) KryptosDecrypt(_ context.Context, input *api.KryptosInput) (*api.KryptosOutput, error) {
	return &api.KryptosOutput{
		To: kryptos.Decrypt(input.From),
	}, nil
}

func (svc *svc) TpyoEnocde(_ context.Context, input *api.TpyoEnocdeIpunt) (*api.TpyoEnocdeOuptut, error) {
	/*
		r.Post("/tpyo/enocde", func(w http.ResponseWriter, r *http.Request) {
			var data struct {
				Message string
			}
			if err := c.BindJSON(&data); err == nil {
				enedocr := tpyo.NewTpyo()
				c.JSON(http.StatusOK, gin.H{
					"result": enedocr.Enocde(data.Message),
				})
			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf("Invalid input: %v", err),
				})
			}
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Dashboard(_ context.Context, input *api.Void) (*api.DashboardOutput, error) {
	/*
		r.Get("/dashboard/random", func(w http.ResponseWriter, r *http.Request) {
			dashboard, err := dashboard.Random()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"result": dashboard,
				})
			}
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Crew(_ context.Context, input *api.Void) (*crew.Crew, error) {
	return &crew.CALC, nil
}

func (svc *svc) Numberinfo(_ context.Context, input *api.NumberinfoInput) (*api.NumberinfoOutput, error) {
	/*
		r.Get("/numberinfo/all/:number", func(w http.ResponseWriter, r *http.Request) {
			number, err := strconv.ParseFloat(c.Param("number"), 64)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf("Invalid number: %v (%v)", c.Param("number"), err),
				})
				return
			}

			info := calcnumberinfo.New(number).All()
			c.JSON(http.StatusOK, gin.H{
				"result": info,
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Recettator(_ context.Context, input *api.RecettatorInput) (*api.RecettatorOutput, error) {
	/*
		r.Get("/recettator/json/:seed", func(w http.ResponseWriter, r *http.Request) {
			seed, err := strconv.ParseInt(c.Param("seed"), 10, 64)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf("Invalid seed: %v (%v)", c.Param("seed"), err),
				})
				return
			}

			rctt := recettator.New(seed)
			rctt.SetSettings(recettator.Settings{
				MainIngredients:      2,
				SecondaryIngredients: 2,
				Steps:                5,
			})

			output := rctt.ToMap()

			c.JSON(http.StatusOK, gin.H{
				"result": output,
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Moijaime(_ context.Context, input *api.Void) (*api.MoijaimeOutput, error) {
	/*
		r.Get("/moijaime", func(w http.ResponseWriter, r *http.Request) {
			phrases := []string{}
			for i := 0; i < 20; i++ {
				phrases = append(phrases, moijaime.Generate())
			}
			c.JSON(http.StatusOK, gin.H{
				"result": phrases,
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtRandom(_ context.Context, input *api.Void) (*api.SpreadshirtRandomOutput, error) {
	/*
		r.Get("/spreadshirt/random", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetRandomProduct(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtAll(_ context.Context, input *api.Void) (*api.SpreadshirtAllOutput, error) {
	/*
		r.Get("/spreadshirt/all", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetAllProducts(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Wotd(_ context.Context, input *api.Void) (*api.WotdOutput, error) {
	/*
		r.Get("/random/wotd", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcrand.WOTD(),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) AlternateLogo(_ context.Context, input *api.Void) (*api.AlternateLogoOutput, error) {
	/*
		r.Get("/random/alternate-logo", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcrand.AlternateLogo(),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudMe(_ context.Context, input *api.Void) (*api.SoundcloudMeOutput, error) {
	/*
		r.Get("/soundcloud/me", func(w http.ResponseWriter, r *http.Request) {
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
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudPlaylists(_ context.Context, input *api.Void) (*api.SoundcloudPlaylistsOutput, error) {
	/*
		r.Get("/soundcloud/playlists", func(w http.ResponseWriter, r *http.Request) {
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
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudPlaylist(_ context.Context, input *api.SoundcloudPlaylistInput) (*api.SoundcloudPlaylistOutput, error) {
	/*
		r.Get("/soundcloud/playlists/:id", func(w http.ResponseWriter, r *http.Request) {
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
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudTracks(_ context.Context, input *api.Void) (*api.SoundcloudTracksOutput, error) {
	/*
		r.Get("/soundcloud/tracks", func(w http.ResponseWriter, r *http.Request) {
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
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudTrack(_ context.Context, input *api.SoundcloudTrackInput) (*api.SoundcloudTrackOutput, error) {
	/*
		r.Get("/soundcloud/tracks/:id", func(w http.ResponseWriter, r *http.Request) {
			if c.Param("id") == "random" {
				track, err := soundcloud.RandomTrack()
				if err != nil {
					log.Warnf("failed to get /api/soundcloud/tracks/random: %v", err)
					c.JSON(http.StatusNotFound, gin.H{
						"error": soundcloud.EscapeString(fmt.Sprintf("%v", err)),
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"result": track,
					})
				}
				return
			}
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
	*/
	return nil, fmt.Errorf("not implemented")
}
