package models

import (
	"time"
)

type Karyawan struct {
	IdKaryawan   int                `json:"id_karyawan" gorm:"primaryKey; autoIncrement:true"`
	Nama         string             `json:"nama_karyawan" gorm:"not null; type:varchar(255)"`
	IdJabatan    int64              `gorm:"type:int" json:"id_jabatan"`
	Jabatan      Jabatan            `json:"jabatan" gorm:"foreignKey:IdJabatan"`
	Alamat       string             `json:"alamat" gorm:"not null; type:text"`
	Email        string             `json:"email" gorm:"not null; type:varchar(255)"`
	Status       string             `json:"status_karyawan" gorm:"not null; type:varchar(255)"`
	Telp         string             `json:"telp" gorm:"not null; type:varchar(255)"`
	JenisKelamin string             `json:"jenis_kelamin" gorm:"type:varchar(50)"`
	// TotalGaji    int64 				`json:"total_gaji" gorm:"type:int"`
	CreatedAt    *time.Time         `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time         `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time         `json:"deletedAt" sql:"index" gorm:"default:CURRENT_TIMESTAMP"`
}

type KaryawanResponse struct {
	IdKaryawan   int                `json:"id_karyawan"`
	Nama         string             `json:"nama_karyawan"`
	IdJabatan    int64              `json:"-"`
	Jabatan      JabatanResponse    `json:"jabatan" gorm:"foreignKey:IdJabatan"`
	Alamat       string             `json:"alamat"`
	Email        string             `json:"email"`
	Status       string             `json:"status_karyawan"`
	Telp         string             `json:"telp"`
	JenisKelamin string             `json:"jenis_kelamin"`
	// TotalGaji    int64 				`json:"total_gaji" gorm:"-"`
	CreatedAt    *time.Time         `json:"createdAt"`
	UpdatedAt    *time.Time         `json:"updatedAt"`
	DeletedAt    *time.Time         `json:"deletedAt" sql:"index"`
}

func (KaryawanResponse) TableName() string {
	return "karyawans"
}

// func TotalGaji() (int64, error) {
// 	// var karyawan []models.KaryawanResponse

// 	rows, err := DB.Table("jabatans").Select("sum(gaji_tunjangans.nominal) as total_gaji").Joins("JOIN gaji_jabatans ON gaji_jabatans.jabatan_id = jabatans.id").Joins("JOIN gaji_tunjangans ON gaji_tunjangans.id = gaji_jabatans.gaji_tunjangan_id").Group("jabatans.id").Rows()

// 	if err != nil {
// 		return 0, err
// 	} else {
// 		var total_gaji int64
// 		for rows.Next() {
// 			rows.Scan(&total_gaji)
// 		}
// 		return total_gaji, nil
// 	}
// }

// type TotalGajiInterface interface {
// 	TotalGaji() (int64, error)
// }
