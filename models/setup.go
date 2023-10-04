package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/project_karyawan?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Karyawan{}, &Jabatan{}, &GajiTunjangan{})

	DB = database
}	

func GetConnection() *gorm.DB {
	return DB
}