package calcapi

import (
	"time"

	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	"ultre.me/calcbiz/pkg/dashboard"
	"ultre.me/calcbiz/pkg/soundcloud"
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
	db         *gorm.DB
}

func New(db *gorm.DB, opts Options) (ServiceServer, error) {
	if err := setupDB(db); err != nil {
		return nil, err
	}

	svc := &svc{opts: opts, startTime: time.Now(), db: db}
	svc.soundcloud = soundcloud.New(opts.SoundcloudClientID, uint64(opts.SoundcloudUserID))
	svc.dashboard = dashboard.New(&dashboard.Options{Soundcloud: svc.soundcloud})
	// svc.dashboard.SetSoundCloud(&soundcloud)
	return svc, nil
}
