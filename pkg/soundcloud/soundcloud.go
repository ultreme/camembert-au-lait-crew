package calcsoundcloud

import (
	"net/url"
	"strings"

	"github.com/shazow/memoizer"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

type CALCSoundcloud struct {
	client   *soundcloud.Api
	clientID string
	userID   uint64
	cache    memoizer.Memoize
}

func New(clientID string, userID uint64) CALCSoundcloud {
	return CALCSoundcloud{
		client: &soundcloud.Api{
			ClientId: clientID,
		},
		clientID: clientID,
		userID:   userID,
		cache: memoizer.Memoize{
			Cache: memoizer.NewMemoryCache(),
		},
	}
}

func (s *CALCSoundcloud) EscapeString(input string) string {
	return strings.Replace(input, s.clientID, "<SOUNDCLOUD_CLIENT_ID>", -1)
}

func (s *CALCSoundcloud) Me() (*soundcloud.User, error) {
	user, err := s.cache.Call(s.client.User(s.userID).Get, url.Values{})
	return user.(*soundcloud.User), err
}

func (s *CALCSoundcloud) Playlists() ([]*soundcloud.Playlist, error) {
	playlists, err := s.cache.Call(s.client.User(s.userID).Playlists, url.Values{})
	return playlists.([]*soundcloud.Playlist), err
}

func (s *CALCSoundcloud) Playlist(playlistID uint64) (*soundcloud.Playlist, error) {
	inline := func(playlistID uint64) (*soundcloud.Playlist, error) {
		return s.client.Playlist(playlistID).Get(nil)
	}
	playlist, err := s.cache.Call(inline, playlistID)
	return playlist.(*soundcloud.Playlist), err
}

func (s *CALCSoundcloud) Track(trackID uint64) (*soundcloud.Track, error) {
	inline := func(trackID uint64) (*soundcloud.Track, error) {
		return s.client.Track(trackID).Get(nil)
	}
	track, err := s.cache.Call(inline, trackID)
	return track.(*soundcloud.Track), err
}

func (s *CALCSoundcloud) Tracks() ([]*soundcloud.Track, error) {
	tracks, err := s.cache.Call(s.client.User(s.userID).Tracks, url.Values{})
	return tracks.([]*soundcloud.Track), err
}
