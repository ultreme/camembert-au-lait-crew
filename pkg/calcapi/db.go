package calcapi

import "github.com/jinzhu/gorm"

type DevKeyValueString struct {
	Key   string `gorm:"primary_key"`
	Value string
}

type DevKeyValueFloat struct {
	Key   string `gorm:"primary_key"`
	Value float64
}

func setupDB(db *gorm.DB) error {
	db.LogMode(true)

	err := db.AutoMigrate(&DevKeyValueString{}, &DevKeyValueFloat{}).Error
	if err != nil {
		return err
	}
	// FIXME: more tuning?

	return nil
}
