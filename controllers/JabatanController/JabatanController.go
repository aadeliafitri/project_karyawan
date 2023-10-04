package JabatanController

import (
	"net/http"
	"strconv"

	// "strconv"

	"github.com/aadeliafitri/project-karyawan/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	perPage := 2
	offset := (page - 1) * perPage

	var jabatan []models.JabatanResponse

	if err := models.DB.Preload("Gaji").Limit(perPage).Offset(offset).Find(&jabatan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jabatan": jabatan})
}

func Show(c *gin.Context) {
	var jabatan models.Jabatan

	id := c.Param("id")

	if err := models.DB.Preload("Gaji").First(&jabatan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"jabatan": jabatan})
}

func Create(c *gin.Context) {
	var jabatan models.Jabatan

	if err := c.ShouldBindJSON(&jabatan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&jabatan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Ditambahkan"})
		return
	}

	for _, gajitunjanganID := range jabatan.GajiID {
		gajiJabatan := new(models.GajiJabatan)
		gajiJabatan.JabatanID = jabatan.ID
		gajiJabatan.GajiTunjanganID = gajitunjanganID
		models.DB.Debug().Create(&gajiJabatan)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambahkan", "jabatan": jabatan})
}

func Update(c *gin.Context) {
	var jabatan models.Jabatan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&jabatan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&jabatan).Where("id = ?", id).Updates(&jabatan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Diupdate"})
		return
	}
	if err := models.DB.Debug().Where("jabatan_id = ?", id).Delete(&models.GajiJabatan{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var dataInsertGajiJabatan []models.GajiJabatan

	idInt, _ := strconv.Atoi(id)

	for _, gajitunjanganID := range jabatan.GajiID {
		dataInsertGajiJabatan = append(dataInsertGajiJabatan, models.GajiJabatan{
			JabatanID: int64(idInt),
			GajiTunjanganID: gajitunjanganID, 
		})
	}

	if err := models.DB.Create(&dataInsertGajiJabatan).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diupdate"})
}

func Delete(c *gin.Context) {
	var jabatan models.Jabatan

	id := c.Param("id")
	

	if err := models.DB.Model(&jabatan).Where("id = ?", id).First(&jabatan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Dihapus"})
		return
	}

	models.DB.Delete(&jabatan)
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})
}
