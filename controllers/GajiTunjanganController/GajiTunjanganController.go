package GajiTunjanganController

import (
	"net/http"
	"strconv"

	"github.com/aadeliafitri/project-karyawan/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	perPage := 3
	offset := (page - 1) * perPage
	
	var gajiTunjangan []models.GajiJabatanResponse

	if err := models.DB.Preload("Jabatan").Limit(perPage).Offset(offset).Find(&gajiTunjangan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"gajiTunjangan": gajiTunjangan})
}

func Show(c *gin.Context) {
	var gajiTunjangan models.GajiTunjangan

	id := c.Param("id")

	if err := models.DB.First(&gajiTunjangan, id).Error; err != nil {
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"gajiTunjangan": gajiTunjangan})
}

func Create(c *gin.Context) {
	var gajiTunjangan models.GajiTunjangan

	if err := c.ShouldBindJSON(&gajiTunjangan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&gajiTunjangan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Ditambahkan"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambahkan", "gajiTunjangan": gajiTunjangan})
}

func Update(c *gin.Context) {
	var gajiTunjangan models.GajiTunjangan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&gajiTunjangan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&gajiTunjangan).Where("id = ?", id).Updates(&gajiTunjangan).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diupdate"})
}

func Delete(c *gin.Context) {
	var gajiTunjangan models.GajiTunjangan

	id := c.Param("id")

	if err := models.DB.Model(&gajiTunjangan).Where("id = ?", id).First(&gajiTunjangan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Dihapus"})
		return
	}

	models.DB.Delete(&gajiTunjangan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})
}