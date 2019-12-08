package calcapi

import (
	"encoding/json"
	"fmt"
	"time"

	engineio "github.com/googollee/go-engine.io"
	socketio "github.com/googollee/go-socket.io"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"moul.io/godev"
)

const socketioRoomLogSize int = 50

func (svc *svc) sioOnConnect(s socketio.Conn) error {
	s.SetContext(&SIO_Context{NoMetadata: true})
	svc.sio.logger.Info(
		"connect",
		zap.String("id", s.ID()),
	)
	return nil
}

func (svc *svc) sioOnError(s socketio.Conn, err error) {
	svc.sio.logger.Warn(
		"error",
		zap.String("id", s.ID()),
		zap.Error(err),
	)
	s.Close()
}

func (svc *svc) sioOnDisconnect(s socketio.Conn, reason string) {
	svc.sio.logger.Info(
		"disconnect",
		zap.String("id", s.ID()),
		zap.String("reason", reason),
		zap.Strings("rooms", s.Rooms()),
	)

	fmt.Println("onDisconnect.Lock")
	defer fmt.Println("onDisconnect.Unlock")
	svc.sio.mutex.Lock()
	defer svc.sio.mutex.Unlock()

	if s == nil || s.Context() == nil {
		svc.sio.logger.Warn("socket is nil", zap.Any("s", s))
	}
	context := s.Context().(*SIO_Context)
	for room, roomPeers := range svc.sio.roomPeers {
		if _, ok := roomPeers[s.ID()]; !ok {
			continue
		}
		delete(svc.sio.roomPeers[room], s.ID())
		broadcast := SIO_Disconnect_Event{
			Room:  room,
			Peer:  context.Peer,
			Peers: svc.peersForRoom(room),
			//PeerCount: int32(svc.sio.server.RoomLen(room)),
		}
		go svc.broadcastToRoom(room, "event:disconnect", broadcast)
		s.Leave(room)
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
		zap.Any("previous-peer", context.Peer),
		zap.Any("peer", in.Peer),
		zap.String("id", s.ID()),
	)

	s.Join(in.Room)

	context.NoMetadata = false
	context.Peer = in.Peer
	s.SetContext(context)

	fmt.Println("onEventJoin.Lock")
	defer fmt.Println("onEventJoin.Unlock")
	svc.sio.mutex.Lock()
	defer svc.sio.mutex.Unlock()

	// update room peers stats
	if _, ok := svc.sio.roomPeers[in.Room]; !ok {
		svc.sio.roomPeers[in.Room] = make(map[string]socketio.Conn)
	}

	svc.sio.roomPeers[in.Room][s.ID()] = s

	// emit old logs
	if in.MaxLogEntries > 0 {
		if logs, ok := svc.sio.logs[in.Room]; ok {
			start := 0
			if len(logs) > int(in.MaxLogEntries) {
				start = len(logs) - int(in.MaxLogEntries)
			}
			for _, log := range logs[start:] {
				out, _ := json.Marshal(log)
				s.Emit("event:broadcast", string(out))
			}
		}
	}

	fmt.Println(godev.PrettyJSON(svc.sio.roomPeers))

	// broadcast join event to the room
	broadcast := SIO_Join_Event{
		Room:  in.Room,
		Peer:  in.Peer,
		Peers: svc.peersForRoom(in.Room),
		//PeerCount: int32(svc.sio.server.RoomLen(in.Room)),
	}

	go svc.broadcastToRoom(in.Room, "event:join", broadcast)

	// reply
	ret := SIO_Join_Output{
		Peers: broadcast.Peers,
	}

	return &ret, nil
}

func (svc *svc) broadcastToRoom(room string, event string, data interface{}) {

	svc.sio.logger.Debug(
		"broadcast to room",
		zap.String("room", room),
		zap.String("event", event),
		//zap.Int("room-len", svc.sio.server.RoomLen(room)),
		zap.Any("data", data),
	)
	out, _ := json.Marshal(data)

	svc.sio.server.BroadcastToRoom(room, event, string(out))

}

func (svc *svc) peersForRoom(room string) []*SIO_Peer {
	peers := []*SIO_Peer{}

	roomPeers, ok := svc.sio.roomPeers[room]
	if !ok {
		return peers
	}

	for _, peer := range roomPeers {
		context := peer.Context().(*SIO_Context)
		peers = append(peers, context.Peer)
	}

	return peers
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

	broadcast := SIO_Broadcast_Event{
		Room:   in.Room,
		Msg:    in.Msg,
		Peer:   context.Peer,
		IsLive: true,
	}
	go svc.broadcastToRoom(in.Room, "event:broadcast", broadcast)

	ret := SIO_Broadcast_Output{}

	// store log
	fmt.Println("onEventBroadcast.Lock")
	defer fmt.Println("onEventBroadcast.Unlock")
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
	server, err := socketio.NewServer(&engineio.Options{
		PingInterval: 5 * time.Second,
		PingTimeout:  10 * time.Second,
	})
	if err != nil {
		return nil, errors.Wrap(err, "socketio.NewServer")
	}
	svc.sio.logger = svc.opts.Logger.Named("sio")
	svc.sio.server = server
	svc.sio.logs = make(map[string][]SIO_Broadcast_Event)
	svc.sio.roomPeers = make(map[string]map[string]socketio.Conn)

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
