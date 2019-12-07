package calcapi

import (
	"time"

	"github.com/gobuffalo/packr"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"ultre.me/calcbiz/pkg/dashboard"
	"ultre.me/calcbiz/pkg/soundcloud"
)

type Options struct {
	SoundcloudUserID   int
	SoundcloudClientID string
	StaticBox          *packr.Box
	Logger             *zap.Logger
}

type svc struct {
	opts       Options
	soundcloud *soundcloud.Soundcloud
	dashboard  *dashboard.Dashboard
	startTime  time.Time
	db         *gorm.DB
}

type Service interface {
	ServiceServer

	SocketIOServer() (*socketio.Server, error)
}

func New(db *gorm.DB, opts Options) (Service, error) {
	if err := setupDB(db); err != nil {
		return nil, err
	}

	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}

	svc := &svc{opts: opts, startTime: time.Now(), db: db}
	svc.soundcloud = soundcloud.New(opts.SoundcloudClientID, uint64(opts.SoundcloudUserID))
	svc.dashboard = dashboard.New(&dashboard.Options{Soundcloud: svc.soundcloud})
	// svc.dashboard.SetSoundCloud(&soundcloud)
	return svc, nil
}
