package calcapi // import "ultre.me/calcbiz/pkg/calcapi"

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	tpyo "github.com/tpyolang/tpyo-cli"
	"go.uber.org/zap"
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

func New(opts Options) (ServiceServer, error) {
	svc := &svc{opts: opts, startTime: time.Now()}
	svc.soundcloud = soundcloud.New(opts.SoundcloudClientID, uint64(opts.SoundcloudUserID))
	svc.dashboard = dashboard.New(&dashboard.Options{Soundcloud: svc.soundcloud})
	// svc.dashboard.SetSoundCloud(&soundcloud)
	return svc, nil
}

func (svc *svc) Ping(_ context.Context, input *Ping_Input) (*Ping_Output, error) {
	return &Ping_Output{Pong: "pong"}, nil
}

func (svc *svc) KryptosEncrypt(_ context.Context, input *Kryptos_Input) (*Kryptos_Output, error) {
	return &Kryptos_Output{To: kryptos.Encrypt(input.From)}, nil
}

func (svc *svc) KryptosDecrypt(_ context.Context, input *Kryptos_Input) (*Kryptos_Output, error) {
	return &Kryptos_Output{To: kryptos.Decrypt(input.From)}, nil
}

func (svc *svc) TpyoEnocde(_ context.Context, input *TpyoEnocde_Ipunt) (*TpyoEnocde_Ouptut, error) {
	enedocr := tpyo.NewTpyo()
	return &TpyoEnocde_Ouptut{To: enedocr.Enocde(input.Form)}, nil
}

func (svc *svc) Dashboard(_ context.Context, input *Dashboard_Input) (*Dashboard_Output, error) {
	entries, err := svc.dashboard.Random()
	if err != nil {
		return nil, err
	}
	return &Dashboard_Output{Entries: entries}, nil
}

func (svc *svc) Hackz(_ context.Context, input *Hackz_Input) (*Hackz_Output, error) {
	entries, err := svc.dashboard.Hackz()
	if err != nil {
		return nil, err
	}
	return &Hackz_Output{Entries: entries}, nil
}

func (svc *svc) Crew(_ context.Context, input *Crew_Input) (*Crew_Output, error) {
	return &Crew_Output{Crew: &crew.CALC}, nil
}

func (svc *svc) Numberinfo(_ context.Context, input *Numberinfo_Input) (*Numberinfo_Output, error) {
	// FIXME: validate: input.Number is mandatory
	facts := map[string]string{}
	for k, v := range numberinfo.New(float64(input.Number)).All() {
		facts[k] = fmt.Sprintf("%v", v)
	}
	return &Numberinfo_Output{Facts: facts}, nil
}

func (svc *svc) Recettator(_ context.Context, input *Recettator_Input) (*Recettator_Output, error) {
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
	output := &Recettator_Output{
		Title:                rctt.Title(),
		People:               rctt.People(),
		Markdown:             markdown,
		Steps:                rctt.Steps(),
		Seed:                 input.Seed,
		MainIngredients:      []*Recettator_Ingredient{},
		SecondaryIngredients: []*Recettator_Ingredient{},
	}
	for _, ingredient := range rctt.Pool().MainIngredients.Picked {
		output.MainIngredients = append(output.MainIngredients, &Recettator_Ingredient{
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

func (svc *svc) Moijaime(_ context.Context, input *Moijaime_Input) (*Moijaime_Output, error) {
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

func (svc *svc) SpreadshirtRandom(_ context.Context, input *SpreadshirtRandom_Input) (*SpreadshirtRandom_Output, error) {
	/*
		r.Get("/spreadshirt/random", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetRandomProduct(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtAll(_ context.Context, input *SpreadshirtAll_Input) (*SpreadshirtAll_Output, error) {
	/*
		r.Get("/spreadshirt/all", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetAllProducts(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Wotd(_ context.Context, input *Wotd_Input) (*Wotd_Output, error) {
	return &Wotd_Output{Word: random.WOTD()}, nil
}

func (svc *svc) AlternateLogo(_ context.Context, input *AlternateLogo_Input) (*AlternateLogo_Output, error) {
	return &AlternateLogo_Output{Path: random.AlternateLogo()}, nil
}

func (svc *svc) SoundcloudMe(_ context.Context, input *SoundcloudMe_Input) (*SoundcloudMe_Output, error) {
	me, err := svc.soundcloud.Me()
	if err != nil {
		return nil, err
	}
	return &SoundcloudMe_Output{Me: me}, nil
}

func (svc *svc) SoundcloudPlaylists(_ context.Context, input *SoundcloudPlaylists_Input) (*SoundcloudPlaylists_Output, error) {
	playlists, err := svc.soundcloud.GetPlaylists()
	if err != nil {
		return nil, err
	}
	return &SoundcloudPlaylists_Output{Playlists: playlists}, nil
}

func (svc *svc) SoundcloudPlaylist(_ context.Context, input *SoundcloudPlaylist_Input) (*SoundcloudPlaylist_Output, error) {
	if input.PlaylistId < 1 { // pick random
		playlist, err := svc.soundcloud.GetRandomPlaylist()
		if err != nil {
			return nil, err
		}
		return &SoundcloudPlaylist_Output{Playlist: playlist}, nil
	}
	playlist, err := svc.soundcloud.GetPlaylist(input.PlaylistId)
	if err != nil {
		return nil, err
	}
	return &SoundcloudPlaylist_Output{Playlist: playlist}, nil
}

func (svc *svc) SoundcloudTracks(_ context.Context, input *SoundcloudTracks_Input) (*SoundcloudTracks_Output, error) {
	tracks, err := svc.soundcloud.GetTracks()
	if err != nil {
		return nil, err
	}
	return &SoundcloudTracks_Output{Tracks: tracks}, nil
}

func (svc *svc) SoundcloudTrack(_ context.Context, input *SoundcloudTrack_Input) (*SoundcloudTrack_Output, error) {
	if input.TrackId < 1 { // pick random
		track, err := svc.soundcloud.GetRandomTrack()
		if err != nil {
			return nil, err
		}
		return &SoundcloudTrack_Output{Track: track}, nil
	}
	track, err := svc.soundcloud.GetTrack(input.TrackId)
	if err != nil {
		return nil, err
	}
	return &SoundcloudTrack_Output{Track: track}, nil
}

func (svc *svc) Metrics(_ context.Context, input *Metrics_Input) (*Metrics_Output, error) {
	staticBoxSize := 0
	err := svc.opts.StaticBox.Walk(func(filepath string, file packd.File) error {
		staticBoxSize++
		return nil
	})
	if err != nil {
		return nil, err
	}
	out := &Metrics_Output{
		StaticBoxSize:     int32(staticBoxSize),
		ServerStartTime:   svc.startTime.Format(time.RFC3339),
		ServerCurrentTime: time.Now().Format(time.RFC3339),
		ServerUptime:      time.Since(svc.startTime).String(),
	}
	return out, nil
}
