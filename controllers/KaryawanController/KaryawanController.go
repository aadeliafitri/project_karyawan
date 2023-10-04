package KaryawanController

import (
	// "database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aadeliafitri/project-karyawan/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	perPage := 5
	offset := (page - 1) * perPage

	var karyawan []models.KaryawanResponse

	// var jabatan []models.JabatanResponse

	sql := "SELECT distinct karyawans.* FROM karyawans join jabatans ON jabatans.id = karyawans.id_jabatan JOIN gaji_jabatans ON gaji_jabatans.jabatan_id = jabatans.id JOIN gaji_tunjangans ON gaji_tunjangans.id = gaji_jabatans.gaji_tunjangan_id"

	if search := c.Query("search"); search != "" {
		sql = fmt.Sprintf("%s WHERE (karyawans.nama LIKE '%%%s%%' OR karyawans.email LIKE '%%%s%%' OR karyawans.telp LIKE '%%%s%%')", sql, search, search, search)
		
		jenisKelamin := c.Query("jenis_kelamin")
		status := c.Query("status")
		jabatan := c.QueryArray("jabatan")

		if jenisKelamin != "" || status != ""{
			sql = fmt.Sprintf("%s AND (karyawans.jenis_kelamin = '%s' OR karyawans.status = '%s'", sql, jenisKelamin, status )
			if jabatan != nil{
				for _, value := range jabatan {
					sql = fmt.Sprintf("%s OR karyawans.id_jabatan IN (%s))", sql, value)
				}
			}else{
				sql = fmt.Sprintf("%s)", sql)
			}
		}else{
			for _, value := range jabatan {
				sql = fmt.Sprintf("%s AND karyawans.id_jabatan IN (%s)", sql, value)
			}
		}
		
		if minGaji := c.Query("min_gaji"); minGaji != "" {

			sql = fmt.Sprintf("%s group by karyawans.id_karyawan having SUM(gaji_tunjangans.nominal) between '%s'", sql, minGaji)

		}
		if maxGaji := c.Query("max_gaji"); maxGaji != "" {
			sql = fmt.Sprintf("%s AND '%s'", sql, maxGaji)
		}
		// fmt.Println(sql)

		models.DB.Debug().Preload("Jabatan.Gaji").Limit(perPage).Offset(offset).Raw(sql).Find(&karyawan)
		c.JSON(http.StatusOK, gin.H{"karyawan": karyawan})
		return
	} else {
		if err := models.DB.Preload("Jabatan.Gaji").Limit(perPage).Offset(offset).Find(&karyawan).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"karyawan": karyawan})
		return
	}

}

func Show(c *gin.Context) {
	var karyawan models.Karyawan

	id := c.Param("id")

	if err := models.DB.Preload("Jabatan.Gaji").First(&karyawan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"karyawan": karyawan})
}

func Create(c *gin.Context) {
	var karyawan models.Karyawan

	// TotalGaji()

	// karyawan.TotalGaji = TotalGaji()
	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&karyawan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Ditambahkan"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambahkan", "karyawan": karyawan})
}

func Update(c *gin.Context) {
	var karyawan models.Karyawan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&karyawan).Where("id_karyawan = ?", id).Updates(&karyawan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diupdate"})
}

func Delete(c *gin.Context) {
	var karyawan models.Karyawan

	id := c.Param("id")

	if err := models.DB.Model(&karyawan).Where("id_karyawan = ?", id).First(&karyawan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Dihapus"})
		return
	}

	models.DB.Delete(&karyawan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil dihapus"})
}

func TotalGaji(c *gin.Context) {
	// var karyawan []models.KaryawanResponse

	var total_gaji int64

	// rows, err := models.DB.Preload("Jabatan.Gaji").Table("jabatans").Select("sum(gaji_tunjangans.nominal) as total_gaji").Joins("JOIN gaji_jabatans ON gaji_jabatans.jabatan_id = jabatans.id").Joins("JOIN gaji_tunjangans ON gaji_tunjangans.id = gaji_jabatans.gaji_tunjangan_id").Where("jabatans.id = karyawans.id_jabatan").Rows()

	if minGaji := c.Query("min_gaji"); minGaji != "" {

	}

	models.DB.Debug().
		Preload("Jabatan.Gaji").
		Table("karyawans").
		Select("sum(gaji_tunjangans.nominal)").
		Joins("JOIN jabatans ON jabatans.id = karyawans.id_jabatan").
		Joins("JOIN gaji_jabatans ON gaji_jabatans.jabatan_id = jabatans.id").
		Joins("JOIN gaji_tunjangans ON gaji_tunjangans.id = gaji_jabatans.gaji_tunjangan_id").
		Where("karyawans.id_jabatan = jabatans.id").
		Group("karyawans.id_karyawan").
		Having("sum(gaji_tunjangans.nominal) BETWEEN ? AND ?").
		Find(&total_gaji)

	// fmt.Println(rows)
	// if err != nil {
	// 	return 0, err
	// } else {
	// 	var total_gaji int64
	// 	for rows.Next() {
	// 		rows.Scan(&total_gaji)
	// 	}
	// 	fmt.Println(total_gaji)
	// 	return total_gaji, nil
	// }
}
