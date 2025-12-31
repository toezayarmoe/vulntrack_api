package models

type Environment struct {
	ID       int    `gorm:"column:env_id"`
	Name     string `gorm:"column:name"`
	Critical int    `gorm:"column:critical"`
	High     int    `gorm:"column:high"`
	Medium   int    `gorm:"column:medium"`
	Low      int    `gorm:"column:low"`
	Info     int    `gorm:"column:info"`
}
