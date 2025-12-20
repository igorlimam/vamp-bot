package model

type Info struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"index"`
	Kind string
	Info string
}
