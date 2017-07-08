package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type BaseModel struct {
	ID uuid.UUID `gorm:"primary_key;type:varchar(36)"`
}

func Gorm() *gorm.DB {
	return db
}

func FirstWhere() {

}
