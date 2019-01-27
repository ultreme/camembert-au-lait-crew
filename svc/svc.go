package svc // import "ultre.me/calcbiz/svc"

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	tpyo "github.com/tpyolang/tpyo-cli"
	"go.uber.org/zap"

	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/pkg/crew"
	"ultre.me/calcbiz/pkg/dashboard"
	"ultre.me/calcbiz/pkg/numberinfo"
	"ultre.me/calcbiz/pkg/random"
	"ultre.me/calcbiz/pkg/soundcloud"
	"ultre.me/kryptos"
	"ultre.me/recettator"
)

type Options struct {
	SoundcloudUserID   int
	SoundcloudClientID string
	StaticBox          *packr.Box
}

type svc struct {
	opts       Options
	soundcloud *soundcloud.Soundcloud
	dashboard  *dashboard.Dashboard
	startTime  time.Time
}

func New(opts Options) (api.ServerServer, error) {
	svc := &svc{
		opts:      opts,
		startTime: time.Now(),
	}
	svc.soundcloud = soundcloud.New(
		opts.SoundcloudClientID,
		uint64(opts.SoundcloudUserID),
	)
	svc.dashboard = dashboard.New(&dashboard.Options{
		Soundcloud: svc.soundcloud,
	})
	// svc.dashboard.SetSoundCloud(&soundcloud)
	return svc, nil
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
	enedocr := tpyo.NewTpyo()
	return &api.TpyoEnocdeOuptut{
		To: enedocr.Enocde(input.Form),
	}, nil
}

func (svc *svc) Dashboard(_ context.Context, input *api.Void) (*dashboard.Entries, error) {
	return svc.dashboard.Random()
}

func (svc *svc) Hackz(_ context.Context, input *api.Void) (*dashboard.Entries, error) {
	return svc.dashboard.Hackz()
}

func (svc *svc) Crew(_ context.Context, input *api.Void) (*crew.Crew, error) {
	return &crew.CALC, nil
}

func (svc *svc) Numberinfo(_ context.Context, input *api.NumberinfoInput) (*api.NumberinfoOutput, error) {
	// FIXME: validate: input.Number is mandatory
	facts := map[string]string{}
	for k, v := range numberinfo.New(float64(input.Number)).All() {
		facts[k] = fmt.Sprintf("%v", v)
	}
	return &api.NumberinfoOutput{Facts: facts}, nil
}

func (svc *svc) Recettator(_ context.Context, input *api.RecettatorInput) (*api.RecettatorOutput, error) {
	if input.Seed == 0 {
		input.Seed = int64(rand.Intn(1000))
	}
	if input.Steps == 0 {
		input.Steps = uint64(rand.Intn(4) + 3)
	}
	if input.MainIngredients == 0 {
		input.MainIngredients = uint64(rand.Intn(2) + 1)
	}
	if input.SecondaryIngredients == 0 {
		input.SecondaryIngredients = uint64(rand.Intn(2) + 1)
	}
	rctt := recettator.New(input.Seed)
	rctt.SetSettings(recettator.Settings{
		MainIngredients:      input.MainIngredients,
		SecondaryIngredients: input.SecondaryIngredients,
		Steps:                input.Steps,
	})

	markdown, err := rctt.Markdown()
	if err != nil {
		zap.L().Warn("failed to marshal recettator in markdown", zap.Error(err))
	}
	output := &api.RecettatorOutput{
		Title:                rctt.Title(),
		People:               rctt.People(),
		Markdown:             markdown,
		Steps:                rctt.Steps(),
		Seed:                 input.Seed,
		MainIngredients:      []*api.RecettatorIngredient{},
		SecondaryIngredients: []*api.RecettatorIngredient{},
	}
	for _, ingredient := range rctt.Pool().MainIngredients.Picked {
		output.MainIngredients = append(output.MainIngredients, &api.RecettatorIngredient{
			Name: ingredient.Name(),
			// Quantity:        ingredient.Quantity(),
			NameAndQuantity: ingredient.NameAndQuantity(),
			Kind:            ingredient.Kind(),
			Method:          ingredient.GetMethod().NameAndQuantity(),
			Gender:          ingredient.GetGender(),
			Multiple:        ingredient.IsMultiple(),
		})
	}

	// fmt.Println(rctt.JSON())

	return output, nil
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
	return &api.WotdOutput{
		Word: random.WOTD(),
	}, nil
}

func (svc *svc) AlternateLogo(_ context.Context, input *api.Void) (*api.AlternateLogoOutput, error) {
	return &api.AlternateLogoOutput{
		Path: random.AlternateLogo(),
	}, nil
}

func (svc *svc) SoundcloudMe(_ context.Context, input *api.Void) (*soundcloud.User, error) {
	return svc.soundcloud.Me()
}

func (svc *svc) SoundcloudPlaylists(_ context.Context, input *api.Void) (*soundcloud.Playlists, error) {
	return svc.soundcloud.GetPlaylists()
}

func (svc *svc) SoundcloudPlaylist(_ context.Context, input *api.SoundcloudPlaylistInput) (*soundcloud.Playlist, error) {
	if input.PlaylistId < 1 { // pick random
		return svc.soundcloud.GetRandomPlaylist()
	}
	return svc.soundcloud.GetPlaylist(input.PlaylistId)
}

func (svc *svc) SoundcloudTracks(_ context.Context, input *api.Void) (*soundcloud.Tracks, error) {
	return svc.soundcloud.GetTracks()
}

func (svc *svc) SoundcloudTrack(_ context.Context, input *api.SoundcloudTrackInput) (*soundcloud.Track, error) {
	if input.TrackId < 1 { // pick random
		return svc.soundcloud.GetRandomTrack()
	}
	return svc.soundcloud.GetTrack(input.TrackId)
}

func (svc *svc) Metrics(_ context.Context, input *api.Void) (*api.MetricsOutput, error) {
	staticBoxSize := 0
	svc.opts.StaticBox.Walk(func(filepath string, file packd.File) error {
		staticBoxSize++
		return nil
	})
	out := &api.MetricsOutput{
		StaticBoxSize:     int32(staticBoxSize),
		ServerStartTime:   svc.startTime.Format(time.RFC3339),
		ServerCurrentTime: time.Now().Format(time.RFC3339),
		ServerUptime:      time.Since(svc.startTime).String(),
	}
	return out, nil
}
