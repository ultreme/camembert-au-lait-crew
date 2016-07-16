package calcsoundcloud

import (
	"strings"

	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

type CALCSoundcloud struct {
	client   *soundcloud.Api
	clientID string
	userID   uint64
}

func New(clientID string, userID uint64) CALCSoundcloud {
	return CALCSoundcloud{
		client: &soundcloud.Api{
			ClientId: clientID,
		},
		clientID: clientID,
		userID:   userID,
	}
}

func (s *CALCSoundcloud) EscapeString(input string) string {
	return strings.Replace(input, s.clientID, "<SOUNDCLOUD_CLIENT_ID>", -1)
}

func (s *CALCSoundcloud) Me() (*soundcloud.User, error) {
	return s.client.User(s.userID).Get(nil)
}

func (s *CALCSoundcloud) Playlists() ([]*soundcloud.Playlist, error) {
	return s.client.User(s.userID).Playlists(nil)
}

func (s *CALCSoundcloud) Playlist(playlistID uint64) (*soundcloud.Playlist, error) {
	return s.client.Playlist(playlistID).Get(nil)
}

func (s *CALCSoundcloud) Track(trackID uint64) (*soundcloud.Track, error) {
	return s.client.Track(trackID).Get(nil)
}

func (s *CALCSoundcloud) Tracks() ([]*soundcloud.Track, error) {
	return s.client.User(s.userID).Tracks(nil)
}
