package calcdashboard

import "github.com/camembertaulaitcrew/camembert-au-lait-crew/pkg/soundcloud"

type CALCDashboard struct {
	soundcloud *calcsoundcloud.CALCSoundcloud
}

func New() *CALCDashboard {
	return &CALCDashboard{}
}

func (d *CALCDashboard) SetSoundCloud(soundcloud *calcsoundcloud.CALCSoundcloud) {
	d.soundcloud = soundcloud
}
