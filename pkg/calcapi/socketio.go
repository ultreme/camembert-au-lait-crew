package calcapi

import (
	"encoding/json"
	"strings"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (svc *svc) sioOnConnect(s socketio.Conn) error {
	s.SetContext(SIO_Context{NoMetadata: true})
	svc.sio.connectedPeers++
	svc.sio.logger.Info(
		"connect",
		zap.String("id", s.ID()),
		zap.Int("peers", svc.sio.connectedPeers),
	)
	return nil
}

func (svc *svc) sioOnError(err error) {
	svc.sio.logger.Warn("error", zap.Error(err))
}

func (svc *svc) sioOnDisconnect(s socketio.Conn, reason string) {
	svc.sio.connectedPeers--
	svc.sio.logger.Info(
		"disconnect",
		zap.String("id", s.ID()),
		zap.Int("peers", svc.sio.connectedPeers),
		zap.String("reason", reason),
	)

	context := s.Context().(SIO_Context)
	for _, room := range s.Rooms() {
		broadcast := SIO_Disconnect_Event{Room: room, Peer: context.Peer}
		out, _ := json.Marshal(broadcast)
		svc.sio.server.BroadcastToRoom(room, "event:disconnet", string(out))
	}
}

func (svc *svc) onEventPing(s socketio.Conn, msg string) string {
	svc.sio.logger.Debug("ping requested", zap.String("msg", msg), zap.String("id", s.ID()))
	return msg
}

func (svc *svc) onEventRooms(s socketio.Conn) string {
	svc.sio.logger.Debug("rooms requested", zap.String("id", s.ID()))
	return strings.Join(s.Rooms(), ",")
}

func (svc *svc) onEventJoin(s socketio.Conn, in *SIO_Join_Input) (*SIO_Join_Output, error) {
	svc.sio.logger.Debug(
		"join",
		zap.String("room", in.Room),
		zap.String("id", s.ID()),
	)

	s.Join(in.Room)

	context := s.Context().(SIO_Context)
	context.NoMetadata = false
	context.Peer = in.Peer
	s.SetContext(context)

	broadcast := SIO_Join_Event{Room: in.Room, Peer: in.Peer}
	out, _ := json.Marshal(broadcast)
	svc.sio.server.BroadcastToRoom(in.Room, "event:join", string(out))

	ret := SIO_Join_Output{
		Peers: []*SIO_Peer{{Name: "foo"}, {Name: "bar"}},
	}
	return &ret, nil
}

func (svc *svc) onEventBroadcast(s socketio.Conn, in *SIO_Broadcast_Input) (*SIO_Broadcast_Output, error) {
	svc.sio.logger.Debug(
		"broadcast",
		zap.String("room", in.Room),
		zap.String("msg", in.Msg),
		zap.String("id", s.ID()),
	)

	context := s.Context().(SIO_Context)

	broadcast := SIO_Broadcast_Event{Room: in.Room, Msg: in.Msg, Peer: context.Peer}
	out, _ := json.Marshal(broadcast)
	svc.sio.server.BroadcastToRoom(in.Room, "event:broadcast", string(out))

	ret := SIO_Broadcast_Output{}
	return &ret, nil
}

func (svc *svc) SocketIOServer() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, errors.Wrap(err, "socketio.NewServer")
	}
	svc.sio.logger = svc.opts.Logger.Named("sio")
	svc.sio.server = server

	// core events
	server.OnConnect("/", svc.sioOnConnect)
	server.OnError("/", svc.sioOnError)
	server.OnDisconnect("/", svc.sioOnDisconnect)

	// api
	server.OnEvent("/", "ping", svc.onEventPing)
	server.OnEvent("/", "rooms", svc.onEventRooms)
	server.OnEvent("/", "join", func(s socketio.Conn, msg string) string {
		var in SIO_Join_Input
		err := json.Unmarshal([]byte(msg), &in)
		if err != nil {
			ret := sioErrToString(err)
			s.Emit("error", ret)
			return ret
		}
		ret, err := svc.onEventJoin(s, &in)
		if err != nil {
			ret := sioErrToString(err)
			s.Emit("error", ret)
			return ret
		}
		out, _ := json.Marshal(ret)
		return string(out)
	})
	server.OnEvent("/", "broadcast", func(s socketio.Conn, msg string) string {
		var in SIO_Broadcast_Input
		err := json.Unmarshal([]byte(msg), &in)
		if err != nil {
			ret := sioErrToString(err)
			s.Emit("error", ret)
			return ret
		}
		ret, err := svc.onEventBroadcast(s, &in)
		if err != nil {
			ret := sioErrToString(err)
			s.Emit("error", ret)
			return ret
		}
		out, _ := json.Marshal(ret)
		return string(out)
	})

	return server, nil
}

func sioErrToString(err error) string {
	msg := SIO_Error{IsError: true, ErrMsg: err.Error()}
	out, _ := json.Marshal(msg)
	return string(out)
}
