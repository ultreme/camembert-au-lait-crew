package calcapi

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (svc *svc) SocketIOServer() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, errors.Wrap(err, "socketio.NewServer")
	}

	logger := svc.opts.Logger.Named("sio")

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		logger.Info("connected", zap.String("id", s.ID()))
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		logger.Debug("notice", zap.String("msg", msg))
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(err error) {
		logger.Warn("error", zap.Error(err))
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		logger.Info("closed", zap.String("reason", reason))
	})

	return server, nil
}
