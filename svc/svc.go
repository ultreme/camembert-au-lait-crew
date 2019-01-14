package svc // import "ultre.me/calcbiz/svc"

import (
	"context"
	"fmt"

	"ultre.me/calcbiz/api"
)

type Options struct {
	SoundcloudUserID   int
	SoundcloudClientID string
}

type svc struct {
	opts Options
}

func New(opts Options) (api.ServerServer, error) {
	return &svc{opts: opts}, nil
}

func (svc *svc) KryptosEncrypt(context.Context, *api.KryptosInput) (*api.KryptosOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) KryptosDecrypt(context.Context, *api.KryptosInput) (*api.KryptosOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) TpyoEnocde(context.Context, *api.TpyoEnocdeIpunt) (*api.TpyoEnocdeOuptut, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Ping(context.Context, *api.Void) (*api.Void, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Dashboard(context.Context, *api.Void) (*api.DashboardOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Crew(context.Context, *api.Void) (*api.CrewOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Numberinfo(context.Context, *api.NumberinfoInput) (*api.NumberinfoOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Recettator(context.Context, *api.RecettatorInput) (*api.RecettatorOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Moijaime(context.Context, *api.Void) (*api.MoijaimeOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtRandom(context.Context, *api.Void) (*api.SpreadshirtRandomOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtAll(context.Context, *api.Void) (*api.SpreadshirtAllOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Wotd(context.Context, *api.Void) (*api.WotdOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) AlternateLogo(context.Context, *api.Void) (*api.AlternateLogoOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudMe(context.Context, *api.Void) (*api.SoundcloudMeOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudPlaylists(context.Context, *api.Void) (*api.SoundcloudPlaylistsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudPlaylist(context.Context, *api.SoundcloudPlaylistInput) (*api.SoundcloudPlaylistOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudTracks(context.Context, *api.Void) (*api.SoundcloudTracksOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SoundcloudTrack(context.Context, *api.SoundcloudTrackInput) (*api.SoundcloudTrackOutput, error) {
	return nil, fmt.Errorf("not implemented")
}

/*
	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		// ping
		pong := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"result":"pong"}`))
		}
		r.Get("/ping", pong)
		r.Post("/ping", pong)
		r.Options("/ping", pong)
		r.Trace("/ping", pong)
		r.Connect("/ping", pong)
		r.Head("/ping", pong)
		r.Put("/ping", pong)
		r.Patch("/ping", pong)
		r.Delete("/ping", pong)

		soundcloud := calcsoundcloud.New(c.String("soundcloud-client-id"), uint64(c.Int("soundcloud-user-id")))
		dashboard := calcdashboard.New()
		dashboard.SetSoundCloud(&soundcloud)

		/*
			// dashboard
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

				// crew
				r.Get("/crew", func(w http.ResponseWriter, r *http.Request) {
					c.JSON(http.StatusOK, gin.H{
						"result": calccrew.CALC,
					})
				})

				// numberinfo
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

				// recettator
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

				// moijaime
				r.Get("/moijaime", func(w http.ResponseWriter, r *http.Request) {
					phrases := []string{}
					for i := 0; i < 20; i++ {
						phrases = append(phrases, moijaime.Generate())
					}
					c.JSON(http.StatusOK, gin.H{
						"result": phrases,
					})
				})

				// spreadshirt
				r.Get("/spreadshirt/random", func(w http.ResponseWriter, r *http.Request) {
					c.JSON(http.StatusOK, gin.H{
						"result": calcspreadshirt.GetRandomProduct(250, 250),
					})
				})
				r.Get("/spreadshirt/all", func(w http.ResponseWriter, r *http.Request) {
					c.JSON(http.StatusOK, gin.H{
						"result": calcspreadshirt.GetAllProducts(250, 250),
					})
				})

				// kryptos
				r.Post("/kryptos/encrypt", func(w http.ResponseWriter, r *http.Request) {
					var data struct {
						Message string
					}
					if err := c.BindJSON(&data); err == nil {
						c.JSON(http.StatusOK, gin.H{
							"result": kryptos.Encrypt(data.Message),
						})
					} else {
						c.JSON(http.StatusNotFound, gin.H{
							"error": fmt.Sprintf("Invalid input: %v", err),
						})
					}
				})
				r.Post("/kryptos/decrypt", func(w http.ResponseWriter, r *http.Request) {
					var data struct {
						Message string
					}
					if err := c.BindJSON(&data); err == nil {
						c.JSON(http.StatusOK, gin.H{
							"result": kryptos.Decrypt(data.Message),
						})
					} else {
						c.JSON(http.StatusNotFound, gin.H{
							"error": fmt.Sprintf("Invalid input: %v", err),
						})
					}
				})

				// tpyo
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

				// random
				r.Get("/random/wotd", func(w http.ResponseWriter, r *http.Request) {
					c.JSON(http.StatusOK, gin.H{
						"result": calcrand.WOTD(),
					})
				})
				r.Get("/random/alternate-logo", func(w http.ResponseWriter, r *http.Request) {
					c.JSON(http.StatusOK, gin.H{
						"result": calcrand.AlternateLogo(),
					})
				})

				// soundcloud
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
