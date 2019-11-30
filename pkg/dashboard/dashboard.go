package dashboard // import "ultre.me/calcbiz/pkg/dashboard"

import (
	"fmt"
	"math/rand"

	"go.uber.org/zap"
	"ultre.me/calcbiz/pkg/soundcloud"
	"ultre.me/calcbiz/pkg/spreadshirt"
)

func (e *Entries) append(entries ...*Entry) {
	e.Entries = append(e.Entries, entries...)
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
		"/img/hackz/moijaime/logo.jpg",
		"Générateur de phrase de moi j'aime",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"3615cryptage",
		"hackz/3615cryptage",
		"/img/hackz/3615cryptage/logo.jpg",
		"Messages codés de James Bond",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"Des Phazms",
		"hackz/phazms",
		"/img/hackz/phazms/logo.jpg",
		"Pokedex de phazms",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"Calculatrice.exe",
		"hackz/calculatrice.exe",
		"/img/hackz/calculatrice/logo.jpg",
		"Pour faire des mathématiques ou d'autres sciences",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"Recettator",
		"hackz/recettator",
		"/img/hackz/recettator/logo.jpg",
		"Des recettes équilibrés et festives",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"Paint",
		"hackz/paint",
		"/img/hackz/paint/logo.jpg",
		"Paint.exe en mode MMORPG",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"Ultreme Tetris",
		"hackz/ultreme-tetris",
		"/img/hackz/tetris/logo.jpg",
		"Pour les balèzes",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"m1ch3l",
		"hackz/m1ch3l",
		"/img/hackz/m1ch3l/logo.jpg",
		"Le meilleur d'entre nous",
		Entry_Hack,
		false,
	))
	entries.append(NewManualEntry(
		"2048",
		"hackz/2048",
		"/img/hackz/2048/logo.jpg",
		"de 7 à 77 ans",
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

func (d *Dashboard) Hackz() (*Entries, error) {
	return d.hackEntries(100)
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
		zap.L().Error("failed to fetch hacks", zap.Error(err))
	} else {
		entries.append(hacks.Entries...)
		zap.L().Debug("fetched hack entries", zap.Int("len", len(hacks.Entries)))
	}

	//
	// tracks (soundcloud)
	//
	// FIXME: add timeout
	tracks, err := d.trackEntries(11)
	if err != nil {
		zap.L().Error("failed to fetch tracks", zap.Error(err))
	} else {
		zap.L().Debug("fetched tracks entries", zap.Int("len", len(tracks.Entries)))
		entries.append(tracks.Entries...)
	}

	//
	// merch
	//
	// FIXME: add timeout
	if (false) {
		merchs, err := d.merchEntries(2)
		if err != nil {
			zap.L().Error("failed to fetch merch", zap.Error(err))
		} else {
			zap.L().Debug("fetched merch entries", zap.Int("len", len(merchs.Entries)))
			entries.append(merchs.Entries...)
		}
	}

	// shuffle the compilation
	entries.shuffle()

	// ensure we have exactly `globalLimit` entries
	for len(entries.Entries) < globalLimit {
		entries.Entries = append(entries.Entries, entries.Entries...)
	}
	entries.Entries = entries.Entries[:globalLimit]
	return entries, nil
}
