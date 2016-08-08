package calcdashboard

import "github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/soundcloud"

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
	entries.append(NewManualEntry(typeHack, "Moi j'aime", "Générateur de phrase de moi j'aime", ""))
	entries.append(NewManualEntry(typeHack, "Kryptos", "Messages codés de James Bond", ""))
	entries.shuffle()
	if len(entries) < limit {
		limit = len(entries)
	}
	return entries[:limit], nil
}

func (d *CALCDashboard) trackEntries(limit int) (Entries, error) {
	entries := Entries{}

	entries.shuffle()
	if len(entries) < limit {
		limit = len(entries)
	}
	return entries[:limit], nil
}

func (d *CALCDashboard) merchEntries(limit int) (Entries, error) {
	entries := Entries{}
	entries.shuffle()
	if len(entries) < limit {
		limit = len(entries)
	}
	return entries[:limit], nil
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
