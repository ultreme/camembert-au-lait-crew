package soundcloud

import (
	"math/rand"
	"net/url"
	"strings"

	"github.com/shazow/memoizer"
	gosoundcloud "github.com/yanatan16/golang-soundcloud/soundcloud"
)

type Soundcloud struct {
	client   *gosoundcloud.Api
	clientID string
	userID   uint64
	cache    memoizer.Memoize
}

func New(clientID string, userID uint64) Soundcloud {
	return Soundcloud{
		client: &gosoundcloud.Api{
			ClientId: clientID,
		},
		clientID: clientID,
		userID:   userID,
		cache: memoizer.Memoize{
			Cache: memoizer.NewMemoryCache(),
		},
	}
}

func (s *Soundcloud) EscapeString(input string) string {
	return strings.Replace(input, s.clientID, "<SOUNDCLOUD_CLIENT_ID>", -1)
}

func (s *Soundcloud) Me() (*User, error) {
	user, err := s.cache.Call(s.client.User(s.userID).Get, url.Values{})
	if err != nil {
		return nil, err
	}
	return fromSoundcloudUser(user.(*gosoundcloud.User)), nil
}

func (s *Soundcloud) GetPlaylists() (*Playlists, error) {
	playlists, err := s.cache.Call(s.client.User(s.userID).Playlists, url.Values{})
	if err != nil {
		return nil, err
	}
	return fromSoundcloudPlaylists(playlists.([]*gosoundcloud.Playlist)), nil
}

func (s *Soundcloud) GetPlaylist(playlistID uint64) (*Playlist, error) {
	inline := func(playlistID uint64) (*gosoundcloud.Playlist, error) {
		return s.client.Playlist(playlistID).Get(nil)
	}
	playlist, err := s.cache.Call(inline, playlistID)
	if err != nil {
		return nil, err
	}
	return fromSoundcloudPlaylist(playlist.(*gosoundcloud.Playlist)), nil
}

func (s *Soundcloud) GetTrack(trackID uint64) (*Track, error) {
	inline := func(trackID uint64) (*gosoundcloud.Track, error) {
		return s.client.Track(trackID).Get(nil)
	}
	track, err := s.cache.Call(inline, trackID)
	if err != nil {
		return nil, err
	}
	return fromSoundcloudTrack(track.(*gosoundcloud.Track)), nil
}

func (s *Soundcloud) GetRandomTrack() (*Track, error) {
	tracks, err := s.GetTracks()
	if err != nil {
		return nil, err
	}

	return tracks.Tracks[rand.Intn(len(tracks.Tracks))], nil
}

func (s *Soundcloud) GetRandomPlaylist() (*Playlist, error) {
	playlists, err := s.GetPlaylists()
	if err != nil {
		return nil, err
	}

	return playlists.Playlists[rand.Intn(len(playlists.Playlists))], nil
}

func (s *Soundcloud) GetTracks() (*Tracks, error) {
	tracks, err := s.cache.Call(s.client.User(s.userID).Tracks, url.Values{})
	if err != nil {
		return nil, err
	}
	return fromSoundcloudTracks(tracks.([]*gosoundcloud.Track)), nil
}

func fromSoundcloudPlaylists(input []*gosoundcloud.Playlist) *Playlists {
	playlists := &Playlists{}
	for _, playlist := range input {
		playlists.Playlists = append(playlists.Playlists, fromSoundcloudPlaylist(playlist))
	}
	return playlists
}

func fromSoundcloudTracks(input []*gosoundcloud.Track) *Tracks {
	tracks := &Tracks{}
	for _, track := range input {
		tracks.Tracks = append(tracks.Tracks, fromSoundcloudTrack(track))
	}
	return tracks
}

