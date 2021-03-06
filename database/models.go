package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBScheduleTime struct {
	gorm.Model
	Day string
	Start string `gorm:"unique;not null"`
	End string
	Displayable bool `gorm:"default:false"`
}

type DBScheduleRoom struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type DBScheduleSession struct {
	gorm.Model
	Time *DBScheduleTime
	Room *DBScheduleRoom
	RoomID int
	TimeID int
	UpdaterName string
	UpdaterID string
	Title string
	Speaker string
	UniqueString string `gorm:"unique;type:string"`
	Version int `gorm:"not null;default:0"`
}

func (s DBScheduleSession) ToString() string {
	return fmt.Sprintf("Title: %s, Speaker: %s, Start Time: %s, Room: %s", s.Title, s.Speaker, s.Time.Start, s.Room.Name)
}

func (s DBScheduleSession) ToDmString() string {
	return fmt.Sprintf("Title: %s, Speaker: %s, Start Time: %s, Room: %s, Edit/Delete link: https://talks.twodarek.dev/talks/%s", s.Title, s.Speaker, s.Time.Start, s.Room.Name, s.UniqueString)
}
