package calcapi

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"moul.io/zapgorm"
)

type DevKeyValueString struct {
	Key   string `gorm:"primary_key"`
	Value string
}

type DevKeyValueFloat struct {
	Key   string `gorm:"primary_key"`
	Value float64
}

func setupDB(db *gorm.DB) error {
	db.SetLogger(zapgorm.New(zap.L().Named("gorm")))
	db.LogMode(true)

	err := db.AutoMigrate(&DevKeyValueString{}, &DevKeyValueFloat{}).Error
	if err != nil {
		return err
	}
	// FIXME: more tuning?

	return nil
}
