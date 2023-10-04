package models

import (
	"time"
)

type Jabatan struct {
	ID int64 `gorm:"primaryKey" json:"id_jabatan"`
	Jabatan string `json:"jabatan" gorm:"not null; type:varchar(255)"`
	Gaji []GajiTunjangan `json:"Gaji" gorm:"many2many:gaji_jabatans"`
	GajiID []int64 `json:"gaji_id" gorm:"-"`
	CreatedAt *time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index" gorm:"default:CURRENT_TIMESTAMP"`
}

type JabatanResponse struct {
	ID int64 `json:"id_jabatan"`
	Jabatan string `json:"jabatan"`
	Gaji []GajiTunjangan `json:"Gaji" gorm:"many2many:gaji_jabatans;foreignKey:ID;joinForeignKey:JabatanID;References:ID;joinReferences:GajiTunjanganID"`
	GajiID []int64 `json:"-" gorm:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type JabatanResponseWithoutGaji struct {
	ID int64 `json:"id_jabatan"`
	Jabatan string `json:"jabatan"`
	Gaji []GajiTunjangan `json:"-" gorm:"many2many:gaji_jabatans;foreignKey:ID;joinForeignKey:JabatanID;References:ID;joinReferences:GajiTunjanganID"`
	GajiID []int64 `json:"-" gorm:"-"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

func (JabatanResponse) TableName() string{
	return "jabatans"
}

func (JabatanResponseWithoutGaji) TableName() string{
	return "jabatans"
}