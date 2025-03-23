package database

import "time"

type Users struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Username   string `gorm:"not null;unique"`
	Email      string `gorm:"not null;unique"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null"`
	IsVerified bool   `gorm:"default:false"`

	Link            []Links         `gorm:"foreignKey:UserID"`
	Qrs             []QR            `gorm:"foreignKey:UserID"`
	Links_histories []Links_history `gorm:"foreignKey:UserID"`
}

type Links struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Uri    string `gorm:"not null"`
	Count  uint   `gorm:"default:0"`
	UserID uint

	QRs             []QR            `gorm:"foreignKey:CurrentLink"`
	Links_histories []Links_history `gorm:"foreignKey:LinkId"`
}

type QR struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CurrentLink uint
	UserID      uint
	Status      bool `gorm:"default:true"`
	Qr_type     string

	Links_histories []Links_history `gorm:"foreignKey:QrId"`
}

type Links_history struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	LinkId uint `gorm:"not null"`
	QrId   uint `gorm:"not null"`
}
