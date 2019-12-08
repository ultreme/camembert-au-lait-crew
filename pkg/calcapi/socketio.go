package calcapi

import (
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const socketioRoomLogSize int = 50

func (svc *svc) sioOnConnect(s socketio.Conn) error {
	s.SetContext(&SIO_Context{NoMetadata: true})
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

	context := s.Context().(*SIO_Context)
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

func (svc *svc) onEventRooms(s socketio.Conn) *SIO_Rooms_Output {
	context := s.Context().(*SIO_Context)
	svc.sio.logger.Debug(
		"rooms requested",
		zap.String("id", s.ID()),
		zap.Any("peer", context.Peer),
	)
	ret := SIO_Rooms_Output{
		Rooms: s.Rooms(),
	}
	return &ret
}

func (svc *svc) onEventJoin(s socketio.Conn, in *SIO_Join_Input) (*SIO_Join_Output, error) {
	context := s.Context().(*SIO_Context)
	svc.sio.logger.Debug(
		"join",
		zap.String("room", in.Room),
		zap.Any("preview-peer", context.Peer),
		zap.Any("peer", in.Peer),
		zap.String("id", s.ID()),
	)

	s.Join(in.Room)

	context.NoMetadata = false
	context.Peer = in.Peer
	s.SetContext(context)

	broadcast := SIO_Join_Event{Room: in.Room, Peer: in.Peer}
	out, _ := json.Marshal(broadcast)
	svc.sio.server.BroadcastToRoom(in.Room, "event:join", string(out))

	ret := SIO_Join_Output{
		Peers: []*SIO_Peer{{Name: "foo"}, {Name: "bar"}},
	}

	// emit old logs
	svc.sio.mutex.Lock()
	defer svc.sio.mutex.Unlock()
	if logs, ok := svc.sio.logs[in.Room]; ok {
		for _, log := range logs {
			out, _ := json.Marshal(log)
			s.Emit("event:broadcast", string(out))
		}
	}
	return &ret, nil
}

func (svc *svc) onEventBroadcast(s socketio.Conn, in *SIO_Broadcast_Input) (*SIO_Broadcast_Output, error) {
	context := s.Context().(*SIO_Context)
	svc.sio.logger.Debug(
		"broadcast",
		zap.String("room", in.Room),
		zap.Any("peer", context.Peer),
		zap.Any("msg", in.Msg),
		zap.String("id", s.ID()),
	)

	broadcast := SIO_Broadcast_Event{Room: in.Room, Msg: in.Msg, Peer: context.Peer, IsLive: true}
	out, _ := json.Marshal(broadcast)
	svc.sio.server.BroadcastToRoom(in.Room, "event:broadcast", string(out))

	ret := SIO_Broadcast_Output{}

	// store log
	svc.sio.mutex.Lock()
	defer svc.sio.mutex.Unlock()
	if _, ok := svc.sio.logs[in.Room]; !ok {
		svc.sio.logs[in.Room] = make([]SIO_Broadcast_Event, 0)
	}
	broadcast.IsLive = false
	svc.sio.logs[in.Room] = append(svc.sio.logs[in.Room], broadcast)
	existing := len(svc.sio.logs[in.Room])
	if existing > socketioRoomLogSize {
		svc.sio.logs[in.Room] = svc.sio.logs[in.Room][existing-socketioRoomLogSize:]
	}
	return &ret, nil
}

func (svc *svc) SocketIOServer() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, errors.Wrap(err, "socketio.NewServer")
	}
	svc.sio.logger = svc.opts.Logger.Named("sio")
	svc.sio.server = server
	svc.sio.logs = make(map[string][]SIO_Broadcast_Event)

	// core events
	server.OnConnect("/", svc.sioOnConnect)
	server.OnError("/", svc.sioOnError)
	server.OnDisconnect("/", svc.sioOnDisconnect)

	// api
	server.OnEvent("/", "ping", svc.onEventPing)
	server.OnEvent("/", "rooms", func(s socketio.Conn, data string) string {
		ret := svc.onEventRooms(s)
		out, _ := json.Marshal(ret)
		return string(out)
	})
	server.OnEvent("/", "join", func(s socketio.Conn, data string) string {
		var in SIO_Join_Input
		err := json.Unmarshal([]byte(data), &in)
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
	server.OnEvent("/", "broadcast", func(s socketio.Conn, data string) string {
		var in SIO_Broadcast_Input
		err := json.Unmarshal([]byte(data), &in)
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
	data := SIO_Error{IsError: true, ErrMsg: err.Error()}
	out, _ := json.Marshal(data)
	return string(out)
}
