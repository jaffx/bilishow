package dao

import (
	"time"
)

type UPInfo struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	Name      string `gorm:"name:id"`
	Mid       string `gorm:"name:mid;index"`
	Sex       string
	Sign      string
	Level     int
	Title     string
	Face      string
	Follower  int
	Ext       string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *UPInfo) TableName() string {
	return "up_infos"
}
