package models

import "time"

type GajiTunjangan struct{
	ID int64 `gorm:"primaryKey" json:"id_gaji_tunjangan"`
	Nama string `json:"nama_gaji_tunjangan" gorm:"not null; type:varchar(255)"`
	Nominal int64 `json:"nominal" gorm:"not null"`
	Jenis string `json:"jenis" gorm:"not null; type:varchar(255)"`
	CreatedAt *time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index" gorm:"default:CURRENT_TIMESTAMP"`
}

type GajiJabatanResponse struct{
	ID int64 `json:"id_gaji_tunjangan"`
	Nama string `json:"nama_gaji_tunjangan"`
	Nominal int64 `json:"nominal"`
	Jenis string `json:"jenis"`
	Jabatan []JabatanResponseWithoutGaji `json:"jabatan" gorm:"many2many:gaji_jabatans;foreignKey:ID;joinForeignKey:GajiTunjanganID;References:ID;joinReferences:JabatanID"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func(GajiJabatanResponse) TableName() string {
	return "gaji_tunjangans"
}