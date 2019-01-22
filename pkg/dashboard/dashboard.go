package dashboard // import "ultre.me/calcbiz/pkg/dashboard"

import (
	"fmt"
	"math/rand"

	"go.uber.org/zap"
	"ultre.me/calcbiz/pkg/soundcloud"
	"ultre.me/calcbiz/pkg/spreadshirt"
)

func (e *Entries) append(entries ...*Entry) {
	for _, entry := range entries {
		e.Entries = append(e.Entries, entry)
	}
}

func NewManualEntry(title, URL, imageURL, description string, kind Entry_Kind, isExternal bool) *Entry {
	return &Entry{
		Title:       title,
		URL:         URL,
		Description: description,
		ImageURL:    imageURL,
		Kind:        kind,
		IsExternal:  isExternal,
	}
}

func (e *Entries) shuffle() {
	for i := range e.Entries {
		j := rand.Intn(i + 1)
		e.Entries[i], e.Entries[j] = e.Entries[j], e.Entries[i]
	}
}

type Options struct {
	Soundcloud *soundcloud.Soundcloud
}

type Dashboard struct{ opts *Options }

func New(opts *Options) *Dashboard { return &Dashboard{opts: opts} }

func newEntries() *Entries {
	return &Entries{Entries: make([]*Entry, 0)}
}

func (d *Dashboard) hackEntries(limit int) (*Entries, error) {
	entries := newEntries()
	entries.append(NewManualEntry(
		"Moi j'aime",
		"hackz/moijaime",
		"",
		"Générateur de phrase de moi j'aime",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"3615cryptage",
		"hackz/3615cryptage",
		"",
		"Messages codés de James Bond",
		Entry_Hack,
		false,
	))

	entries.shuffle()
	if len(entries.Entries) < limit {
		limit = len(entries.Entries)
	}
	entries.Entries = entries.Entries[:limit]
	return entries, nil
}

func (d *Dashboard) trackEntries(limit int) (*Entries, error) {
	entries := newEntries()

	tracks, err := d.opts.Soundcloud.GetTracks()
	if err != nil {
		return entries, err
	}
	if len(tracks.Tracks) < limit {
		limit = len(tracks.Tracks)
	}

	// shuffle tracks
	for i := range tracks.Tracks {
		j := rand.Intn(i + 1)
		tracks.Tracks[i], tracks.Tracks[j] = tracks.Tracks[j], tracks.Tracks[i]
	}

	for _, track := range tracks.Tracks[:limit] {
		entries.append(NewManualEntry(
			track.Title,
			fmt.Sprintf("track/%d", track.ID),
			track.ArtworkUrl,
			track.Description,
			Entry_Track,
			false,
		))
	}

	entries.shuffle()
	return entries, nil
}

func (d *Dashboard) merchEntries(limit int) (Entries, error) {
	entries := Entries{}

	products := spreadshirt.GetAllProducts(250, 250)
	if len(products) < limit {
		limit = len(products)
	}
	for _, product := range products[:limit] {
		entries.append(NewManualEntry(
			product.Title,
			product.URL,
			product.ImageURL,
			"",
			Entry_Merch,
			true,
		))
	}

	entries.shuffle()
	return entries, nil
}

func (d *Dashboard) Random() (*Entries, error) {
	entries := newEntries()

	globalLimit := 16

	// FIXME: parallelize slow calls

	//
	// hacks
	//
	hacks, err := d.hackEntries(3)
	if err != nil {
		return nil, err
	}
	entries.append(hacks.Entries...)
	zap.L().Debug("fetched hack entries", zap.Int("len", len(hacks.Entries)))

	//
	// tracks (soundcloud)
	//
	// FIXME: add timeout
	tracks, err := d.trackEntries(11)
	if err != nil {
		return nil, err
	}
	zap.L().Debug("fetched tracks entries", zap.Int("len", len(tracks.Entries)))
	entries.append(tracks.Entries...)

	//
	// merch
	//
	// FIXME: add timeout
	merchs, err := d.merchEntries(2)
	if err != nil {
		return nil, err
	}
	zap.L().Debug("fetched merch entries", zap.Int("len", len(merchs.Entries)))
	entries.append(merchs.Entries...)

	// shuffle the compilation
	entries.shuffle()

	// ensure we have exactly `globalLimit` entries
	for len(entries.Entries) < globalLimit {
		entries.Entries = append(entries.Entries, entries.Entries...)
	}
	entries.Entries = entries.Entries[:globalLimit]
	return entries, nil
}
