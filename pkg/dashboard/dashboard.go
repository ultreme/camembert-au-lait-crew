package calcdashboard

import (
	"fmt"
	"math/rand"

	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/soundcloud"
	"github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/spreadshirt"
)

type CALCDashboard struct {
	soundcloud *calcsoundcloud.CALCSoundcloud
}

const (
	typeHack  = "hack"
	typeTrack = "track"
	typeMerch = "merch"
)

func New() *CALCDashboard {
	return &CALCDashboard{}
}

func (d *CALCDashboard) SetSoundCloud(soundcloud *calcsoundcloud.CALCSoundcloud) {
	d.soundcloud = soundcloud
}

func (d *CALCDashboard) hackEntries(limit int) (Entries, error) {
	entries := Entries{}
	entries.append(NewManualEntry(typeHack, "Moi j'aime", "hackz/moijaime", "", "Générateur de phrase de moi j'aime"))
	entries.append(NewManualEntry(typeHack, "3615cryptage", "hackz/3615cryptage", "", "Messages codés de James Bond"))

	entries.shuffle()
	if len(entries) < limit {
		limit = len(entries)
	}
	return entries[:limit], nil
}

func (d *CALCDashboard) trackEntries(limit int) (Entries, error) {
	entries := Entries{}

	tracks, err := d.soundcloud.Tracks()
	if err != nil {
		return entries, err
	}
	if len(tracks) < limit {
		limit = len(tracks)
	}

	// shuffle tracks
	for i := range tracks {
		j := rand.Intn(i + 1)
		tracks[i], tracks[j] = tracks[j], tracks[i]
	}

	for _, track := range tracks[:limit] {
		entries.append(NewManualEntry(typeTrack, track.Title, fmt.Sprintf("track/%d", track.Id), track.ArtworkUrl, track.Description))
	}

	entries.shuffle()
	return entries, nil
}

func (d *CALCDashboard) merchEntries(limit int) (Entries, error) {
	entries := Entries{}

	products := calcspreadshirt.GetAllProducts(250, 250)
	if len(products) < limit {
		limit = len(products)
	}
	for _, product := range products[:limit] {
		entries.append(NewManualEntry(typeMerch, product.Title, product.URL, product.ImageURL, ""))
	}

	entries.shuffle()
	return entries, nil
}

func (d *CALCDashboard) Random() (Entries, error) {
	entries := Entries{}

	hacks, err := d.hackEntries(3)
	if err != nil {
		return nil, err
	}
	entries = append(entries, hacks...)

	tracks, err := d.trackEntries(11)
	if err != nil {
		return nil, err
	}
	entries = append(entries, tracks...)

	merchs, err := d.merchEntries(2)
	if err != nil {
		return nil, err
	}
	entries = append(entries, merchs...)

	entries.shuffle()
	return entries, nil
}