func fromSoundcloudPlaylist(input *gosoundcloud.Playlist) *Playlist {
	return &Playlist{
		ID:            input.Id,
		CreatedAt:     input.CreatedAt,
		Title:         input.Title,
		Sharing:       input.Sharing,
		EmbeddableBy:  input.EmbeddableBy,
		PurchaseUrl:   input.PurchaseUrl,
		ArtworkUrl:    input.ArtworkUrl,
		Description:   input.Description,
		Duration:      input.Duration,
		Genre:         input.Genre,
		SharedToCount: input.SharedToCount,
		TagList:       input.TagList,
		ReleaseDay:    uint32(input.ReleaseDay),
		ReleaseMonth:  uint32(input.ReleaseMonth),
		ReleaseYear:   uint32(input.ReleaseYear),
		Streamable:    input.Streamable,
		Downloadable:  input.Downloadable,
		Ean:           input.EAN,
		PlaylistType:  input.PlaylistType,
		Tracks:        fromSoundcloudTracks(input.Tracks).Tracks,
		Uri:           input.Uri,
		Label:         fromSoundcloudUser(input.Label),
		LabelId:       input.LabelId,
		LabelName:     input.LabelName,
		User:          fromSoundcloudUser(input.User),
		UserId:        input.UserId,
		Permalink:     input.Permalink,
		PermalinkUrl:  input.PermalinkUrl,
	}
}

func fromSoundcloudTrack(input *gosoundcloud.Track) *Track {
	return &Track{
		ID:                  input.Id,
		CreatedAt:           input.CreatedAt,
		Title:               input.Title,
		Sharing:             input.Sharing,
		EmbeddableBy:        input.EmbeddableBy,
		PurchaseUrl:         input.PurchaseUrl,
		ArtworkUrl:          input.ArtworkUrl,
		Description:         input.Description,
		Duration:            input.Duration,
		Genre:               input.Genre,
		SharedToCount:       input.SharedToCount,
		TagList:             input.TagList,
		ReleaseDay:          uint32(input.ReleaseDay),
		ReleaseMonth:        uint32(input.ReleaseMonth),
		ReleaseYear:         uint32(input.ReleaseYear),
		Streamable:          input.Streamable,
		Downloadable:        input.Downloadable,
		State:               input.State,
		License:             input.License,
		TrackType:           input.TrackType,
		WaveformUrl:         input.WaveformUrl,
		DownloadUrl:         input.DownloadUrl,
		StreamUrl:           input.StreamUrl,
		VideoUrl:            input.VideoUrl,
		Bpm:                 float32(input.Bpm),
		Commentable:         input.Commentable,
		ISRC:                input.ISRC,
		KeySignature:        input.KeySignature,
		CommentCount:        input.CommentCount,
		DownloadCount:       input.DownloadCount,
		PlaybackCount:       input.PlaybackCount,
		FavoritingsCount:    input.FavoritingsCount,
		OriginalFormat:      input.OriginalFormat,
		OriginalContentSize: input.OriginalContentSize,
		AssetData:           input.AssetData,
		ArtworkData:         input.ArtworkData,
		UserFavorite:        input.UserFavorite,
		Uri:                 input.Uri,
		Label:               fromSoundcloudUser(input.Label),
		LabelId:             input.LabelId,
		LabelName:           input.LabelName,
		User:                fromSoundcloudUser(input.User),
		UserId:              input.UserId,
		Permalink:           input.Permalink,
		PermalinkUrl:        input.PermalinkUrl,
	}
}

func fromSoundcloudUser(input *gosoundcloud.User) *User {
	if input == nil {
		return nil
	}
	return &User{
		ID:                   input.Id,
		Username:             input.Username,
		AvatarURL:            input.AvatarUrl,
		Country:              input.Country,
		FullName:             input.FullName,
		City:                 input.City,
		Description:          input.Description,
		DiscogsName:          input.DiscogsName,
		MyspaceName:          input.MyspaceName,
		Website:              input.Website,
		WebsiteTitle:         input.WebsiteTitle,
		Online:               input.Online,
		TrackCount:           input.TrackCount,
		PlaylistCount:        input.PlaylistCount,
		FollowersCount:       input.FollowersCount,
		FollowingsCount:      input.FollowingsCount,
		PublicFavoritesCount: input.PublicFavoritesCount,
		AvatarData:           input.AvatarData,
		Uri:                  input.Uri,
		Permalink:            input.Permalink,
		PermalinkURL:         input.PermalinkUrl,
	}
}
