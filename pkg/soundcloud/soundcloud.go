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
	meApi := s.client.User(s.userID)
	return meApi.Get(nil)
}
